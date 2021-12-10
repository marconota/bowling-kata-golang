package main

import "fmt"

type Game interface {
	Roll(int) error
	Score() int
}

type BowlingGame struct {
	score int
	CurrentFrame Frame
}

type Frame struct {
	AvailableRolls int
}

func (b *BowlingGame) Roll(pins int) error {
	if b.CurrentFrame.AvailableRolls == 0 {
		return fmt.Errorf("can't roll more times")
	}

	if pins > 10 {
		return fmt.Errorf("more than 10 pins are not allowed")
	}

	b.CurrentFrame.AvailableRolls -= 1

	if pins == 10 {
		b.CurrentFrame.AvailableRolls = 0
	}

	b.score += pins

	return nil
}

func (b *BowlingGame) Score() int {
	return b.score
}

func NewBowlingGame() BowlingGame {
	return BowlingGame{
		CurrentFrame: Frame{
			AvailableRolls: 2,
		},
	}
}
