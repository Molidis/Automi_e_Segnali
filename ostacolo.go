//alessandro_cacciolo_molica_31254A

package main

// ostacolo crea un ostacolo nel piano cartesiano, se nello spazio in cui dovrebbe essere agggiunto l'ostacolo è presente un automa non fa nulla
func (p piano) ostacolo(a, b, c, d int) {
	ost := Ostacolo{a, b, c, d}
	if p.contieneAutoma(ost) {
		return
	}
	p.ostacoli.insert(ost)
}

// contieneAutoma restituisce true se lo spazio che dovrebbe occupare l'ostacolo contiene un automa, false altrimenti
func (p piano) contieneAutoma(ost Ostacolo) bool {
	for _, automa := range p.automi {
		if ost.x0 <= automa.x && automa.x <= ost.x1 && ost.y0 <= automa.y && automa.y <= ost.y1 {
			return true
		}
	}
	return false
}
