package stringfuncs

import (
	"strings"
)

// Erwartet einen string s und zählt, wie oft der Buchstabe 'A' in s vorkommt.
func CountA(s string) int {
	result := 0
	// begin:hint
	// Laufen Sie in einer for-Schleife über alle Buchstaben des Strings.
	// Prüfen Sie in jedem Schleifendurchlauf, ob der aktuelle Buchstabe ein 'A' ist.
	// Wenn ja, erhöhen Sie result um 1.
	// Alternativ können Sie auch die Funktion CountChar() verwenden,
	// die Sie weiter unten implementieren sollen.
	// end:hint
	// begin:solution
	for _, char := range s {
		if char == 'A' {
			result++
		}
	}
	// end:solution
	return result
}

// Erwartet einen string s und einen Buchstaben c.
// Zählt, wie oft c in s vorkommt.
func CountChar(s string, c rune) int {
	result := 0
	// begin:hint
	// Laufen Sie in einer for-Schleife über alle Buchstaben des Strings.
	// Prüfen Sie in jedem Schleifendurchlauf, ob der aktuelle Buchstabe gleich c ist.
	// Wenn ja, erhöhen Sie result um 1.
	// end:hint
	// begin:solution
	for _, char := range s {
		if char == c {
			result++
		}
	}
	// end:solution
	return result
}

// Erwartet einen String s und liefert einen neuen String,
// in dem jeder Buchstabe aus s zweimal hintereinander vorkommt.
func DuplicateChars(s string) string {
	result := ""
	// begin:hint
	// Laufen Sie in einer for-Schleife über alle Buchstaben des Strings.
	// Fügen Sie in jedem Schleifendurchlauf den aktuellen Buchstaben zweimal
	// an den result-String an.
	// end:hint
	// begin:solution
	for _, char := range s {
		result += string(char)
		result += string(char)
	}
	// end:solution
	return result
}

// Erwartet einen String s und liefert s rückwärts zurück.
func Reverse(s string) string {
	result := ""
	// begin:hint
	// Laufen Sie in einer for-Schleife rückwärts über alle Buchstaben des Strings.
	// Fügen Sie in jedem Schleifendurchlauf den aktuellen Buchstaben an den result-String an.
	// end:hint
	// begin:solution
	for i := len(s) - 1; i >= 0; i-- {
		result += string(s[i])
	}
	// end:solution
	return result
}

// Erwartet zwei Strings s1 und s2 und prüft, ob der eine der andere umgedreht ist.
func IsReverse(s1, s2 string) bool {
	// begin:hint
	// Verwenden Sie die Funktion Reverse(), um s2 umzudrehen
	// und prüfen Sie dann, ob s1 und das umgedrehte s2 gleich sind.
	// end:hint
	// begin:solution
	return s1 == Reverse(s2)
	// end:solution
	// iftask: return false
}

// Erwartet einen String s und prüft, ob s ein Palindrom ist.
// Ein Palindrom ist ein String, der vorwärts und rückwärts gelesen gleich ist.
func IsPalindrome(s string) bool {
	// begin:hint
	// Verwenden Sie die Funktion IsReverse(), um zu prüfen, ob s gleich seinem
	// umgedrehten String ist.
	// end:hint
	// begin:solution
	return IsReverse(s, s)
	// end:solution
	// iftask: return false
}

// Erwartet zwei Strings s1 und s2 und prüft, ob die beiden Anagramme voneinander sind.
// Ein Anagramm von s1 ist ein String, der exakt die gleichen Buchstaben wie s1
// enthält, aber nicht unbedingt in der gleichen Reihenfolge.
func IsAnagram(s1, s2 string) bool {
	// begin:hint
	// Lauft in einer for-Schleife über alle Buchstaben von s1.
	// Prüfen Sie in jedem Schleifendurchlauf mittels der Funktion CountChar(),
	// ob der aktuelle Buchstabe in s1 gleich oft vorkommt wie in s2.
	// end:hint
	// begin:solution
	for _, c := range s1 {
		if CountChar(s1, c) != CountChar(s2, c) {
			return false
		}
	}
	// end:solution
	return true
}

// Erwartet zwei Strings s1 und s2 und prüft, ob die beiden Anagramme voneinander sind.
// Ein Anagramm von s1 ist ein String, der exakt die gleichen Buchstaben wie s1
// enthält, aber nicht unbedingt in der gleichen Reihenfolge.
// Diese Funktion soll keinen Unterschied zwischen Groß- und Kleinschreibung machen.
func IsAnagramIgnoreCase(s1, s2 string) bool {
	// begin:hint
	// Verwenden Sie die Funktion strings.ToLower(), um s1 und s2 in Kleinbuchstaben
	// umzuwandeln, bevor Sie die Funktion IsAnagram() aufrufen.
	// end:hint
	// begin:solution
	return IsAnagram(strings.ToLower(s1), strings.ToLower(s2))
	// end:solution
	// iftask: return false
}

