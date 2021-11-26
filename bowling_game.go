package main

import "fmt"

type Game interface {
    Roll(int) error
    Score() int
}

type BowlingGame struct {
    score int
}

func (b *BowlingGame) Roll(i int) error {
    return fmt.Errorf("error")
}

func (b *BowlingGame) Score() int {
    return 1
}
