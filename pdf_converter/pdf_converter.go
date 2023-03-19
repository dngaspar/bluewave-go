package pdf_converter

import (
	"fmt"
	"image"
	"image/jpeg"

	// "io/ioutil"
	"os"
	"path/filepath"

	"github.com/gen2brain/go-fitz"
)

func PdfToImage(filePath string, savePath string, pages []int) []image.Image {
	doc, _ := fitz.New(filePath)
	// len := toPageNumber - fromPageNumber
	var result []image.Image
	defer doc.Close()
	for _, page := range pages {
		img, err := doc.Image(page - 1)
		if err != nil {
			panic(err)
		}
		// fmt.Printf("img: %s\n", reflect.TypeOf(img))
		f, err := os.Create(filepath.Join(savePath, fmt.Sprintf("page%03d.png", page)))
		if err != nil {
			panic(err)
		}

		err = jpeg.Encode(f, img, &jpeg.Options{jpeg.DefaultQuality})
		if err != nil {
			panic(err)
		}

		f.Close()
		result = append(result, img)
		// result[n] = img
	}
	// fmt.Println(result)
	return result
}

func PdfToText(filePath string, savePath string, pages []int) []string {
	doc, _ := fitz.New(filePath)
	var result []string
	defer doc.Close()
	for _, page := range pages {
		text, err := doc.Text(page - 1)
		if err != nil {
			panic(err)
		}
		f, err := os.Create(filepath.Join(savePath, fmt.Sprintf("page%03d.text", page)))
		if err != nil {
			panic(err)
		}
		_, err = f.WriteString(text)
		if err != nil {
			panic(err)
		}
		result = append(result, text)
		f.Close()
	}
	return result
}

func PdfToHtml(filePath string, savePath string, pages []int) []string {
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

		f, err := os.Create(filepath.Join(savePath, fmt.Sprintf("test%03d.html", page)))
		if err != nil {
			panic(err)
		}

		_, err = f.WriteString(html)
		if err != nil {
			panic(err)
		}
		result = append(result, html)
		f.Close()
	}
	return result
}
