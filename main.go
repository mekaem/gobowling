package main

import (
	"fmt"
)

type Bowling struct {
	rolls []int
}

func NewBowling() *Bowling {
	return &Bowling{
		rolls: make([]int, 0, 21),
	}
}

func (b *Bowling) roll(pins int) {
	b.rolls = append(b.rolls, pins)
}

func (b *Bowling) score() int {
	score := 0
	frameIndex := 0

	for frame := 0; frame < 10; frame++ {
		if frameIndex >= len(b.rolls) {
			break
		}
		if b.rolls[frameIndex] == 10 {
			score += 10 + b.rolls[frameIndex+1] + b.rolls[frameIndex+2]
			frameIndex++
		} else if b.rolls[frameIndex]+b.rolls[frameIndex+1] == 10 {
			if frameIndex+2 >= len(b.rolls) {
				break
			}
			score += 10 + b.rolls[frameIndex+2]
			frameIndex += 2
		} else {
			score += b.rolls[frameIndex] + b.rolls[frameIndex+1]
			frameIndex += 2
		}
	}
	return score
}

func assertEqual(name string, expected int, actual int) {
	if expected != actual {
		msg := fmt.Sprintf("FAILED => %s", name)
		fmt.Println(msg)
		fmt.Println("----------------------------------")
		fmt.Printf("Expected: %d\n", expected)
		fmt.Printf("Your Answer: %d\n", actual)
		fmt.Println("!!!  FAIL  !!!")
		panic(msg)
	} else {
		fmt.Printf("SUCCESS => %s\n", name)
	}
}

func main() {
	b := &Bowling{}
	b.roll(3)
	b.roll(4)
	assertEqual("one frame", 7, b.score())

	b = &Bowling{}
	for i := 1; i <= 20; i++ {
		b.roll(2)
	}
	assertEqual("all twos", 40, b.score())

	b = &Bowling{}
	b.roll(6)
	b.roll(4)
	b.roll(3)
	b.roll(4)
	assertEqual("one spare", 20, b.score())

	b = &Bowling{}
	b.roll(10)
	b.roll(3)
	b.roll(4)
	assertEqual("one strike", 24, b.score())

	b = &Bowling{}
	b.roll(2)
	b.roll(2)
	b.roll(6)
	b.roll(4)
	b.roll(3)
	b.roll(4)
	assertEqual("second frame spare", 24, b.score())

	b = &Bowling{}
	b.roll(10)
	b.roll(10)
	b.roll(3)
	b.roll(4)
	assertEqual("two strikes", 47, b.score())

	b = &Bowling{}
	b.roll(10)
	b.roll(10)
	b.roll(10)
	b.roll(3)
	b.roll(4)
	b.roll(6)
	b.roll(4)
	b.roll(0)
	b.roll(10)
	b.roll(5)
	b.roll(1)
	assertEqual("this and that", 108, b.score())

	b = &Bowling{}
	b.roll(10)
	b.roll(10)
	b.roll(10)
	b.roll(10)
	b.roll(10)
	b.roll(10)
	b.roll(10)
	b.roll(10)
	b.roll(10)
	b.roll(10)
	b.roll(10)
	b.roll(10)
	assertEqual("perfect game", 300, b.score())

	b = &Bowling{}
	b.roll(10)
	b.roll(10)
	b.roll(10)
	b.roll(10)
	b.roll(10)
	b.roll(10)
	b.roll(10)
	b.roll(10)
	b.roll(10)
	b.roll(5)
	b.roll(5)
	b.roll(5)
	assertEqual("last frame spare", 270, b.score())
}
