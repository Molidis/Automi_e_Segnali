//alessandro_cacciolo_molica_31254A

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func esegui(p piano, s string) piano {
	words := strings.Split(s, " ")
	key := words[0]
	switch key {
	case "s":
		x, _ := strconv.Atoi(words[1])
		y, _ := strconv.Atoi(words[2])
		fmt.Println(p.stato(x, y))
	case "S":
		p.stampa()
	case "a":
		a, _ := strconv.Atoi(words[1])
		b, _ := strconv.Atoi(words[2])
		w := words[3]
		p.automa(a, b, w)
	case "o":
		a, _ := strconv.Atoi(words[1])
		b, _ := strconv.Atoi(words[2])
		c, _ := strconv.Atoi(words[3])
		w, _ := strconv.Atoi(words[4])
		p.ostacolo(a, b, c, w)
	case "r":
		a, _ := strconv.Atoi(words[1])
		b, _ := strconv.Atoi(words[2])
		w := words[3]
		p.richiamo(a, b, w)
	case "p":
		w := words[1]
		p.posizioni(w)
	case "e":
		a, _ := strconv.Atoi(words[1])
		b, _ := strconv.Atoi(words[2])
		w := words[3]
		p.esistePercorso(a, b, w)
	case "f":
		os.Exit(0)
	}
	return p
}

func main() {
	var p piano
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	for {
		scanner.Scan()
		line := scanner.Text()
		if line == "c" {
			p = newPiano()
		} else {
			esegui(p, line)
		}
	}
}
