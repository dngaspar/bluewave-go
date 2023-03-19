package main

import (
	"bluewave/pdfconverter"
	"fmt"
)

func main() {
	// doc, err := fitz.New("test.pdf")
	// if err != nil {
	// 	panic(err)
	// }
	result := pdfconverter.PdfToHtml("test.pdf", "assets/html", 1, 4)
	fmt.Println(result)
}
