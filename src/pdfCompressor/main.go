package main

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
	"github.com/jung-kurt/gofpdf/contrib/gofpdi"
)

func main() {
	pdf := gofpdf.New("P", "pt", "A4", "")

	template1 := gofpdi.ImportPage(pdf, "test.pdf", 1, "/CropBox")
	sizes := gofpdi.GetPageSizes()
	pageOneSizes := sizes[1]
	mediaBoxSize := pageOneSizes["/MediaBox"]

	fmt.Printf("%f", mediaBoxSize["w"])
	pdf.AddPage()

	gofpdi.UseImportedTemplate(pdf, template1, 0, 0, mediaBoxSize["w"], mediaBoxSize["h"])
	err := pdf.OutputFileAndClose("result.pdf")
	if err != nil {
		panic(err)
	}
}
