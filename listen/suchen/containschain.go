package suchen

// Erwartet eine Liste von Zahlen.
// Liefert true, falls die Liste an irgendeiner
// Stelle eine aufsteigende Kette von mindestens
// drei aufeinanderfolgenden Zahlen enthält.
func ContainsChain(list []int) bool {
	if len(list) < 3 {
		return false
	}
	counter := 1
	// Verwenden Sie eine for-Schleife, um die Elemente der Liste zu durchlaufen.
	// Prüfen Sie in jedem Schleifendurchlauf, ob das aktuelle Element gleich dem
	// vorherigen Element plus 1 ist.
	// Wenn ja, erhöhen Sie die Variable counter um 1.
	// Wenn nein, setzen Sie counter auf 1 zurück, da die Kette unterbrochen ist.
	// Wenn dabei counter >= 3 wird, geben Sie true zurück.
	// Wenn Sie die Schleife durchlaufen haben, ohne eine Kette gefunden zu haben,
	// geben Sie false zurück.
	// TODO
	_ = counter // Entfernen Sie diese Zeile, wenn Sie die Variable counter verwenden.
	return false
}
