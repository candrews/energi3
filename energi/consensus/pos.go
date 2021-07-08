// Copyright 2019 The Energi Core Authors
// This file is part of the Energi Core library.
//
// The Energi Core library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Energi Core library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Energi Core library. If not, see <http://www.gnu.org/licenses/>.

package consensus

import (
	"encoding/binary"
	"errors"
	"math/big"
	"sort"
	"time"

	"energi.world/core/gen3/common"
	"energi.world/core/gen3/consensus"
	"energi.world/core/gen3/core/types"
	"energi.world/core/gen3/crypto"
	"energi.world/core/gen3/energi/params"
	"energi.world/core/gen3/log"
)

// removing this because why clutter namespaces more than they have to be?
// const (
//	MaturityPeriod    = params.MaturityPeriod
//	AveragingWindow = params.AveragingWindow
//	TargetBlockGap    = params.TargetBlockGap
//	MinBlockGap       = params.MinBlockGap
//	MaxFutureGap      = params.MaxFutureGap
//	TargetPeriodGap   = params.TargetPeriodGap
// )

var (
	minStake    = big.NewInt(1e18) // 1000000000000000000
	diff1Target = new(big.Int).Exp(
		big.NewInt(2), big.NewInt(256), big.NewInt(0),
	)

	errBlockMinTime    = errors.New("block is before minimum time")
	errInvalidPoSHash  = errors.New("invalid PoS hash")
	errInvalidPoSNonce = errors.New("invalid stake weight")
)

type timeTarget struct {
	min uint64
	max uint64
	blockTarget uint64
	periodTarget uint64
}

/**
 * Implements block time consensus
 *
 * POS-11: Block time restrictions
 * POS-12: Block interval enforcement
 */
func (e *Energi) calcTimeTarget(
	chain ChainReader, parent *types.Header,
) (ret *timeTarget) {

	ret = new(timeTarget)
	now := e.now()
	parentNumber := parent.Number.Uint64()
	blockNumber := parentNumber + 1

	// POS-11: Block time restrictions
	ret.max = now + params.MaxFutureGap

	// POS-11: Block time restrictions
	ret.min = parent.Time + params.MinBlockGap
	ret.blockTarget = parent.Time + params.TargetBlockGap
	ret.periodTarget = ret.blockTarget

	// POS-12: Block interval enforcement
	// ---
	if blockNumber >= params.AveragingWindow {
		// TODO: LRU cache here for extra DoS mitigation
		past := parent

		// NOTE: we have to do this way as parent may be not part of canonical
		//       chain. As no mutex is held, we cannot do checks for canonical.
		for i := params.AveragingWindow - 1; i > 0; i-- {
			past = chain.GetHeader(past.ParentHash, past.Number.Uint64()-1)

			if past == nil {
				log.Trace("Inconsistent tree, shutdown?")
				return
			}
		}

		ret.periodTarget = past.Time + params.TargetPeriodGap
		periodMinTime := ret.periodTarget - params.MinBlockGap

		if periodMinTime > ret.min {
			ret.min = periodMinTime
		}
	}

	log.Trace(
		"PoS time", "block", blockNumber,
		"min", ret.min, "max", ret.max,
		"blockTarget", ret.blockTarget,
		"periodTarget", ret.periodTarget,
	)
	return
}

func (e *Energi) enforceMinTime(
	header *types.Header, timeTarget *timeTarget,
) error {

	// NOTE: allow Miner to hint already tried period by
	if header.Time < timeTarget.min {
		header.Time = timeTarget.min
	}

	return nil
}

func (e *Energi) checkTime(header *types.Header, timeTarget *timeTarget) error {

	if header.Time < timeTarget.min {
		return errBlockMinTime
	}

	// Check if able to mine
	if header.Time > timeTarget.max {
		return consensus.ErrFutureBlock
	}

	return nil
}

/**
 * Implements check modifier consensus
 *
 * POS-14: Stake modifier
 */
