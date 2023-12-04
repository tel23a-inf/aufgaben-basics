package flaechen

import "math"

// Erwartet zwei Seitenlängen a und b.
// Liefert die Fläche des entsprechenden Rechtecks.
func AreaRectangle(a, b float64) float64 {
	// begin:hint
	// Benutzen Sie den *-Operator, um a und b zu multiplizieren.
	// end:hint
	// begin:solution
	return a * b
	// end:solution
	// iftask: return 0
}

// Erwartet eine Seitenlänge a.
// Liefert die Fläche des entsprechenden Quadrats.
func AreaSquare(a float64) float64 {
	// begin:hint
	// Benutzen Sie die Funktion AreaRectangle().
	// end:hint
	// begin:solution
	return AreaRectangle(a, a)
	// end:solution
	// iftask: return 0
}

// Erwartet zwei Seitenlängen a und b.
// Liefert die Fläche eines rechtwinkligen Dreiecks, dessen Katheten a und b sind.
func AreaRightTriangle(a, b float64) float64 {
	// begin:hint
	// Benutzen Sie die Funktion AreaRectangle() und teilen Sie das Ergebnis durch 2.
	// end:hint
	// begin:solution
	return AreaRectangle(a, b) / 2
	// end:solution
	// iftask: return 0
}

// Erwartet zwei Seitenlängen a und b.
// Liefert den Umfang des entsprechenden Rechtecks.
func PerimeterRectangle(a, b float64) float64 {
	// begin:hint
	// Der Umfang eines Rechtecks mit den Seitenlängen a und b ist 2 * (a + b).
	// end:hint
	// begin:solution
	return 2 * (a + b)
	// end:solution
	// iftask: return 0
}

// Erwartet eine Seitenlänge a.
// Liefert den Umfang des entsprechenden Quadrats.
func PerimeterSquare(a float64) float64 {
	// begin:hint
	// Benutzen Sie die Funktion PerimeterRectangle().
	// end:hint
	// begin:solution
	return PerimeterRectangle(a, a)
	// end:solution
	// iftask: return 0
}

// Erwartet die Längen der Katheten eines rechtwinkligen Dreiecks.
// Liefert die Länge der Hypotenuse.
func Hypotenuse(a, b float64) float64 {
	// begin:hint
	// Benutzen Sie die Funktion math.Sqrt().
	// end:hint
	// begin:solution
	return math.Sqrt(a*a + b*b)
	// end:solution
	// iftask: return 0
}

// Erwartet zwei Seitenlängen a und b.
// Liefert den Umfang eines rechtwinkligen Dreiecks, dessen Katheten a und b sind.
func PerimeterRightTriangle(a, b float64) float64 {
	// begin:hint
	// Der Umfang eines rechtwinkligen Dreiecks mit den Katheten a und b ist a + b + Hypotenuse(a, b).
	// end:hint
	// begin:solution
	return a + b + Hypotenuse(a, b)
	// end:solution
	// iftask: return 0
}
