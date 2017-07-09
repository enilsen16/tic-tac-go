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
	empty int8 = iota
	x
	o
)

type line [3]int8

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

func (g *grid) place(x, y, val int8) {
	if g[x][y] != empty {
		fmt.Println("This spot is taken")
	} else {
		g[x][y] = val
		g[x+3][y] = val
		if x == 1 && y == 1 {
			g[6][y] = val
			g[7][y] = val
		} else if x == y {
			g[6][y] = val
		} else if x%2 == 0 && y%2 == 0 {
			g[7][y] = val
		}
	}
}

func start(in io.Reader, out io.Writer) {
	gb := &grid{}
	scanner := bufio.NewScanner(in)
	players := [2]string{"X", "O"}
	var player int8 = -1

	for {
		player = (player + 1) % 2
		fmt.Printf("%s: ", players[player])
		scanner.Scan()
		s := strings.Split(scanner.Text(), ",")
		x, _ := strconv.ParseInt(s[0], 10, 8)
		y, _ := strconv.ParseInt(s[1], 10, 8)
		x8, y8 := int8(x), int8(y)

		gb.place(x8, y8, player+1)
		if gb.check() {
			fmt.Printf("%s is the winner!", players[player])
			return
		}
	}
}

func main() {
	start(os.Stdin, os.Stdout)
}