// Erwartet einen String s und einen Buchstaben c.
// Prüft, ob c in s vorkommt.
func Contains(s string, c byte) bool {
	// begin:hint
	// Laufen Sie in einer for-Schleife über alle Buchstaben des Strings.
	// Prüfen Sie in jedem Schleifendurchlauf, ob der aktuelle Buchstabe gleich c ist.
	// Wenn ja, geben Sie true zurück (vorzeitiger Abbruch der Funktion).
	// Alternativ können Sie auch CountChar() verwenden.
	// Bemerkung: Im Package strings gibt es eine Funktion strings.Contains(),
	// die genau diese Aufgabe erfüllt.
	// end:hint
	// begin:solution
	for i := range s {
		if s[i] == c {
			return true
		}
	}
	// end:solution
	return false
}

// Erwartet einen String s und einen Buchstaben c.
// Liefert die Position, an der c in s vorkommt.
// Liefert die Länge von s, falls c nicht in s vorkommt.
// Kommt c mehrfach vor, soll die erste Position geliefert werden.
func PositionOf(s string, c byte) int {
	// begin:hint
	// Laufen Sie in einer for-Schleife über alle Buchstaben des Strings.
	// Prüfen Sie in jedem Schleifendurchlauf, ob der aktuelle Buchstabe gleich c ist.
	// Wenn ja, geben Sie die aktuelle Position zurück (vorzeitiger Abbruch der Funktion).
	// end:hint
	// begin:solution
	for i := range s {
		if s[i] == c {
			return i
		}
	}
	// end:solution
	return len(s)
}

// Erwartet zwei Strings s und t und prüft, ob t in s als Teilstring vorkommt.
func ContainsSubstring(s, t string) bool {
	// begin:hint
	// Laufen Sie in einer for-Schleife über alle Positionen in s.
	// Prüfen Sie in jedem Schleifendurchlauf, ob s ab der aktuellen Position mit t beginnt.
	// Verwenden Sie dazu den Slice-Operator, um einen Teilstring aus s zu extrahieren.
	// end:hint
	// begin:solution
	for i := 0; i < len(s)-len(t)+1; i++ {
		if s[i:i+len(t)] == t {
			return true
		}
	}
	// end:solution
	return false
}

// Erwartet einen String und prüft, ob er korrekte Klammerpaare enthält.
// D.h. der Eingabestring enthält Klammern '(' und ')'.
// Die Funktion soll nun prüfen, ob der String für jede öffnende Klammer auch eine
// schließende Klammer enthält.
// Außerdem darf es keine schließenden Klammern geben, für die es nicht vorher eine
// passende öffnende Klammer gegeben hat.
// Die Funktion soll true liefern, falls der String korrekt geklammert ist.
func CheckParentheses(s string) bool {
	counter := 0
	// begin:hint
	// Laufen Sie in einer for-Schleife über alle Buchstaben des Strings.
	// Jedes Mal, wenn Sie eine öffnende Klammer finden, erhöhen Sie den counter um 1.
	// Jedes Mal, wenn Sie eine schließende Klammer finden, verringern Sie den counter um 1.
	// Wenn der counter negativ wird, ist der String nicht korrekt geklammert.
	// Wenn der counter am Ende der Schleife nicht 0 ist, ist der String nicht korrekt geklammert.
	// end:hint
	// begin:solution
	for _, c := range s {
		if c == '(' {
			counter++
		}
		if c == ')' {
			counter--
			if counter < 0 {
				return false
			}
		}
	}
	// end:solution
	return counter == 0
}

// Erwartet zwei Strings s und sep sowie eine Zahl n.
// Liefert einen String, der aus n Kopien von s besteht, die duch sep getrennt werden.
func ConcatN(s, sep string, n int) string {
	result := ""
	// begin:hint
	// Hängen Sie in einer Schleife n-1 mal s und sep an den result-String an.
	// Hängen Sie dann noch ein weiteres Mal s an.
	// end:hint
	// begin:solution
	for i := 0; i < n-1; i++ {
		result += s + sep
	}
	result += s
	// end:solution
	return result
}

// Erwartet zwei strings s1 und s2.
// Liefert einen String, der abwechselnd aus den Buchstaben des einen und des anderen
// Strings besteht.
func Zip(s1, s2 string) string {
	result := ""

	min := len(s1)
	if len(s2) < min {
		min = len(s2)
	}

	// begin:hint
	// Laufen Sie in einer for-Schleife über alle Buchstaben des kürzeren Strings.
	// Fügen Sie in jedem Schleifendurchlauf den aktuellen Buchstaben aus s1 und s2
	// an den result-String an.
	// Am Ende der Schleife müssen Sie noch die restlichen Buchstaben des längeren
	// Strings an den result-String anhängen.
	// end:hint
	// begin:solution
	for i := 0; i < min; i++ {
		result += string(s1[i]) + string(s2[i])
	}
	result += s1[min:]
	result += s2[min:]
	// end:solution
	return result
}
