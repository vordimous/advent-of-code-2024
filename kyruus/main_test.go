package main

import (
	"testing"
)

var s string

func init() {
	s = "init"
}

func Test_doWork(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test",
			args: args{
				s: s,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doWork(tt.args.s)
		})
	}
}

func Test_scoreOpenFrames(t *testing.T) {
	g := Bowling{}
	frames := []int{1, 2, 3, 4, 5}

	ballScore := 15
	for _, v := range frames {
		g.rollBall(v)
	}

	if g.getScore() != ballScore {
		t.Errorf("Score does match the value, found: %d ", g.getScore())
	}
}

func Test_scoreCompleteNoMarks(t *testing.T) {
	g := Bowling{}
	gameScoreMax := 90

	frames := make([]int, 10)

	for range frames {
		g.rollBall(0)
		g.rollBall(9)
	}

	if g.getScore() != gameScoreMax {
		t.Errorf("Score reach the max score, found: %d ", g.getScore())
	}
}

func Test_scoreCompleteOneSpare(t *testing.T) {
	g := Bowling{
		scoredFrames: make([][]int, 10),
	}
	gameScoreMax := 100

	g.rollBall(1)
	g.rollBall(9)

	for range make([]int, 9) {
		g.rollBall(9)
		g.rollBall(0)
	}

	if g.getScore() != gameScoreMax {
		t.Errorf("Score reach the max score, found: %d ", g.getScore())
	}
}
