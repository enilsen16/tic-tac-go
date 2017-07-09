package main

import "fmt"

const (
	empty int = iota
	x
	o
)

type line struct {
	first, second, third int
}

type grid [8]line

func (l *line) check() bool {
	return l.first != empty && l.first == l.second && l.second == l.third
}

func (g *grid) check() (bool, string) {
	for _, line := range g {
		if line.check() {
			return true, fmt.Sprintf("%d", line.first)
		}
	}
	return false, ""
}

func main() {
	// grid := [8]line{}

}
