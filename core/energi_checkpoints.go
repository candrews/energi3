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

package core

import (
	"errors"
	"math/big"
	"sort"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"

	energi_params "energi.world/core/gen3/energi/params"
)

type CheckpointValidateChain interface {
	GetHeaderByNumber(number uint64) *types.Header
}

type CheckpointChain interface {
	CheckpointValidateChain

	EnforceCheckpoint(cp Checkpoint) error
	Config() *params.ChainConfig
}

type Checkpoint struct {
	Since  uint64
	Number uint64
	Hash   common.Hash
}

type CheckpointSignature []byte

type FutureCheckpoint struct {
	Checkpoint
	signatures map[common.Address]CheckpointSignature
}

type checkpointManager struct {
	validated map[uint64]Checkpoint
	latest    uint64
	future    map[uint64]FutureCheckpoint
	mtx       sync.RWMutex
}

func newCheckpointManager() *checkpointManager {
	return &checkpointManager{
		validated: make(map[uint64]Checkpoint),
		future:    make(map[uint64]FutureCheckpoint),
	}
}

func (cm *checkpointManager) setup(chain CheckpointChain) {
	genesis_hash := chain.GetHeaderByNumber(0).Hash()

	if checkpoints, ok := energi_params.EnergiCheckpoints[genesis_hash]; ok {
		for k, v := range checkpoints {
			cm.addCheckpoint(
				chain,
				Checkpoint{
					Number: k,
					Hash:   v,
				},
				[]CheckpointSignature{},
				true,
			)
		}
	}
}

func (cm *checkpointManager) validate(chain CheckpointValidateChain, num uint64, hash common.Hash) error {
	cm.mtx.Lock()
	defer cm.mtx.Unlock()

	// Check against validated checkpoints & mismatch
	if cp, ok := cm.validated[num]; ok {
		if cp.Hash != hash {
			return ErrCheckpointMismatch
		}

		return nil
	}

	// Check if before the latest checkpoint & mismatch
	if num < cm.latest {
		header := chain.GetHeaderByNumber(num)

		if header != nil && header.Hash() != hash {
			return ErrCheckpointMismatch
		}

		return nil
	}

	// TODO: proper future checkpoint processing
	if cp, ok := cm.future[num]; ok {
		if cp.Hash != hash {
			return ErrCheckpointMismatch
		}

		return nil
	}

	return nil
}

func (bc *BlockChain) AddCheckpoint(
	cp Checkpoint,
	sigs []CheckpointSignature,
	local bool,
) error {
	return bc.checkpoints.addCheckpoint(bc, cp, sigs, local)
}

func (cm *checkpointManager) addCheckpoint(
	chain CheckpointChain,
	cp Checkpoint,
	sigs []CheckpointSignature,
	local bool,
) (err error) {
	cm.mtx.Lock()
	defer cm.mtx.Unlock()

	if curr, ok := cm.validated[cp.Number]; ok {
		if curr == cp {
			return nil
		}

		if curr.Since > cp.Since {
			return nil
		}
	}

	if !local {
		// TODO: proper validation and use of future checkpoints
		if len(sigs) == 0 {
			log.Warn("Checkpoint: missing signatures",
				"num", cp.Number, "hash", cp.Hash)
			return errors.New("missing checkpoint signatures")
		}

		// The first one must always be CPP_signer
		pubkey, err := crypto.Ecrecover(cm.hashToSign(&cp), sigs[0][:])
		if err != nil {
			log.Warn("Checkpoint: failed to extract signature",
				"num", cp.Number, "hash", cp.Hash, "err", err)
			return err
		}

		// Check the primary signature
		var signer common.Address
		copy(signer[:], crypto.Keccak256(pubkey[1:])[12:])
		if nrgconf := chain.Config().Energi; nrgconf == nil || signer != nrgconf.CPPSigner {
			log.Warn("Checkpoint: invalid CPP signature", "num", cp.Number, "hash", cp.Hash)
			return errors.New("invalid CPP signature")
		}
	}

	cm.validated[cp.Number] = cp
	log.Info("Added new checkpoint", "checkpoint", cp, "local", local)

	return chain.EnforceCheckpoint(cp)
}

func (cm *checkpointManager) hashToSign(cp *Checkpoint) []byte {
	data := []byte("||Energi Blockchain Checkpoint||")
	data = append(data, common.BigToHash(new(big.Int).SetUint64(cp.Number)).Bytes()...)
	data = append(data, cp.Hash.Bytes()...)
	return crypto.Keccak256(data)
}

func (bc *BlockChain) EnforceCheckpoint(cp Checkpoint) error {
	header := bc.GetHeaderByNumber(cp.Number)

	if header != nil && header.Hash() != cp.Hash {
		log.Error("Side chain is detected as canonical", "number", cp.Number, "hash", cp.Hash)

		if cp_block := bc.GetBlock(cp.Hash, cp.Number); cp_block != nil {
			// Known block
			bc.mu.Lock()
			defer bc.mu.Unlock()

			if err := bc.reorg(bc.GetBlock(header.Hash(), cp.Number), cp_block); err != nil {
				log.Crit("Failed to reorg", "err", err)
				// should terminate
				return err
			}
			log.Warn("Chain reorg was successful, resuming normal operation")
		} else {
			// Unknown block
			if err := bc.SetHead(cp.Number - 1); err != nil {
				log.Crit("Failed to rewind before fork point", "err", err)
				// should terminate
				return err
			}
			log.Warn("Chain rewind was successful, resuming normal operation")
		}
	}

	return nil
}

func (bc *BlockChain) ListCheckpoints() []Checkpoint {
	cm := bc.checkpoints

	cm.mtx.Lock()
	defer cm.mtx.Unlock()

	res := make([]Checkpoint, 0, len(cm.validated))

	for _, v := range cm.validated {
		res = append(res, v)
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i].Number < res[j].Number
	})

	return res
}
