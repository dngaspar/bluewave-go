package main

import (
	"bluewave/pdf"
	"fmt"
)

func main() {
	result := pdf.PdfToText("test.pdf", 1, 4)
	fmt.Println(result[1])
}
