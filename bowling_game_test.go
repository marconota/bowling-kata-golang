package main

import "testing"

func TestBowlingGame(t *testing.T) {
	t.Run(
		"When we start a new game AND we roll once THEN the score is equal to the roll",
		func(t *testing.T) {
			bowlingGame := BowlingGame{}

			bowlingGame.Roll(1)

			if bowlingGame.Score() != 1 {
				t.Fail()
			}
		},
	)

	t.Run(
		"When we start a new game AND we roll twice THEN the score is equal to the sum of rolls",
		func(t *testing.T) {
			bowlingGame := BowlingGame{}

			bowlingGame.Roll(1)
			bowlingGame.Roll(1)

			if bowlingGame.Score() != 2 {
				t.Fail()
			}
		},
	)

	t.Run(
		"When we roll AND we roll a number of pin over 10 THEN error",
		func(t *testing.T) {
			bowlingGame := BowlingGame{}

			if bowlingGame.Roll(11) == nil {
				t.Fail()
			}
		},
	)

	t.Run(
		"When we roll a strike (10 pins) THEN we can not roll again in the same frame",
		func(t *testing.T) {
			bowlingGame := BowlingGame{}

			bowlingGame.Roll(10)

			err := bowlingGame.Roll(1)
			if err == nil {
				t.Fail()
			}
		},
	)
}
