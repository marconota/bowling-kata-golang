package main

type Game interface {
    Roll(int)
    Score() int
}
