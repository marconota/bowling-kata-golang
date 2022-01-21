# Bowling Kata Golang

## Test-Driven Development code kata in Golang

### Bowling Rules

The game consists of 10 frames. In each frame the player has two rolls to knock down 10 pins. The score for the frame is the total number of pins knocked down, plus bonuses for strikes and spares.

A spare is when the player knocks down all 10 pins in two rolls. The bonus for that frame is the number of pins knocked down by the next roll.

A strike is when the player knocks down all 10 pins on his first roll. The frame is then completed with a single roll. The bonus for that frame is the value of the next two rolls.

In the tenth frame a player who rolls a spare or strike is allowed to roll the extra balls to complete the frame. However no more than three balls can be rolled in tenth frame.

[Here](https://www.bowlinggenius.com/) you can find a visual support of the rules.

### Requirements

Write a struct `Game` that has two methods

`Roll(pins int)` is called each time the player rolls a ball. The argument is the number of pins knocked down.
`Score() int` returns the total score for that game.

### How to write a test in Go

Writing a test is just like writing a function, with a few rules:

* It needs to be in a file with a name like xxx_test.go
* The test function must start with the word Test
* The test function takes one argument only t *testing.T
* In order to use the *testing.T type, you need to import “testing”, like we did with fmt in the other file
* You can do things like t.Fail() when you want to fail.

### TDD cycle in GO

1. Write a test
2. Make the compiler pass
3. Run the test, see that it fails and check the error message is meaningful
4. Write enough code to make the test pass
5. Refactor
