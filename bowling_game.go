package main

import "fmt"

type Game interface {
	Roll(int) error
	Score() int
}

type ResultType int

const (
	Standard ResultType = iota
	Strike
	Spare
)

type BowlingGame struct {
	currentFrameIndex int
	frames            []*frame
}

type frame struct {
	score          int
	scoredPins     int
	availableRolls int
	resultType     ResultType
}

func (f *frame) Score() int {
	return f.score
}

func (f *frame) DownedPins() int {
	return f.scoredPins
}

func (f *frame) ResultType() ResultType {
	return f.resultType
}

func (f *frame) IncreaseScoredPins(pins int) {
	f.scoredPins += pins
}

func (f *frame) Close() {
	if f.availableRolls == 1 && f.scoredPins == 10 {
		f.resultType = Strike
	}

	if f.availableRolls == 0 && f.scoredPins == 10 {
		f.resultType = Spare
	}

	f.availableRolls = 0
}

func (f *frame) AvailableRolls() int {
	return f.availableRolls
}

func (f *frame) DecreaseAvailableRolls() {
	f.availableRolls--
}

func NewBowlingGame() BowlingGame {
	return BowlingGame{
		frames: []*frame{
			{
				scoredPins:     0,
				availableRolls: 2,
				resultType:     Standard,
			},
			{
				scoredPins:     0,
				availableRolls: 2,
				resultType:     Standard,
			},
			{
				scoredPins:     0,
				availableRolls: 2,
				resultType:     Standard,
			},
		},
	}
}

func (b *BowlingGame) Roll(pins int) error {
	if err := b.assertValidRoll(pins); err != nil {
		return err
	}

	frame := b.frames[b.currentFrameIndex]
	frame.DecreaseAvailableRolls()
	frame.IncreaseScoredPins(pins)

	b.applyBonus(pins)

	frame.score += pins

	if pins == 10 || frame.AvailableRolls() == 0 {
		frame.Close()
		b.currentFrameIndex++
	}

	return nil
}

func (b *BowlingGame) Score() int {
	s := 0
	for _, f := range b.frames {
		s += f.score
	}
	return s
}

func (b *BowlingGame) assertValidRoll(pins int) error {
	if len(b.frames) == b.currentFrameIndex {
		return fmt.Errorf("the game is over")
	}

	if b.frames[b.currentFrameIndex].AvailableRolls() == 0 {
		return fmt.Errorf("can't roll more times")
	}

	if pins > 10 {
		return fmt.Errorf("more than 10 pins are not allowed in a single roll")
	}

	if b.frames[b.currentFrameIndex].DownedPins()+pins > 10 {
		return fmt.Errorf("more than 10 pins are not allowed in the same frame")
	}

	return nil
}

//todo: please refactor me
func (b *BowlingGame) applyBonus(rolledPins int) {
	if b.currentFrameIndex == 0 {
		return
	}

	previousFrame := b.frames[b.currentFrameIndex-1]
	currentFrame := b.frames[b.currentFrameIndex]

	if previousFrame.ResultType() == Spare &&
		currentFrame.AvailableRolls() == 1 {
		previousFrame.score += rolledPins
		return
	}
	/**

	caso 1: sono al primo tiro del current frame e ho uno strike nel frame precedente e non ho strike nel frame -2			==> bonus frame precedente
	caso 2: sono al secondo tiro del current frame e ho uno strike nel frame precedente			==> bonus frame precedente
	caso 3: sono al primo tiro del current frame e non ho uno strike nel frame precedente		==> no bonus
	caso 4: sono al secondo tiro del current frame e non ho uno strike nel frame precedente		==> no bonus
	caso 5: sono al primo tiro del current frame e ho strike nei 2 frame precedenti				==> bonus frame precedente + bonus frame -2
	caso 6: sono al secondo tiro del current frame e ho uno strike nei 2 frame precedenti		==> bonus applicato solo al frame precedente

	*/

	if previousFrame.ResultType() != Strike {
		return
	}

	if previousFrame.ResultType() == Strike &&
		currentFrame.AvailableRolls() == 0 {
		previousFrame.score += rolledPins
		return
	}

	if previousFrame.ResultType() == Strike &&
		currentFrame.AvailableRolls() == 1 {
		previousFrame.score += rolledPins
	}

	if b.currentFrameIndex == 1 {
		return
	}

	thirdToLast := b.frames[b.currentFrameIndex-2]

	if thirdToLast.ResultType() == Strike &&
		currentFrame.AvailableRolls() == 1 {
		thirdToLast.score += rolledPins
	}
}
