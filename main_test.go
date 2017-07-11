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
		if tt.line.check() != tt.expectedResult {
			t.Fatalf("Nope")
			return
		}
	}
}

func TestInput(t *testing.T) {
	test := struct {
		g        *grid
		expected bool
	}{
		&grid{
			{x, empty, x},
			{o, x, o},
			{x, empty, o},
			{x, o, x},
			{empty, x, empty},
			{x, o, o},
			{x, x, o},
			{x, x, x},
		}, true}

	if winner := test.g.check(); winner != test.expected {
		t.Fatalf("Nope")
	}
}

func TestWinningGame(t *testing.T) {
	gb := &grid{}
	gb.place(int8(0), int8(0), 1)
	assertEquals(t, gb.check(), false)

	gb.place(int8(2), int8(0), 2)
	assertEquals(t, gb.check(), false)

	gb.place(int8(0), int8(1), 1)
	assertEquals(t, gb.check(), false)

	gb.place(int8(2), int8(2), 2)
	assertEquals(t, gb.check(), false)

	gb.place(int8(0), int8(2), 1)
	assertEquals(t, gb.check(), true)
}

// Unfortuantely while this is convient, it hides exactly where the error failed
// Leave for now
func assertEquals(t *testing.T, value, expected bool) {
	if value != expected {
		t.Fatalf("Values didn't match")
	}
}