func (e *Energi) calcPoSModifier(
	chain ChainReader, time uint64, parent *types.Header,
) (ret common.Hash) {
	// TODO: LRU cache here for extra DoS mitigation

	// Find maturity period border
	maturityBorder := time

	if maturityBorder < params.MaturityPeriod {
		// This should happen only in testing
		maturityBorder = 0
	} else {
		maturityBorder -= params.MaturityPeriod
	}

	// Find the oldest inside maturity period
	// NOTE: we have to do this walk as parent may not be part of the canonical chain
	parentHeight := parent.Number.Uint64()
	oldest := parent

	for header, num := oldest, oldest.Number.Uint64(); (header.Time > maturityBorder) && (num > 0); {

		oldest = header
		num--
		header = chain.GetHeader(header.ParentHash, num)
	}

	// Create Stake Modifier
	//
	// The stake modifier is computed by hashing the parent coinbase and the root state of the block nearest to the
	// maturityBorder
	ret = crypto.Keccak256Hash(
		parent.Coinbase.Bytes(),
		oldest.Root.Bytes(),
	)

	log.Trace(
		"PoS modifier", "block", parentHeight+1,
		"modifier", ret, "oldest", oldest.Number.Uint64(),
	)
	return
}

/**
 * Implements difficulty consensus
 */
func (e *Energi) calcPoSDifficulty(
	chain ChainReader,
	time uint64,
	parent *types.Header,
	tt *timeTarget,
) (ret *big.Int) {
	ret = e.diffFn(chain, time, parent, tt)
	log.Trace(
		"PoS difficulty", "block", parent.Number.Uint64()+1, "time", time,
		"diff", ret,
	)
	return ret
}

/**
 * POS-13: Difficulty algorithm (Proposal v1)
 */
const (
	diffV1_BMax     uint64 = 30
	diffV1_AMax     uint64 = 120
	diffV1_DivPlain uint64 = 100

	// Roughly get 2x difficulty decrease
	diffV1_MigrationStakerDelay  uint64 = 15
	diffV1_MigrationStakerTarget uint64 = 0xFFFF
)

var (
	diffV1_BTable []*big.Int
	diffV1_ATable []*big.Int
	diffV1_Div    = new(big.Int).SetUint64(diffV1_DivPlain)
)

func initDiffTable(l uint64, c float64) []*big.Int {
	t := make([]*big.Int, l+1)
	t[0] = common.Big1
	var acc float64 = 1
	for i := 1; i < len(t); i++ {
		acc *= c
		t[i] = big.NewInt(int64(acc * float64(diffV1_DivPlain)))
	}
	return t
}

func init() {
	diffV1_BTable = initDiffTable(diffV1_BMax, 1.1)
	diffV1_ATable = initDiffTable(diffV1_AMax, 1.05)
}

func calcPoSDifficultyV1(
	chain ChainReader,
	time uint64,
	parent *types.Header,
	tt *timeTarget,
) (D *big.Int) {
	// Find the target anchor
	target := (tt.blockTarget + tt.periodTarget) / 2
	if target < tt.min {
		target = tt.min
	}

	if time < target {
		targetDelta := target - time
		if targetDelta > diffV1_BMax {
			targetDelta = diffV1_BMax
		}

		B := diffV1_BTable[targetDelta]
		D = new(big.Int).Div(new(big.Int).Mul(parent.Difficulty, B), diffV1_Div)
		log.Trace("Diff multiplier", "before", targetDelta, "mult", B)

	} else if time > target {
		targetDelta := time - target
		// clamp the target delta to max
		if targetDelta > diffV1_AMax {
			targetDelta = diffV1_AMax
		}
		A := diffV1_ATable[targetDelta]
		D = new(big.Int).Div(
			new(big.Int).Mul(parent.Difficulty, diffV1_Div),
			A,
		)
		log.Trace("Diff multiplier", "after", targetDelta, "div", A)

	} else {
		log.Trace("No difficulty change", "parent", parent.Difficulty)
		return parent.Difficulty
	}

	if D.Cmp(common.Big1) < 0 {
		D = common.Big1
	}

	log.Trace(
		"Difficulty change",
		"parent", parent.Difficulty, "new", D, "time", time, "target", target,
	)
	return D
}

/**
 * Implements hash consensus
 *
 * POS-18: PoS hash generation
 * POS-22: Partial stake amount
 */
