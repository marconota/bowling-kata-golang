package main

import "testing"

func TestBowlingGame(t *testing.T) {
	t.Run("When we start a new game AND we roll once THEN the score is equal to the roll", func(t *testing.T) {
		bowlingGame := BowlingGame{}

		bowlingGame.Roll(1)

		if bowlingGame.Score() != 1 {
			t.Fail()
		}
    })
}
