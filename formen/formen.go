package formen

import "fmt"

// PrintRow erwartet eine Zahl n und gibt eine Zeile mit n Sternen auf die Konsole aus.
func PrintRow(n int) {
	// begin:hint
	// Benutzen Sie eine for-Schleife, um n-mal ein Sternchen auszugeben.
	// end:hint
	// begin:solution
	for i := 0; i < n; i++ {
		fmt.Print("*")
	}
	fmt.Println()
	// end:solution
}

// PrintSquare erwartet eine Zahl n und gibt ein Quadrat der Seitenlänge n auf die Konsole aus.
// Das Quadrat soll mit Hilfe des Zeichens "*" gezeichnet werden.
func PrintSquare(n int) {
	// begin:hint
	// Benutzen Sie die Funktion PrintRow() in einer for-Schleife.
	// end:hint
	// begin:solution
	for i := 0; i < n; i++ {
		PrintRow(n)
	}
	// end:solution
}

// PrintEmptyRow erwartet eine Zahl n und gibt eine Zeile auf die Konsole aus,
// die mit einem Stern beginnt und mit einem Stern endet und dazwischen n-2 Leerzeichen enthält.
func PrintEmptyRow(n int) {
	// begin:hint
	// Benutzen Sie eine for-Schleife, um n-2-mal ein Leerzeichen auszugeben.
	// Geben Sie davor und danach jeweils ein Sternchen aus.
	// end:hint
	// begin:solution
	fmt.Print("*")
	for i := 0; i < n-2; i++ {
		fmt.Print(" ")
	}
	fmt.Println("*")
	// end:solution
}

// PrintEmptySquare erwartet eine Zahl n und gibt ein Quadrat der Seitenlänge n auf die Konsole aus.
// Der Rand des Quadrat soll mit Hilfe des Zeichens "*" gezeichnet werden.
// Das Innere soll aus Leerzeichen bestehen.
func PrintEmptySquare(n int) {
	// begin:hint
	// Benutzen Sie die Funktion PrintEmptyRow() in einer for-Schleife.
	// Rufen Sie vorher und nachher jeweils einmal PrintRow() auf.
	// end:hint
	// begin:solution
	PrintRow(n)
	for i := 0; i < n-2; i++ {
		PrintEmptyRow(n)
	}
	PrintRow(n)
	// end:solution
}

// PrintCustomRow erwartet eine Zahl n sowie zwei Strings border und fill.
// Gibt eine Zeile auf die Konsole aus, die mit border beginnt und endet
// und dazwischen n-2 mal fill enthält.
func PrintCustomRow(n int, border, fill string) {
	// begin:hint
	// Gehen Sie ähnlich vor wie in PrintEmptyRow().
	// Hier sind Sie aber nicht auf Sterne und Leerzeichen festgelegt,
	// sondern verwenden die Parameter border und fill.
	// Anmerkung: Wenn diese Funktion fertig ist, können Sie PrintEmptyRow()
	// mit PrintCustomRow() implementieren.
	// end:hint
	// begin:solution
	fmt.Print(border)
	for i := 0; i < n-2; i++ {
		fmt.Print(fill)
	}
	fmt.Println(border)
	// end:solution
}

// PrintCustomSquare erwartet eine Zahl n und gibt ein Quadrat der Seitenlänge n auf die Konsole aus.
// Die Zeichen für Rand und Inneres werden als Parameter angegeben.
func PrintCustomSquare(n int, border, fill string) {
	// begin:hint
	// Benutzen Sie die Funktion PrintCustomRow() in einer for-Schleife,
	// ähnlich wie in PrintEmptySquare().
	// end:hint
	// begin:solution
	PrintCustomRow(n, border, border)
	for i := 0; i < n-2; i++ {
		PrintCustomRow(n, border, fill)
	}
	PrintCustomRow(n, border, border)
	// end:solution
}

// Erwartet zwei Zahlen m und n und gibt ein Rechteck
// der Breite m und der Höhe n auf die Konsole aus.
// Das Rechteck soll mit dem Zeichen "*" ausgefüllt sein.
func PrintRectangle(m, n int) {
	// begin:hint
	// Gehen Sie analog zu PrintSquare() vor.
	// end:hint
	// begin:solution
	for i := 0; i < n; i++ {
		PrintRow(m)
	}
	// end:solution
}

// Erwartet eine Zahl n und gibt ein Dreieck auf die Konsole aus.
// Das Dreieck soll ein halbes n x n-Quadrat sein, das auf der
// linken oberen Seite ausgefüllt ist (siehe Test).
// Das Dreieck soll mit dem Zeichen "*" ausgefüllt sein.
func PrintTriangle(n int) {
	// begin:hint
	// Gehen Sie ähnlich wie bei PrintSquare() vor.
	// D.h. verwenden Sie eine for-Schleife, die n-mal PrintRow() aufruft.
	// Als Argument für PrintRow() müssen Sie die Anzahl der Sterne
	// übergeben, die in der jeweiligen Zeile ausgegeben werden sollen.
	// Dies sollte die Laufvariable der for-Schleife sein.
	// end:hint
	// begin:solution
	for i := n; i > 0; i-- {
		PrintRow(i)
	}
	// end:solution
}
