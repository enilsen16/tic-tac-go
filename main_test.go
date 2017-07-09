package main

import "testing"

func TestLine(t *testing.T) {
	tests := []struct {
		line           *line
		expectedResult bool
	}{
		{&line{empty, empty, empty}, false},
		{&line{o, x, x}, false},
		{&line{x, x, x}, true},
	}

	for _, tt := range tests {
		if tt.line.checkLine() != tt.expectedResult {
			t.Fatalf("Nope")
			return
		}
	}
}

// func TestInput(t *testing.T) {
// 	tests := []line{
// 		{x, null, x},
// 		{o, x, o},
// 		{x, null, o},
// 		{x, o, x},
// 		{null, x, null},
// 		{x, o, o},
// 		{x, x, o},
// 		{x, x, x},
// 	}
// }
