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

	t.Run(
		"GIVEN a spare is scored WHEN the next frame is rolled THEN the bonus is the score of the next roll",
		func(t *testing.T) {
			bowlingGame := NewBowlingGame()

			firstRoll := 6
			secondRoll := 4

			if err := bowlingGame.Roll(firstRoll); err != nil {
				t.Fail()
			}
			if err := bowlingGame.Roll(secondRoll); err != nil {
				t.Fail()
			}

			firstRollAfterSpare := 2
			secondRollAfterSpare := 3
			if err := bowlingGame.Roll(firstRollAfterSpare); err != nil {
				t.Fail()
			}
			if err := bowlingGame.Roll(secondRollAfterSpare); err != nil {
				t.Fail()
			}

			expectedBonus := firstRollAfterSpare

			totalScore := firstRoll + secondRoll + firstRollAfterSpare + secondRollAfterSpare + expectedBonus
			if bowlingGame.Score() != totalScore {
				t.Errorf("bowling score: %d", bowlingGame.Score())
			}
		},
	)

	t.Run(
		"GIVEN a strike is scored WHEN the next frame is rolled THEN the bonus is the score of the next two rolls",
		func(t *testing.T) {
			bowlingGame := NewBowlingGame()

			firstRoll := 10
			if err := bowlingGame.Roll(firstRoll); err != nil {
				t.Fail()
			}

			firstRollAfterStrike := 2
			secondRollAfterStrike := 3
			if err := bowlingGame.Roll(firstRollAfterStrike); err != nil {
				t.Fail()
			}
			if err := bowlingGame.Roll(secondRollAfterStrike); err != nil {
				t.Fail()
			}

			expectedBonus := firstRollAfterStrike + secondRollAfterStrike
			totalScore := firstRoll + firstRollAfterStrike + secondRollAfterStrike + expectedBonus
			if bowlingGame.Score() != totalScore {
				t.Errorf("bowling score: %d", bowlingGame.Score())
			}
		},
	)

	t.Run(
		"GIVEN a strike is scored WHEN the next frame is rolled THEN the bonus is the score of the next two rolls",
		func(t *testing.T) {
			bowlingGame := NewBowlingGame()

			expectedFirstFrameBonus := 0
			firstRoll := 10
			if err := bowlingGame.Roll(firstRoll); err != nil {
				t.Fail()
			}

			firstRollAfterStrike := 10
			if err := bowlingGame.Roll(firstRollAfterStrike); err != nil {
				t.Fail()
			}
			expectedFirstFrameBonus += firstRollAfterStrike

			secondRollAfterStrike := 10
			if err := bowlingGame.Roll(secondRollAfterStrike); err != nil {
				t.Fail()
			}
			expectedSecondFrameBonus := secondRollAfterStrike
			expectedFirstFrameBonus += secondRollAfterStrike

			totalScore := firstRoll + firstRollAfterStrike + secondRollAfterStrike + expectedFirstFrameBonus +
				expectedSecondFrameBonus
			if bowlingGame.Score() != totalScore {
				t.Errorf("bowling score: %d", bowlingGame.Score())
			}
		},
	)
}
