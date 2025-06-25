package main

import (
	"github.com/phpdave11/gofpdf"
)

func main() {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 20)
	pdf.Cell(40, 10, "Hello, This is PDF generated in pure Go!")

	pdf.Ln(15)

	// --- 2. First Table ---
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(0, 10, "Main Table")
	pdf.Ln(10)
	pdf.SetFont("Arial", "", 12)

	header := []string{"Name", "Email", "Country"}
	data := [][]string{
		{"Alice", "alice@example.com", "USA"},
		{"Bob", "bob@example.com", "UK"},
		{"Charlie", "charlie@example.com", "Canada"},
	}
	drawTable(pdf, header, data)

	pdf.Ln(10)

	// --- 3. Second Table ---
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(0, 10, "Other Table")
	pdf.Ln(10)
	pdf.SetFont("Arial", "", 12)

	otherHeader := []string{"Product", "Quantity", "Price"}
	otherData := [][]string{
		{"Apples", "10", "$5"},
		{"Bananas", "6", "$2.5"},
		{"Oranges", "8", "$4"},
	}
	drawTable(pdf, otherHeader, otherData)

	// Output
	if err := pdf.OutputFileAndClose("example.pdf"); err != nil {
		panic(err)
	}
}

func drawTable(pdf *gofpdf.Fpdf, header []string, data [][]string) {
	colWidths := make([]float64, len(header))
	for i := range colWidths {
		colWidths[i] = 190.0 / float64(len(header)) // auto-fit width
	}

	// Header
	pdf.SetFont("Arial", "B", 12)
	for i, str := range header {
		pdf.CellFormat(colWidths[i], 10, str, "1", 0, "C", false, 0, "")
	}
	pdf.Ln(-1)

	// Rows
	pdf.SetFont("Arial", "", 12)
	for _, row := range data {
		for i, cell := range row {
			pdf.CellFormat(colWidths[i], 10, cell, "1", 0, "", false, 0, "")
		}
		pdf.Ln(-1)
	}
}
