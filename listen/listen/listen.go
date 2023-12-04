package listen

// Erwartet eine Liste und eine Zahl n.
// Liefert zurück, wie oft n in der Liste vorkommt.
func CountElement(list []int, n int) int {
	// ANMERKUNG: Dies ist ein Beispiel, hier ist nichts zu tun.
	result := 0
	for _, element := range list {
		if element == n {
			result++
		}
	}
	return result
}

// Erwartet eine Liste und eine Zahl n.
// Liefert zurück, wie viele Elemente in der Liste größer als n sind.
func CountGreaterThan(list []int, n int) int {
	result := 0
	// begin:hint
	// Verwenden Sie eine for-Schleife, um die Elemente der Liste zu durchlaufen.
	// Prüfen Sie in jedem Schleifendurchlauf, ob das aktuelle Element größer als n ist.
	// Wenn ja, erhöhen Sie die Variable result um 1.
	// end:hint
	// begin:solution
	for _, element := range list {
		if element > n {
			result++
		}
	}
	// end:solution
	return result
}

// Erwartet eine Liste und eine Zahl n.
// Liefert zurück, wie viele Elemente in der Liste ungleich n sind.
func CountNotN(list []int, n int) int {
	result := 0
	// begin:hint
	// Verwenden Sie eine for-Schleife, um die Elemente der Liste zu durchlaufen.
	// Prüfen Sie in jedem Schleifendurchlauf, ob das aktuelle Element ungleich n ist.
	// Wenn ja, erhöhen Sie die Variable result um 1.
	// Anmerkung: Das ist beinahe die gleiche Aufgabe wie CountGreaterThan().
	// end:hint
	// begin:solution
	for _, element := range list {
		if element != n {
			result++
		}
	}
	// end:solution
	return result
}

// Erwartet eine Liste und eine Zahl n.
// Liefert zurück, wie viele Elemente in der Liste durch n teilbar sind.
func CountDivisors(list []int, n int) int {
	result := 0
	// begin:hint
	// Auch diese Aufgabe ist sehr ähnlich zu den beiden vorherigen.
	// Um Teilbarkeit zu prüfen, können Sie den %-Operator ("Modulo") verwenden.
	// end:hint
	// begin:solution
	for _, element := range list {
		if element%n == 0 {
			result++
		}
	}
	// end:solution
	return result
}

// Erwartet zwei Listen.
// Liefert zurück, an wie vielen Stellen die Elemente in den beiden Listen gleich sind.
func CountEqualElements(list1, list2 []int) int {
	// TODO: Diese Lösung funktioniert für gleich lange Listen, aber nicht für
	// unterschiedlich lange Listen. Es gibt zwei Tests zu dieser Funktion,
	// einer ist ok, der andere noch nicht.

	result := 0
	// begin:hint
	// Passen Sie die Bedingung in der for-Schleife an, sodass geprüft wird,
	// ob der aktuelle Index kleiner ist als die Länge der zweiten Liste.
	// Nur dann kann auch auf das Element der zweiten Liste zugegriffen werden.
	// end:hint
	// begin:solution
	for i := range list1 {
		if i < len(list2) && list1[i] == list2[i] {
			result++
		}
	}
	return result
	// end:solution
	/**iftask:
	for i := range list1 {
		if list1[i] == list2[i] {
			result++
		}
	}
	return result
	**/
}

// Erwartet zwei Listen.
// Liefert zurück, an wie vielen Stellen die Elemente in list1 größer
// als in list2 sind. Sind die Listen nicht gleich lang, sollen die Stellen
// ignoriert werden, die nur in einer Liste existieren.
func CountGreaterElements(list1, list2 []int) int {
	result := 0
	// begin:hint
	// Gehen Sie ähnlich vor wie in CountEqualElements().
	// end:hint
	// begin:solution
	for i := range list1 {
		if i < len(list2) && list1[i] > list2[i] {
			result++
		}
	}
	// end:solution
	return result
}
