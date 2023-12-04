package suchen

// Erwartet eine Liste von Zahlen und eine Zahl x.
// Liefert true, falls die Liste x mehr als einmal enthält.
func ContainsMultiple(list []int, x int) bool {
	counter := 0
	// begin:hint
	// Verwenden Sie eine for-Schleife, um die Elemente der Liste zu durchlaufen.
	// Prüfen Sie in jedem Schleifendurchlauf, ob das aktuelle Element gleich x ist.
	// Wenn ja, erhöhen Sie die Variable counter um 1.
	// Wenn der counter größer als 1 wird, geben Sie true zurück (vorzeitiger Abbruch der Funktion).
	// end:hint
	// begin:solution
	for _, el := range list {
		if el == x {
			counter++
			if counter > 1 {
				return true
			}
		}
	}
	// end:solution
	return false
}
