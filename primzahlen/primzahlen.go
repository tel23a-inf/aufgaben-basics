package primzahlen

import "fmt"

// Erwartet zwei Zahlen m und n.
// Liefert true, falls m ein Teiler von n ist.
func Divides(m, n int) bool {
	// Hinweis: Es ist eine Lösung vorgegeben, die man auch in der Praxis verwenden würde.
	// D.h. eigentlich würde man diese Funktion nicht schreiben
	// oder sie so einfach lassen, wie hier in der Vorgabe.
	//
	// Als Übungsaufgabe ersetzen Sie diese Lösung dennoch durch eine,
	// die den Modulo-Operator nicht verwendet.

	// Verwenden Sie eine for-Schleife, um n schrittweise um m zu verringern.
	// Wenn n dabei auf 0 kommt, ist m ein Teiler von n.
	// Prüfen Sie vorher, ob m und n beide positiv sind und ändern Sie sie ggf. das Vorzeichen.
	if m == 0 {
		return false
	}
	if n < 0 {
		n = -n
	}
	if m < 0 {
		m = -m
	}
	for n >= m {
		n -= m
	}
	return n == 0
}

// Erwartet eine Zahl n.
// Liefert true, falls n eine Primzahl ist.
func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	// Verwenden Sie eine for-Schleife, um alle Zahlen von 2 bis n-1 zu durchlaufen.
	// Wenn eine dieser Zahlen n teilt, ist n keine Primzahl.
	for i := 2; i < n; i++ {
		if Divides(i, n) {
			return false
		}
	}
	return true
}

// Erwartet eine Zahl n.
// Gibt alle Primzahlen auf der Konsole aus, die kleiner als n sind.
func PrintPrimes(n int) {
	// Verwenden Sie eine for-Schleife, um alle Zahlen von 2 bis n-1 zu durchlaufen.
	// Prüfen Sie in jedem Schleifendurchlauf, ob die aktuelle Zahl eine Primzahl ist.
	// Wenn ja, geben Sie sie aus.
	for i := 2; i < n; i++ {
		if IsPrime(i) {
			fmt.Println(i)
		}
	}
}

// Erwartet eine Zahl n.
// Liefert die nächstgrößere Primzahl.
// Liefert n, falls n selbst eine Primzahl ist.
func NextPrime(n int) int {
	result := n
	// Verwenden Sie eine for-Schleife, um result so lange zu erhöhen, bis es eine Primzahl ist.
	// Verwenden Sie die Funktion IsPrime, um zu prüfen, ob result eine Primzahl ist.
	for ; !IsPrime(result); result++ {
	}
	return result
}

// Erwartet eine Zahl n.
// Liefert den nächsten Primzahlzwilling, der größer ist als n.
// Genauer: Liefert die kleinste Zahl k, so dass:
// * k >= n
// * k ist eine Primzahl
// * k + 2 ist eine Primzahl
func NextPrimeTwin(n int) int {
	// begin:hint
	// Verwenden Sie eine for-Schleife, um n schrittweise zu erhöhen.
	// Prüfen Sie in jedem Schleifendurchlauf, ob n und n+2 beide Primzahlen sind.
	for n = NextPrime(n); !IsPrime(n + 2); n = NextPrime(n + 1) {
	}
	return n
}

// Erwartet eine Zahl n.
// Liefert die größte Primzahl, die echt kleiner als n ist.
// Falls es keine solche Zahl gibt, wird 0 geliefert.
func GreatestPrimeBelow(n int) int {
	if n <= 2 {
		return 0
	}
	// begin:hint
	// Verwenden Sie eine for-Schleife, um n schrittweise zu verringern,
	// bis es eine Primzahl ist.
	for ; !IsPrime(n); n-- {
	}
	return n
}
