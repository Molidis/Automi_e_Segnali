//alessandro_cacciolo_molica_31254A

package main

import (
	"container/heap"
	"fmt"
	"strings"
)

// newPiano crea un nuovo piano, se esiste già un piano lo sovrascrive
func newPiano() piano {
	return crea()
}

func crea() piano {
	return piano{
		automi:   make(map[string]Punto),
		ostacoli: &ostacoli{},
	}
}

// stampa lo stato del punto inserito, se il punto ha un automa stampa 'A', se è un ostacolo stampa 'O', altrimenti 'E'
func (p piano) stato(x, y int) string {
	for current := p.ostacoli.head; current != nil; current = current.next {
		if current.ost.x0 <= x && x <= current.ost.x1 && current.ost.y0 <= y && y <= current.ost.y1 {
			return "O"
		}
	}
	for _, pos := range p.automi {
		if pos.x == x && pos.y == y {
			return "A"
		}
	}
	return "E"
}

// stampa l'elenco di tutti gli automi seguito dall'elenco degli ostacoli
func (p piano) stampa() {
	fmt.Println("(")
	for nome, pos := range p.automi {
		fmt.Printf("%s: %d,%d\n", nome, pos.x, pos.y)
	}
	fmt.Println(")")
	fmt.Println("[")
	for current := p.ostacoli.head; current != nil; current = current.next {
		fmt.Printf("(%d,%d)(%d,%d)\n", current.ost.x0, current.ost.y0, current.ost.x1, current.ost.y1)
	}
	fmt.Println("]")
}

func (p piano) esistePercorso(a, b int, w string) {
	if p.findPercorso(a, b, w) {
		fmt.Println("SI")
	} else {
		fmt.Println("NO")
	}
}

// findPercorso restituisce true se esiste un percorso tra il punto{a, b} e l'automa di nome w, false altrimenti
func (p piano) findPercorso(a, b int, w string) bool {

	// controlla se l'automa esiste nella mappa automi, se non esiste restituisce false
	goal, ok := p.automi[w]
	if !ok {
		return false
	}

	// dato che il numero di passi sarebbe 0 e il numero di passi dev'essere maggiore o uguale a 1 per essere un percorso, nel caso la distanza sia 0 restituisce false
	// questo implica che quando verrà fatto un richiamo, un automa già in questo punto non verrà considerato
	if manhattanDistance(Punto{a, b}, goal) == 0 {
		return false
	}

	// controlla se il punto{a, b} è un ostacolo, se lo è restituisce false
	if p.stato(a, b) == "O" {
		return false
	}

	// inizializza la coda con il punto{a, b}
	coda := &PuntoHeap{}
	heap.Init(coda)
	heap.Push(coda, Punto{a, b})

	// inizializza le mappe mostEfficient, minDaOgniPunto e distStimata
	mostEfficient := make(map[Punto]Punto)
	minDaOgniPunto := make(map[Punto]int)
	minDaOgniPunto[Punto{a, b}] = 0

	// distStimata contiene la distanza stimata tra il punto{a, b} e goal
	distStimata := make(map[Punto]int)
	distStimata[Punto{a, b}] = manhattanDistance(Punto{a, b}, goal)

	// finchè la coda non è vuota
	for coda.Len() > 0 {
		// estrae il punto corrente dalla coda e controlla se è uguale a goal, se lo è restituisce true
		current := heap.Pop(coda).(Punto)
		if current == goal {
			return true
		}

		// per ogni vicino di current controlla se il passo è minore di minDaOgniPunto[vicino] o se minDaOgniPunto[vicino] è uguale a 0
		for _, vicino := range p.vicini(current, goal) {
			passo := minDaOgniPunto[current] + 1

			if passo < minDaOgniPunto[vicino] || minDaOgniPunto[vicino] == 0 {
				// se il passo è minore di minDaOgniPunto[vicino] o minDaOgniPunto[vicino] è uguale a 0 allora aggiorna mostEfficient, minDaOgniPunto e distStimata
				mostEfficient[vicino] = current
				minDaOgniPunto[vicino] = passo
				distStimata[vicino] = minDaOgniPunto[vicino] + manhattanDistance(vicino, goal)
				heap.Push(coda, vicino)
			}
		}
	}
	return false
}

// vicini restituisce i punti vicini (x+1, y) (x-1, y) (x y+1) (x y-1) a current che non sono ostacoli e che non allontanano da goal
func (p piano) vicini(current, goal Punto) []Punto {
	vicini := []Punto{}
	if goal.x > current.x && p.stato(current.x+1, current.y) != "O" {
		vicini = append(vicini, Punto{current.x + 1, current.y})
	}
	if goal.x < current.x && p.stato(current.x-1, current.y) != "O" {
		vicini = append(vicini, Punto{current.x - 1, current.y})
	}
	if goal.y > current.y && p.stato(current.x, current.y+1) != "O" {
		vicini = append(vicini, Punto{current.x, current.y + 1})
	}
	if goal.y < current.y && p.stato(current.x, current.y-1) != "O" {
		vicini = append(vicini, Punto{current.x, current.y - 1})
	}
	return vicini
}

// richiamo sposta gli automi di prefisso alpha dal punto in cui si trovano al punto{a, b} se esiste un percorso minimo tra i due punti che è anche il minore tra tutti i percorsi minimi
func (p piano) richiamo(a, b int, alpha string) {
	// minPerchoso contiene i nomi degli automi che hanno un percorso minimo tra il punto in cui si trovano e il punto{a, b}
	var minPercorso []string

	// min tiene traccia della distanza minima tra il punto{a, b} e l'automa più vicino
	var min = -1

	if p.stato(a, b) == "O" {
		return
	}

	// aggiunge gli automi che hanno il prefisso alpha e hanno un percorso minimo tra il punto in cui si trovano e il punto{a, b}
	for nome := range p.automi {
		if !strings.HasPrefix(nome, alpha) {
			return
		}
		if p.findPercorso(a, b, nome) {
			minPercorso = append(minPercorso, nome)
		}
	}

	// mi permette di trovare min
	for _, nome := range minPercorso {
		if min == -1 {
			min = manhattanDistance(p.automi[nome], Punto{a, b})
		}
		if manhattanDistance(p.automi[nome], Punto{a, b}) < min {
			min = manhattanDistance(p.automi[nome], Punto{a, b})
		}
	}

	// sposta gli automi di minPercorso nel punto{a, b} se la distanza tra il punto{a, b} e l'automa è uguale a min
	for _, nome := range minPercorso {
		if manhattanDistance(p.automi[nome], Punto{a, b}) == min {
			p.automi[nome] = Punto{a, b}
		}
	}
}
