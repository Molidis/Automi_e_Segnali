//alessandro_cacciolo_molica_31254A

package main

// Punto rappresenta un punto nel piano cartesiano
type Punto struct {
	x, y int
}

// ostacolo rappresenta un ostacolo nel piano cartesiano, botSx è il punto in basso a sinistra, TopDx è il punto in alto a destra
type Ostacolo struct {
	x0, y0, x1, y1 int
}

// piano rappresenta il piano cartesiano con ostacoli e automi
type piano struct {
	automi   map[string]Punto
	ostacoli *ostacoli
}

// calcola la distanza di Manhattan tra due punti
func manhattanDistance(start, goal Punto) int {
	return intAbs(start.x-goal.x) + intAbs(start.y-goal.y)
}

// intAbs restituisce il valore assoluto di x, nella libreria <math> c'è già il metodo math.Abs ma solo per i float64, invece di svolgere tre cast ho preferito implementare la funzione con int
func intAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type ostacolo struct {
	ost  Ostacolo
	next *ostacolo
}

type ostacoli struct {
	head *ostacolo
}

// inserisce un ostacolo in testa alla lista
func (l *ostacoli) insert(ost Ostacolo) {
	l.head = &ostacolo{ost, l.head}
}

// PuntoHeap rappresenta una coda di punti
type PuntoHeap []Punto

// ritorna la lunghezza della coda
func (h PuntoHeap) Len() int {
	return len(h)
}

// ritorna true se il punto i ha indice maggiore del punto j
func (h PuntoHeap) Less(i, j int) bool {
	return h[i].x < h[j].x || (h[i].x == h[j].x && h[i].y < h[j].y)
}

// scambia h[i] e h[j]
func (h PuntoHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// aggiunge x alla coda
func (h *PuntoHeap) Push(x interface{}) {
	*h = append(*h, x.(Punto))
}

// rimuove e ritorna l'ultimo elemento della coda
func (h *PuntoHeap) Pop() interface{} {
	precedente := *h
	n := len(precedente)
	popped := precedente[n-1]
	*h = precedente[0 : n-1]
	return popped
}
