package pdf

import (
	"image"

	"github.com/gen2brain/go-fitz"
)

func PdfPageInfo(filePath string, page int) image.Image {
	doc, err := fitz.New(filePath)
	if err != nil {
		panic(err)
	}
	var result image.Image
	doc.Image(page)
	defer doc.Close()
	return result
}
