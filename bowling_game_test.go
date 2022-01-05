package main

import "testing"

func TestBowlingGame(t *testing.T) {
	t.Run(
		"When we start a new game AND we roll once THEN the score is equal to the roll",
		func(t *testing.T) {
			bowlingGame := NewBowlingGame()

			if err := bowlingGame.Roll(1); err != nil {
				t.Fail()
			}

			if bowlingGame.Score() != 1 {
				t.Fail()
			}
		},
	)

	t.Run(
		"When we start a new game AND we roll twice THEN the score is equal to the sum of rolls",
		func(t *testing.T) {
			bowlingGame := NewBowlingGame()

			if err := bowlingGame.Roll(1); err != nil {
				t.Fail()
			}
			if err := bowlingGame.Roll(1); err != nil {
				t.Fail()
			}

			if bowlingGame.Score() != 2 {
				t.Fail()
			}
		},
	)

	t.Run(
		"When we roll AND we roll a number of pin over 10 THEN error",
		func(t *testing.T) {
			bowlingGame := NewBowlingGame()

			if bowlingGame.Roll(11) == nil {
				t.Fail()
			}
		},
	)

	//t.Run(
	//	"GIVEN a game with 1 frame WHEN we roll a strike (10 pins) THEN we can not roll again in the same frame",
	//	func(t *testing.T) {
	//		bowlingGame := NewBowlingGame()
	//
	//		bowlingGame.Roll(10)
	//
	//		err := bowlingGame.Roll(1)
	//		if err == nil {
	//			t.Fail()
	//		}
	//	},
	//)

	t.Run(
		"",
		func(t *testing.T) {
			bowlingGame := NewBowlingGame()

			if err := bowlingGame.Roll(1); err != nil {
				t.Fail()
			}
			if err := bowlingGame.Roll(1); err != nil {
				t.Fail()
			}
			if err := bowlingGame.Roll(1); err != nil {
				t.Fail()
			}

			if bowlingGame.Score() != 3 {
				t.Fail()
			}
		},
	)

	t.Run(
		"When ",
		func(t *testing.T) {
			bowlingGame := NewBowlingGame()

			if err := bowlingGame.Roll(6); err != nil {
				t.Fail()
			}
			if err := bowlingGame.Roll(4); err != nil {
				t.Fail()
			}
			if err := bowlingGame.Roll(2); err != nil {
				t.Fail()
			}

			if bowlingGame.Score() != 14 {
				t.Fail()
			}
		},
	)
}
