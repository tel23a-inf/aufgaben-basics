package suchen

// Erwartet eine Liste von Zahlen und eine Zahl x.
// Falls es eine Stelle gibt, an der x zwei Mal
// hintereinander vorkommt, liefert die Funktion
// diese Stelle.
// Liefert -1, falls die Situation nicht auftritt.
func FindPair(list []int, x int) int {
	if len(list) < 2 {
		return -1
	}

	// begin:hint
	// Gehen Sie ähnlich vor wie in der Funktion Find.
	// Ändern Sie lediglich die Bedingung in der if-Anweisung.
	// end:hint
	// begin:solution
	for pos, el := range list[:len(list)-1] {
		if el == x && list[pos+1] == x {
			return pos
		}
	}
	// end:solution
	return -1
}
