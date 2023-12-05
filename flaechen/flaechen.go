package flaechen

import "math"

// Erwartet zwei Seitenlängen a und b.
// Liefert die Fläche des entsprechenden Rechtecks.
func AreaRectangle(a, b float64) float64 {
	// Benutzen Sie den *-Operator, um a und b zu multiplizieren.
	return a * b
}

// Erwartet eine Seitenlänge a.
// Liefert die Fläche des entsprechenden Quadrats.
func AreaSquare(a float64) float64 {
	// Benutzen Sie die Funktion AreaRectangle().
	return AreaRectangle(a, a)
}

// Erwartet zwei Seitenlängen a und b.
// Liefert die Fläche eines rechtwinkligen Dreiecks, dessen Katheten a und b sind.
func AreaRightTriangle(a, b float64) float64 {
	// Benutzen Sie die Funktion AreaRectangle() und teilen Sie das Ergebnis durch 2.
	return AreaRectangle(a, b) / 2
}

// Erwartet zwei Seitenlängen a und b.
// Liefert den Umfang des entsprechenden Rechtecks.
func PerimeterRectangle(a, b float64) float64 {
	// Der Umfang eines Rechtecks mit den Seitenlängen a und b ist 2 * (a + b).
	return 2 * (a + b)
}

// Erwartet eine Seitenlänge a.
// Liefert den Umfang des entsprechenden Quadrats.
func PerimeterSquare(a float64) float64 {
	// Benutzen Sie die Funktion PerimeterRectangle().
	return PerimeterRectangle(a, a)
}

// Erwartet die Längen der Katheten eines rechtwinkligen Dreiecks.
// Liefert die Länge der Hypotenuse.
func Hypotenuse(a, b float64) float64 {
	// Benutzen Sie die Funktion math.Sqrt().
	return math.Sqrt(a*a + b*b)
}

// Erwartet zwei Seitenlängen a und b.
// Liefert den Umfang eines rechtwinkligen Dreiecks, dessen Katheten a und b sind.
func PerimeterRightTriangle(a, b float64) float64 {
	// Der Umfang eines rechtwinkligen Dreiecks mit den Katheten a und b ist a + b + Hypotenuse(a, b).
	return a + b + Hypotenuse(a, b)
}
