package suchen

// Erwartet eine Liste von Zahlen.
// Liefert die Position der ersten negativen Zahl in der Liste.
// Liefert -1, falls keine negative Zahl enthalten ist.
func FindNegative(list []int) int {
	// Gehen Sie ähnlich vor wie in der Funktion Find.
	// Ändern Sie lediglich die Bedingung in der if-Anweisung.
	for pos, el := range list {
		if el < 0 {
			return pos
		}
	}
	return -1
}
