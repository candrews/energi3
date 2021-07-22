// Copyright 2021 The Energi Core Authors
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
	"math/big"

	"energi.world/core/gen3/core/types"
	"energi.world/core/gen3/energi/params"
	"energi.world/core/gen3/log"
)

const (
	// we compute in microseconds so that we can do integer math with higher precision
	microseconds uint64 = 1000000

	// just here to avoid extra typecasts in calculation
	two uint64 = 2
)

// CalculateBlockTimeEMA computes the exponential moving average of block times
// this will return the EMA of block times as microseconds
// for a description of the EMA algorithm, please see:
// see https://www.itl.nist.gov/div898/handbook/pmc/section4/pmc431.htm
func CalculateBlockTimeEMA(blockTimeDifferences []uint64, emaPeriod uint64) (ema []uint64) {
	sampleSize := len(blockTimeDifferences)
	N := emaPeriod + 1
	ema = make([]uint64, sampleSize)

	// choice of initial condition is important for an EMA. We could use the first
	// block time difference, but instead we'll set it to the target value so our
	// EMA will tend toward the target. However we don't include this value in our
	// EMA series data that we return, we only use it to calculate the first EMA
	emaPrev := params.TargetBlockGap * microseconds
	for i := 0; i < sampleSize; i++ {
		// this formula has a factor of 2/(emaPeriod+1) in a couple places. This is our
		// smoothing coefficient for the EMA, often referred to as alpha. We have
		// not precomputed this value so we don't lose precision on early division
		ema[i] = ((two * blockTimeDifferences[i] * microseconds)/ N) + (emaPrev - ((emaPrev * two)/N))
		emaPrev = ema[i]
	}
	return
}

// CalculateBlockTimeDrift calculates the difference between the target block time
// and the EMA block time. Drift should be a positive value if blocks are too slow
// and a negative value if blocks are too fast
func CalculateBlockTimeDrift(ema []uint64) (drift []int64) {
	target := int64(params.TargetBlockGap * microseconds)
	drift = make([]int64, len(ema))
	for i := range ema {
		drift[i] = int64(ema[i]) - target
	}
	return
}

// CalculateBlockTimeIntegral calculates the integral of the block drift function
// This provides us with some idea fo historical "error", how far the block time
// has been from the target value for the duration of the period
// We use the trapezoidal rule here for integration
func CalculateBlockTimeIntegral(drift []int64) (integral int64) {
	sampleSize := len(drift)
	integral = 0
	// this is a simplification of the trapezoid rule based on uniform spacing
	for i := 1; i < sampleSize - 1; i++ {
		integral += drift[i]
	}
	integral += (drift[0] + drift[sampleSize-1]) / 2
	return
}

// CalculateBlockTimeDerivative computes the derivative series of a data series
// Here we use the central difference formula, for some small step h (each block)
// f'(x) = 1/2h * (f(x+h) - f(x-h))
func CalculateBlockTimeDerivative(drift []int64) (derivative []int64) {
	sampleSize := len(drift)
	derivative = make([]int64, sampleSize - 2)

	for i := 1; i < sampleSize - 1; i++ {
		derivative[i-1] = (drift[i+1] - drift[i-1]) / 2
	}
	return
}

