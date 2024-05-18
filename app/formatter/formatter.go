package formatter

import (
	"fmt"
	"github.com/go-pdf/fpdf"
	"github.com/misuaaki/godoc/app/parser"
	"github.com/misuaaki/godoc/app/scanner"
	"strings"
)

func GeneratePDF(givenPath string) error {
	s := scanner.NewScanner()
	if err := s.Scan(givenPath, "js"); err != nil {
		return err
	}

	p := parser.NewParser()
	pdf := fpdf.New("P", "mm", "A4", "Arial")

	// Set the document properties
	pdf.SetTitle(p.ProjectInfo.Title, true)
	pdf.SetAuthor(p.ProjectInfo.Description, true)
	pdf.SetSubject(p.ProjectInfo.Version, true)

	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.CellFormat(40, 10, "Project Information", "", 1, "C", false, 0, "fff")

	// Iterate over the packages and files
	for _, pkg := range s.GetPackages() {
		for filePath := range pkg.Files {
			// Parse the file
			err := p.Parse(filePath)
			if err != nil {
				return err
			}

			pdf.AddPage()

			// Write the file info as the table title
			pdf.SetFont("Arial", "B", 16)
			pdf.CellFormat(0, 10, fmt.Sprintf("%s", filePath), "1", 1, "C", false, 0, "")

			// 190mm is the width of the A4 paper
			columnWidth := 190.0 / 7.0

			// Write the function info as the table header
			pdf.SetFont("Arial", "B", 12)
			pdf.CellFormat(columnWidth, 6, "Function", "1", 0, "C", false, 0, "")
			pdf.CellFormat(columnWidth, 6, "Description", "1", 0, "C", false, 0, "")
			pdf.CellFormat(columnWidth, 6, "Parameters", "1", 0, "C", false, 0, "")
			pdf.CellFormat(columnWidth, 6, "Returns", "1", 0, "C", false, 0, "")
			pdf.CellFormat(columnWidth, 6, "Complexity", "1", 0, "C", false, 0, "")
			pdf.CellFormat(columnWidth, 6, "Example", "1", 0, "C", false, 0, "")
			pdf.CellFormat(columnWidth, 6, "Authors", "1", 1, "C", false, 0, "")

			// Write the function info as the table data
			pdf.SetFont("Arial", "", 11)
			for _, f := range p.Functions {
				pdf.CellFormat(columnWidth, 6, f.Name, "1", 0, "L", false, 0, "")
				pdf.CellFormat(columnWidth, 6, "", "1", 0, "L", false, 0, "")
				pdf.CellFormat(columnWidth, 6, "", "1", 0, "L", false, 0, "")
				pdf.CellFormat(columnWidth, 6, "", "1", 0, "L", false, 0, "")
				pdf.CellFormat(columnWidth, 6, f.Complexity, "1", 0, "L", false, 0, "")
				pdf.CellFormat(columnWidth, 6, f.Example, "1", 0, "L", false, 0, "")
				pdf.CellFormat(columnWidth, 6, strings.Trim(fmt.Sprintf("%v", f.Authors), "[]"), "1", 1, "L", false, 0, "")
			}

			// Reset the functions
			p.Functions = nil
		}
	}

	return pdf.OutputFileAndClose("output.pdf")
}
