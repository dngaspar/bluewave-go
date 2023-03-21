package src

import (
	"image"

	"github.com/gen2brain/go-fitz"
)

func GetPageInfo(filePath string, page int) image.Image {
	doc, err := fitz.New(filePath)
	if err != nil {
		panic(err)
	}
	img, err := doc.Image(page)
	if err != nil {
		panic(err)
	}
	defer doc.Close()
	return img
}
