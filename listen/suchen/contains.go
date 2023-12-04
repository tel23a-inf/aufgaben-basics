package suchen

// Erwartet eine Liste von Zahlen und eine Zahl x.
// Liefert true, falls die Liste die Zahl x enthält.
func Contains(list []int, x int) bool {
	// begin:hint
	// Verwenden Sie eine for-Schleife, um die Elemente der Liste zu durchlaufen.
	// Prüfen Sie in jedem Schleifendurchlauf, ob das aktuelle Element gleich x ist.
	// Wenn ja, geben Sie true zurück (vorzeitiger Abbruch der Funktion).
	// Wenn Sie die Schleife durchlaufen haben, ohne x gefunden zu haben, geben Sie false zurück.
	// end:hint
	// begin:solution
	for _, el := range list {
		if el == x {
			return true
		}
	}
	// end:solution
	return false
}