func (e *Energi) calcPoSHash(
	header *types.Header,
	target *big.Int,
	weight uint64,
) (posHash *big.Int, usedWeight uint64) {
	serializedTime := [8]byte{}
	binary.BigEndian.PutUint64(serializedTime[:], header.Time)

	posHash = new(big.Int).SetBytes(
		crypto.Keccak256(
			serializedTime[:],
			header.MixDigest.Bytes(),
			header.Coinbase.Bytes(),
		),
	)

	if posHash.Cmp(target) > 0 {
		count, mod := new(big.Int).DivMod(posHash, target, new(big.Int))
		usedWeight = count.Uint64()

		if mod.Cmp(common.Big0) > 0 {
			usedWeight += 1
		}

	} else {
		usedWeight = 1
	}

	if weight < usedWeight {
		return nil, 0
	}

	log.Trace(
		"PoS hash",
		"target", target,
		"posHash", posHash,
		"used_weight", usedWeight,
		"weight", weight,
	)
	return posHash, usedWeight
}

func (e *Energi) verifyPoSHash(chain ChainReader, header *types.Header) error {
	parent := chain.GetHeader(header.ParentHash, header.Number.Uint64()-1)
	if parent == nil {
		return consensus.ErrUnknownAncestor
	}

	weight, err := e.lookupStakeWeight(
		chain, header.Time, parent, header.Coinbase,
	)
	if err != nil {
		return err
	}

	target := new(big.Int).Div(diff1Target, header.Difficulty)

	posHash, usedWeight := e.calcPoSHash(header, target, weight)

	if posHash == nil {
		return errInvalidPoSHash
	}

	if usedWeight != header.Nonce.Uint64() {
		return errInvalidPoSNonce
	}

	return nil
}

/**
 * Implements stake amount calculation.
 *
 * POS-3: Stake maturity period
 * POS-4: Stake amount
 * POS-22: Partial stake amount
 *
 * This is a basic helper for stake amount calculation.
 * There are ways to optimize it for high load, but we need something
 * to start with.
 */
func (e *Energi) lookupStakeWeight(
	chain ChainReader,
	now uint64,
	until *types.Header,
	addr common.Address,
) (weight uint64, err error) {
	var since uint64

	if now > params.MaturityPeriod {
		since = now - params.MaturityPeriod

	} else {
		since = 0
	}

	// NOTE: Do not set to high initial value due to defensive coding approach!
	weight = 0
	totalStaked := uint64(0)
	firstRun := true
	blockState := chain.CalculateBlockState(until.Hash(), until.Number.Uint64())

	// NOTE: we need to ensure at least one iteration with the balance condition
	for (until.Time > since) || firstRun {

		if blockState == nil {
			log.Warn("PoS state root failure", "header", until.Hash())
			return 0, consensus.ErrMissingState
		}

		weightAtBlock := new(big.Int).Div(
			blockState.GetBalance(addr),
			minStake,
		).Uint64()

		if firstRun {
			weight = weightAtBlock
			firstRun = false
		}

		// Find the minimum balance
		if weight > weightAtBlock {
			weight = weightAtBlock
		}

		// No need to lookup further
		if weight < 1 {
			break
		}

		// POS-22: partial stake amount
		if until.Coinbase == addr {
			totalStaked += until.Nonce.Uint64()
		}

		curr := until
		parentNumber := curr.Number.Uint64() - 1
		until = chain.GetHeader(curr.ParentHash, parentNumber)

		if until == nil {

			if curr.Number.Cmp(common.Big0) == 0 {
				break
			}

			log.Error("PoS state missing parent", "parent", curr.ParentHash)
			return 0, consensus.ErrUnknownAncestor
		}

		blockState = chain.CalculateBlockState(curr.ParentHash, parentNumber)
	}

	if weight < totalStaked {
		log.Debug(
			"Nothing to stake",
			"addr", addr, "since", since, "weight", weight, "total_staked",
			totalStaked,
		)
		weight = 0
	} else {
		weight -= totalStaked
	}

	// log.Trace("PoS stake weight", "addr", addr, "weight", weight)
	return weight, nil
}

/**
 * POS-19: PoS miner implementation
 */
