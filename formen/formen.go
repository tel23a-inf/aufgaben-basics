package formen

// PrintRow erwartet eine Zahl n und gibt eine Zeile mit n Sternen auf die Konsole aus.
func PrintRow(n int) {
	// Benutzen Sie eine for-Schleife, um n-mal ein Sternchen auszugeben.
	// TODO
}

// PrintSquare erwartet eine Zahl n und gibt ein Quadrat der Seitenlänge n auf die Konsole aus.
// Das Quadrat soll mit Hilfe des Zeichens "*" gezeichnet werden.
func PrintSquare(n int) {
	// Benutzen Sie die Funktion PrintRow() in einer for-Schleife.
	// TODO
}

// PrintEmptyRow erwartet eine Zahl n und gibt eine Zeile auf die Konsole aus,
// die mit einem Stern beginnt und mit einem Stern endet und dazwischen n-2 Leerzeichen enthält.
func PrintEmptyRow(n int) {
	// Benutzen Sie eine for-Schleife, um n-2-mal ein Leerzeichen auszugeben.
	// Geben Sie davor und danach jeweils ein Sternchen aus.
	// TODO
}

// PrintEmptySquare erwartet eine Zahl n und gibt ein Quadrat der Seitenlänge n auf die Konsole aus.
// Der Rand des Quadrat soll mit Hilfe des Zeichens "*" gezeichnet werden.
// Das Innere soll aus Leerzeichen bestehen.
func PrintEmptySquare(n int) {
	// Benutzen Sie die Funktion PrintEmptyRow() in einer for-Schleife.
	// Rufen Sie vorher und nachher jeweils einmal PrintRow() auf.
	// TODO
}

// PrintCustomRow erwartet eine Zahl n sowie zwei Strings border und fill.
// Gibt eine Zeile auf die Konsole aus, die mit border beginnt und endet
// und dazwischen n-2 mal fill enthält.
func PrintCustomRow(n int, border, fill string) {
	// Gehen Sie ähnlich vor wie in PrintEmptyRow().
	// Hier sind Sie aber nicht auf Sterne und Leerzeichen festgelegt,
	// sondern verwenden die Parameter border und fill.
	// Anmerkung: Wenn diese Funktion fertig ist, können Sie PrintEmptyRow()
	// mit PrintCustomRow() implementieren.
	// TODO
}

// PrintCustomSquare erwartet eine Zahl n und gibt ein Quadrat der Seitenlänge n auf die Konsole aus.
// Die Zeichen für Rand und Inneres werden als Parameter angegeben.
func PrintCustomSquare(n int, border, fill string) {
	// Benutzen Sie die Funktion PrintCustomRow() in einer for-Schleife,
	// ähnlich wie in PrintEmptySquare().
	// TODO
}

// Erwartet zwei Zahlen m und n und gibt ein Rechteck
// der Breite m und der Höhe n auf die Konsole aus.
// Das Rechteck soll mit dem Zeichen "*" ausgefüllt sein.
func PrintRectangle(m, n int) {
	// Gehen Sie analog zu PrintSquare() vor.
	// TODO
}

// Erwartet eine Zahl n und gibt ein Dreieck auf die Konsole aus.
// Das Dreieck soll ein halbes n x n-Quadrat sein, das auf der
// linken oberen Seite ausgefüllt ist (siehe Test).
// Das Dreieck soll mit dem Zeichen "*" ausgefüllt sein.
func PrintTriangle(n int) {
	// Gehen Sie ähnlich wie bei PrintSquare() vor.
	// D.h. verwenden Sie eine for-Schleife, die n-mal PrintRow() aufruft.
	// Als Argument für PrintRow() müssen Sie die Anzahl der Sterne
	// übergeben, die in der jeweiligen Zeile ausgegeben werden sollen.
	// Dies sollte die Laufvariable der for-Schleife sein.
	// TODO
}
