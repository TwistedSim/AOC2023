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
		{"example", util.ReadFile("example.txt"), 4361},
		{"input", util.ReadFile("input.txt"), 509115},
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
		{"example", util.ReadFile("example.txt"), 467835},
		{"input", util.ReadFile("input.txt"), 75220503},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part1alt(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"example", util.ReadFile("example.txt"), 4361},
		{"input", util.ReadFile("input.txt"), 509115},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1alt(tt.input); got != tt.want {
				t.Errorf("part1alt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2alt(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"example", util.ReadFile("example.txt"), 467835},
		{"input", util.ReadFile("input.txt"), 75220503},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2alt(tt.input); got != tt.want {
				t.Errorf("Test_part2alt() = %v, want %v", got, tt.want)
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

func Benchmark_Part1alt(b *testing.B) {
	data := util.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		part1alt(data)
	}
}

func Benchmark_Part2alt(b *testing.B) {
	data := util.ReadFile("input.txt")
	for n := 0; n < b.N; n++ {
		part2alt(data)
	}
}