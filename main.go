package main

const (
	empty int = iota
	x
	o
)

type line struct {
	first, second, third int
}

func (l *line) checkLine() bool {
	return l.first != empty && l.first == l.second && l.second == l.third
}

func main() {
	// grid := [8]line{}

}