/*
 * Block Time Target Calculation V2
 * @chain Current Chain
 * @parent Parent Block Header
 * @ret Time Target structure
 * Populates ret with an updated Time Target
 * Calculates a Target Block Time based on previous block times in order to maintain a 60 second average time
 * Implements the Exponential Moving Average in calculating the block target time
 * Based on the last 60 elapsed block times
 * A block cannot be created with a time greater than 3 seconds in the future
 * ~~The minimum block time is 30 seconds~~ - This should not be enforced
here as an early or late target is for difficulty adjustment not the block
timestamp
*/
func (e *Energi) calcTimeTargetV2(chain ChainReader, parent *types.Header) *timeTarget {

	ret := &timeTarget{}
	parentBlockTime := parent.Time // Defines the original parent block time.
	parentNumber := parent.Number.Uint64()

	// POS-11: Block time restrictions
	ret.max = e.now() + params.MaxFutureGap

	// POS-11: Block time restrictions
	ret.min = parentBlockTime + params.MinBlockGap
	ret.blockTarget = parentBlockTime + params.TargetBlockGap
	ret.periodTarget = ret.blockTarget

	// Block interval enforcement
	// TODO: LRU cache here for extra DoS mitigation
	timeDiffs := make([]uint64, params.AveragingWindow)

	// compute block time differences
	// note that the most recent time difference will be the most
	// weighted by the EMA, and the oldest time difference will be the least
	for i := params.AveragingWindow; i > 0; i-- {
		past := chain.GetHeader(parent.ParentHash, parent.Number.Uint64()-1)
		if past == nil {
			// this normally can't happen because there is more
			// than enough blocks before the hard fork to always
			// get params.AveragingWindow timestamps
			log.Trace("Inconsistent tree, shutdown?")
			return ret
		}
		timeDiffs[i-1] = parent.Time - past.Time
		parent = past
	}

	ema := CalculateBlockTimeEMA(timeDiffs, params.AveragingWindow)
	ret.periodTarget = ema[len(ema)-1]

	log.Trace("PoS time", "block", parentNumber+1,
		"min", ret.min, "max", ret.max,
		"timeTarget", ret.blockTarget,
		"averageBlockTimeMicroseconds", ret.periodTarget,
	)
	return ret
}

/*
 * Difficulty algorithm V2
 * Returns a difficulty value to be used in the next Block
 * @newBlockTime Last Block Time
 * @parent Parent Block Header
 * @timeTarget Target Block Time
 * If the block time is less than the minimum time, the difficulty must be increased
 * If the block time is the target time, the difficulty should stay the same
 * If the block time is more than the target time the difficulty must be reduced
 * New Difficulty = Parent Difficulty * (1.0001 ^ Block Time)
 */
func calcPoSDifficultyV2(
	newBlockTime uint64,
	parent *types.Header,
	timeTarget *timeTarget,
) *big.Int {

	target := timeTarget.blockTarget
	// if the target is the new block time we use the parent difficulty
	if newBlockTime == target {
		log.Trace("No difficulty change", "parent", parent.Difficulty)
		return parent.Difficulty
	}
	// The divergence from the target time to the new block time
	// determines the new difficulty
	targetDivergence := int(newBlockTime) - int(target)
	// clamp to minimum -30
	if targetDivergence < params.MaxTimeDifferenceDrop {
		targetDivergence = params.MaxTimeDifferenceDrop
	}
	// clamp to maximum 60
	if targetDivergence > int(params.TargetBlockGap) {
		targetDivergence = int(params.TargetBlockGap)
	}
	log.Debug(">>>","target", targetDivergence)
	const factorInverse = 10000               // 0.0001 is the same as 1/10000
	const precision = 1000000                 // we want 2 decimal places precision lower
	var scaledPreMultiplier = precision + 100 // this levels it to 1 by
	// dividing the result back by precision

	negative := false
	if targetDivergence < 0 {
		targetDivergence = -targetDivergence
		negative = true
	}
	for i := 0; i < targetDivergence; i++ {
		// the function of 1.0001 ^ timeDiff means the same as
		// repeatedly add 1/10000th to the previous result value as many
		// times as timeDiff, starting with an initial (scaled) value
		if !negative {
			scaledPreMultiplier += scaledPreMultiplier / factorInverse
		} else {
			scaledPreMultiplier -= scaledPreMultiplier / factorInverse
		}
	}
	// multiply the parent difficulty by the multiplier and divide back
	// by the precision value, applying the difficulty change without using
	// floating point numbers
	difficulty := big.NewInt(0).Mul(parent.Difficulty, big.NewInt(int64(scaledPreMultiplier)))
	difficulty = difficulty.Div(difficulty, big.NewInt(int64(precision)))

	log.Trace("Difficulty change",
		"parent", parent.Difficulty, "new difficulty", difficulty,
		"block time", newBlockTime, "target time", target)
	return difficulty
}
