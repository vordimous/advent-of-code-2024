package main

import (
	"fmt"
)

type game interface {
	rollBall(pins int)
	getScore() int
}

type Bowling struct {
	scoredFrames [][]int
	currentFrameIndex int
	game
}

func (b *Bowling) rollBall(pins int) {

	if b.currentFrameIndex > 9 {
		panic("don't hit the board please")
	}

	// write to the score
	if b.scoredFrames[b.currentFrameIndex] == nil {
		b.scoredFrames[b.currentFrameIndex] = []int{}
	}

	b.scoredFrames[b.currentFrameIndex] = append(b.scoredFrames[b.currentFrameIndex], pins)

	// itterate to next frame
	// if ball score length is 2
	if len(b.scoredFrames[b.currentFrameIndex]) > 1 {
		b.currentFrameIndex += 1
	}
	// if ball score is 10
	// last frame logic
}

func (b *Bowling) getScore() int {
	// calculate score
	t := 0
	for i, f := range b.scoredFrames {
		for x, ball := range f {
			if i > 0 && i < len(b.scoredFrames){
				prev := b.scoredFrames[i - 1]
				if len(prev) == 1 {
					// strike
				} else {
					sumPrev := 0
					for _, s := range prev {
						sumPrev += s
					}
					if sumPrev == 10 && x == 0 {
						// add first
						t += ball
					}
				}
			}
			t += ball
		}
	}
	return t
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func doWork(s string) {
	fmt.Println(s)

}

func main() {
	doWork("Kyruus tech interview")
}
