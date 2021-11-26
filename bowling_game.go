package main

import "fmt"

type Game interface {
    Roll(int) error
    Score() int
}

type BowlingGame struct {
    score int
}

func (b *BowlingGame) Roll(pins int) error {
    if pins > 10 {
        return fmt.Errorf("more than 10 pins are not allowed")
    }

    b.score += pins
    return nil
}

func (b *BowlingGame) Score() int {
    return b.score
}
