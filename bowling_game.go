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
    StandingPins int
	AvailableRolls int
}

func (b *BowlingGame) Roll(pins int) error {
    if err := b.assertValidRoll(pins); err != nil {
        return err
    }

	b.CurrentFrame.AvailableRolls -= 1

	if pins == 10 {
		b.CurrentFrame.AvailableRolls = 0
	}

    b.CurrentFrame.StandingPins -= pins
	b.score += pins

	return nil
}

func (b *BowlingGame) Score() int {
	return b.score
}

func NewBowlingGame() BowlingGame {
	return BowlingGame{
		CurrentFrame: Frame{
            StandingPins: 10,
			AvailableRolls: 2,
		},
	}
}

func (b *BowlingGame) assertValidRoll(pins int) error {
	if b.CurrentFrame.AvailableRolls == 0 {
		return fmt.Errorf("can't roll more times")
	}

	if pins > 10 {
		return fmt.Errorf("more than 10 pins are not allowed in a single roll")
	}

    if b.CurrentFrame.StandingPins < pins {
        return fmt.Errorf("more than 10 pins are not allowed in the same frame")
    }

    return nil
}
