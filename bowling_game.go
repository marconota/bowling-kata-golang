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
	score             int
	currentFrameIndex int
	frames            []*frame
}

type frame struct {
	scoredPins     int
	availableRolls int
	resultType     ResultType
}

func (f *frame) DownedPins() int {
	return f.scoredPins
}

func (f *frame) ResultType() ResultType {
	return f.resultType
}

func (f *frame) IncreaseScore(pins int) {
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
		},
	}
}

func (b *BowlingGame) Roll(pins int) error {
	rollScore := pins
	if err := b.assertValidRoll(pins); err != nil {
		return err
	}

	frame := b.frames[b.currentFrameIndex]
	frame.DecreaseAvailableRolls()
	frame.IncreaseScore(pins)

	if b.shouldDoublePoints() {
		rollScore *= 2
	}

	b.score += rollScore
	if pins == 10 || frame.AvailableRolls() == 0 {
		frame.Close()
		b.currentFrameIndex++
	}

	return nil
}

func (b *BowlingGame) Score() int {
	return b.score
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

func (b *BowlingGame) shouldDoublePoints() bool {

	if b.currentFrameIndex != 0 &&
		b.frames[b.currentFrameIndex-1].ResultType() == Spare &&
		b.frames[b.currentFrameIndex].AvailableRolls() == 1 {
		return true
	}
	if b.currentFrameIndex != 0 &&
		b.frames[b.currentFrameIndex-1].ResultType() == Strike {
		return true
	}
	return false
}
