package main

import (
	"aoc/util"
	"testing"
)

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"example1", util.ReadFile("example1.txt"), 2},
		{"example2", util.ReadFile("example2.txt"), 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"example3", util.ReadFile("example3.txt"), 6},
		{"input", util.ReadFile("input.txt"), 10241191004509},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_Part1(b *testing.B) {
	data := util.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		part1(data)
	}
}

func Benchmark_Part2(b *testing.B) {
	data := util.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		part2(data)
	}
}