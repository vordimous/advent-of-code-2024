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
