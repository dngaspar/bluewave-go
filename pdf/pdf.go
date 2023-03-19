package pdf

import (
	"image"

	"github.com/gen2brain/go-fitz"
)

func PdfToImage(filePath string, fromPageNumber int, toPageNumber int) []image.Image {
	doc, err := fitz.New(filePath)
	if err != nil {
		panic(err)
	}
	// len := toPageNumber - fromPageNumber
	var result []image.Image
	defer doc.Close()
	for n := fromPageNumber - 1; n < toPageNumber; n++ {
		img, err := doc.Image(n)
		if err != nil {
			panic(err)
		}
		result = append(result, img)
		// result[n] = img
	}
	// fmt.Println(result)
	return result
}

func PdfToText(filePath string, fromPageNumber int, toPageNumber int) []string {
	doc, err := fitz.New(filePath)
	if err != nil {
		panic(err)
	}
	var result []string
	defer doc.Close()
	for n := fromPageNumber - 1; n < toPageNumber; n++ {
		text, err := doc.Text(n)
		if err != nil {
			panic(err)
		}
		result = append(result, text)
	}
	return result
}

func PdfToHtml(filePath string, fromPageNumber int, toPageNumber int) []string {
	doc, err := fitz.New(filePath)
	if err != nil {
		panic(err)
	}
	var result []string
	defer doc.Close()
	for n := fromPageNumber - 1; n < toPageNumber; n++ {
		html, err := doc.HTML(n, true)
		// fmt.Printf("html: %s\n", reflect.TypeOf(html))
		// fmt.Println()
		if err != nil {
			panic(err)
		}

		result = append(result, html)
	}
	return result
}
