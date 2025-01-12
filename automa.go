//alessandro_cacciolo_molica_31254A

package main

import (
	"fmt"
	"strings"
)

// se in {x, y} è presente un ostacolo o il nome dell'automa è già presente non fa nulla
// altrimenti crea un nuovo automa con nome e la posizione
func (p piano) automa(x, y int, nome string) {
	if p.stato(x, y) == "O" {
		return
	}
	p.automi[nome] = Punto{x, y}
}

// posizioni stampa le posizioni di tutti gli automi che hanno il prefisso alpha
func (p piano) posizioni(alpha string) {
	fmt.Println("(")
	for k, v := range p.automi {
		if strings.HasPrefix(k, alpha) {
			fmt.Printf("%s: %d,%d\n", k, v.x, v.y)
		}
	}
	fmt.Print(")")
}
