package main

import (
	"testing"
)

func Test_parseInput(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "succeeds",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseInput(); got == "" {
				t.Errorf("parseInput() = %v", got)
			}
		})
	}
}

func Test_findAllValidMuls(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		if got :=  findAllValidMuls(""); len(got) > 0 {
			t.Errorf("findValidMul() = %v", got)
		}
	})	
	t.Run("Regex for muls", func(t *testing.T) {
		if got :=  findAllValidMuls("do()mul(1,2)"); 
			len(got) != 2 && got[0][0] != "do()" && got[1][0] != "mul(1,2)" {
			t.Errorf("findValidMul() = %v", got)
		}
	})
}