func (e *Energi) mine(
	chain ChainReader, header *types.Header, stop <-chan struct{},
) (success bool, err error) {

	type Candidates struct {
		addr   common.Address
		weight uint64
	}

	accounts := e.accountsFn()
	// if no accounts are found, just pause and wait for the stop signal
	//
	// todo: is this what is intended? Is there a case where this value can
	//  change but this thread is then dead? I think that very likely this
	//  should be a repeating loop that retries every minute or something like
	//  this in case an account is created that can be used, while the server is
	//  running.
	//  further thoughts: generally it seems like this function returns quite
	//  quickly so probably this is fine
	if len(accounts) == 0 {
		select {
		case <-stop:
			return
		}
	}

	parent := chain.GetHeader(header.ParentHash, header.Number.Uint64()-1)
	if parent == nil {
		err = consensus.ErrUnknownAncestor
		return
	}

	timeTarget := e.calcTimeTarget(chain, parent)
	blockTime := timeTarget.min

	// Special case due to expected very large gap between Genesis and Migration
	if header.IsGen2Migration() && !e.testing {
		blockTime = e.now()
	}

	candidates := make([]Candidates, 0, len(accounts))
	migrationDPoS := false

	for _, a := range accounts {

		if a == params.Energi_MigrationContract {
			migrationDPoS = true
		}

		candidate := Candidates{a, 0}
		candidates = append(candidates, candidate)
	}

	// check if Asgard hardfork is activated use new difficulty algorithm
	isAsgardActive, err := e.hardforkIsActive(chain, header, "Asgard")
	log.Debug("hf check", "isAsgardActive", isAsgardActive)
	if err != nil {
		log.Trace("Asgard hf check failed: " + err.Error())
	}

	// A special workaround to obey target time when migration contract is
	// used for mining to prevent any difficulty bombs.
	if migrationDPoS && !e.testing {
		// new version modifications
		if isAsgardActive && header.Number.Uint64() < params.DiffV2MigrationStakerBlockDelay && blockTime < timeTarget.blockTarget {
			blockTime = timeTarget.blockTarget
		}

		if isAsgardActive && header.Number.Uint64() < params.DiffV2MigrationStakerBlockDelay && header.Difficulty.Uint64() > params.DiffV2MigrationStakerTarget {
			blockTime += params.DiffV2MigrationStakerTimeDelay
		}

		// old version modifications
		if isAsgardActive == false && blockTime < timeTarget.blockTarget {
			blockTime = timeTarget.blockTarget
		}

		if isAsgardActive == false && blockTime < timeTarget.periodTarget {
			blockTime = timeTarget.periodTarget
		}

		if isAsgardActive == false &&  header.Difficulty.Uint64() > diffV1_MigrationStakerTarget {
			blockTime += diffV1_MigrationStakerDelay
		}

	}

	// Try to match target
	for ; ; blockTime++ {

		if max_time := e.now() + params.MaxFutureGap; blockTime > max_time {
			log.Trace("PoS miner is sleeping")
			select {
			case <-stop:
				// NOTE: it's very important to ignore stop until all variants are tried
				//       to prevent rogue stakers taking the initiative.
				return false, nil
			case <-time.After(time.Duration(blockTime-max_time) * time.Second):
			}
		}


		// check account axistance
		if e.peerCountFn() == 0 {
			log.Trace("Skipping PoS miner due to missing peers")
			continue
		}

		header.Time = blockTime
		if timeTarget, err = e.PoSPrepare(
			chain, header, parent,
		); err != nil {
			return false, err
		}

		target := new(big.Int).Div(diff1Target, header.Difficulty)
		log.Trace("PoS miner time", "time", blockTime)

		// It could be done once, but then there is a chance to miss blocks.
		// Some significant algo optimizations are possible, but we start with simplicity.
		for i := range candidates {
			candidate := &candidates[i]
			candidate.weight, err = e.lookupStakeWeight(
				chain, blockTime, parent, candidate.addr)
			if err != nil {
				return false, err
			}
		}
		// Try smaller amounts first
		sort.Slice(candidates, func(i, j int) bool {
			return candidates[i].weight < candidates[j].weight
		})

		// This tries each candidate for each timestamp before progressing the
		// timestamp. If the reverse order was desired, the block time needs to
		// be saved and reset here. Since older is better, this is probably the
		// better sequence to work in.
		for _, candidate := range candidates {

			if candidate.weight < 1 {
				continue
			}

			//log.Trace("PoS stake candidate", "addr", v.addr, "weight", v.weight)
			header.Coinbase = candidate.addr
			posHash, usedWeight := e.calcPoSHash(header, target, candidate.weight)
			nonceCap := e.GetMinerNonceCap()

			header.Nonce = types.EncodeNonce(usedWeight)
			if nonceCap != 0 && nonceCap < usedWeight {
				continue
			} else if posHash != nil {
				log.Trace(
					"PoS stake", "addr", candidate.addr,
					"weight", candidate.weight,
					"used_weight", usedWeight,
				)
				success = true
				return success, err
			}
		}
	}
	// this doesn't need to be strictly stated but this is a long function
	return success, err
}
