package suchen

// Erwartet eine Liste von Zahlen.
// Liefert true, falls die Liste an irgendeiner
// Stelle eine Kette von drei Zahlen enthält,
// deren Summe 42 ist.
func ContainsSum(list []int) bool {
	// begin:hint
	if len(list) < 3 {
		return false
	}
	counter := list[0] + list[1] + list[2]
	// Verwenden Sie eine for-Schleife, um die Elemente der Liste zu durchlaufen.
	// In jedem Schleifendurchlauf addieren Sie das aktuelle Element zur Summe hinzu
	// und ziehen das Element, das drei Positionen vor dem aktuellen Element steht,
	// von der Summe ab.
	// Wenn die Summe 42 ist, geben Sie true zurück.
	// end:hint
	// begin:solution
	for pos, el := range list[3:] {
		if counter == 42 {
			return true
		}
		counter += el
		counter -= list[pos]
	}
	// end:solution
	return counter == 42
}
