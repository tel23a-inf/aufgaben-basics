package suchen

// Erwartet eine Liste von Zahlen und eine Zahl x.
// Liefert die Position von x in der Liste.
// Liefert -1, falls x nicht enthalten ist.
func Find(list []int, x int) int {
	// Verwenden Sie eine for-Schleife, um die Elemente der Liste zu durchlaufen.
	// Prüfen Sie in jedem Schleifendurchlauf, ob das aktuelle Element gleich x ist.
	// Wenn ja, geben Sie die Position des Elements zurück.
	for pos, el := range list {
		if el == x {
			return pos
		}
	}
	return -1
}
