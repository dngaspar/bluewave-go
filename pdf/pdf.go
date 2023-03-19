package pdf

import (
	"image"

	"github.com/gen2brain/go-fitz"
)

func PdfToImage(filePath string, pages []int) []image.Image {
	doc, err := fitz.New(filePath)
	if err != nil {
		panic(err)
	}
	var result []image.Image
	defer doc.Close()
	for _, page := range pages {
		img, err := doc.Image(page - 1)
		if err != nil {
			panic(err)
		}
		result = append(result, img)
		// result[n] = img
	}
	// fmt.Println(result)
	return result
}

func PdfToText(filePath string, pages []int) []string {
	doc, err := fitz.New(filePath)
	if err != nil {
		panic(err)
	}
	var result []string
	defer doc.Close()
	for _, page := range pages {
		text, err := doc.Text(page - 1)
		if err != nil {
			panic(err)
		}
		result = append(result, text)
	}
	return result
}

func PdfToHtml(filePath string, pages []int) []string {
	doc, err := fitz.New(filePath)
	if err != nil {
		panic(err)
	}
	var result []string
	defer doc.Close()
	for _, page := range pages {
		html, err := doc.HTML(page-1, true)
		if err != nil {
			panic(err)
		}
		result = append(result, html)
	}
	return result
}
