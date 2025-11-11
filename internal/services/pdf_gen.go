package services

import (
	"bytes"
	"fmt"

	"codeberg.org/go-pdf/fpdf"
	"github.com/Belixk/10.11.2025/internal/models"
)

func GeneratePDF(reports []*models.CheckLinksResponse) []byte {
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.SetFont("Arial", "B", 16)
	pdf.SetFontSize(22)
	pdf.AddPage()
	pdf.Cell(40, 10, "Links Report") // Заголовок

	pdf.Ln(12)
	pdf.SetFont("Arial", "", 12)

	for _, report := range reports {
		pdf.Cell(40, 10, fmt.Sprintf("Links Num: %d", report.LinksNum))
		pdf.Ln(8)

		for link, status := range report.Links {
			pdf.Cell(40, 10, fmt.Sprintf(" %s: %s", link, status))
			pdf.Ln(8)
		}
		pdf.Ln(4)
	}

	var buf bytes.Buffer
	pdf.Output(&buf)
	return buf.Bytes()
}
