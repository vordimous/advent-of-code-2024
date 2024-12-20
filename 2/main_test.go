package main

import (
	"reflect"
	"testing"
)

var m [][]int

func init() {
	m = parseInput()
}

func Benchmark_partOne(b *testing.B) {
	type args struct {
		m [][]int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "run",
			args: args{
				m,
			},
		},
	}
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				partOne(tt.args.m)
			}
		})
	}
}

func Test_removeIndex(t *testing.T) {
	type args struct {
		s []int
		i int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Remove first",
			args: args{
				s: []int{1,2,3},
				i: 0,
			},
			want: []int{2,3},
		},
		{
			name: "Remove middle",
			args: args{
				s: []int{1,2,3},
				i: 1,
			},
			want: []int{1,3},
		},
		{
			name: "Remove last",
			args: args{
				s: []int{1,2,3},
				i: 2,
			},
			want: []int{1,2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeIndex(tt.args.s, tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}


func Benchmark_partTwo(b *testing.B) {
	type args struct {
		m [][]int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "run",
			args: args{
				m,
			},
		},
	}
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				partTwo(tt.args.m)
			}
		})
	}
}
