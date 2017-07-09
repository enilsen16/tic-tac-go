package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	empty int = iota
	x
	o
)

type line [3]int

type grid [8]line

func (l *line) check() bool {
	return l[0] != empty && l[0] == l[1] && l[1] == l[2]
}

func (g *grid) check() bool {
	for _, line := range g {
		if line.check() {
			return true
		}
	}
	return false
}

func (g *grid) place(x, y, val int) {
	if g[x][y] != empty {
		fmt.Println("This spot is taken")
	} else {
		g[x][y] = val
	}
}

func start(in io.Reader, out io.Writer) {
	gb := &grid{}
	scanner := bufio.NewScanner(in)
	players := [2]string{"X", "O"}
	player := -1

	for {
		player = (player + 1) % 2
		fmt.Printf("%s: ", players[player])
		scanner.Scan()
		s := strings.Split(scanner.Text(), ",")
		x, _ := strconv.Atoi(s[0])
		y, _ := strconv.Atoi(s[1])

		gb.place(x, y, player+1)
		if gb.check() {
			fmt.Printf("%s is the winner!", players[player])
			return
		}
	}
}

func main() {
	start(os.Stdin, os.Stdout)
}
