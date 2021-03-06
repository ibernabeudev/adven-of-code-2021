package main

import (
	"testing"
)

func TestCountDepthMeasurenmentIncrements(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  int
	}{
		"simple": {
			input: []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263},
			want:  7,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := countDepthMeasurenmentIncrements(tc.input)
			if got != tc.want {
				t.Fatalf("%s: expected: %v, got: %v", name, tc.want, got)
			}
		})
	}
}

func TestCountDepthMeasurenmentIncrementsWindow(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  int
	}{
		"complex": {
			input: []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263},
			want:  5,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := countDepthMeasurenmentIncrementsWindow(tc.input)
			if got != tc.want {
				t.Fatalf("%s: expected: %v, got: %v", name, tc.want, got)
			}
		})
	}
}
