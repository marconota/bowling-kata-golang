package main

import "fmt"

type Game interface {
	Roll(int) error
	Score() int
}

type BowlingGame struct {
	score          int
	standingPins   int
	availableRolls int
}

func NewBowlingGame() BowlingGame {
	return BowlingGame{
		standingPins:   10,
		availableRolls: 3,
	}
}

func (b *BowlingGame) Roll(pins int) error {
	if err := b.assertValidRoll(pins); err != nil {
		return err
	}

	b.availableRolls -= 1

	if pins == 10 {
		b.availableRolls = 0
	}

	b.standingPins -= pins
	b.score += pins

	return nil
}

func (b *BowlingGame) Score() int {
	return b.score
}

func (b *BowlingGame) assertValidRoll(pins int) error {
	if b.availableRolls == 0 {
		return fmt.Errorf("can't roll more times")
	}

	if pins > 10 {
		return fmt.Errorf("more than 10 pins are not allowed in a single roll")
	}

	if b.standingPins < pins {
		return fmt.Errorf("more than 10 pins are not allowed in the same frame")
	}

	return nil
}
