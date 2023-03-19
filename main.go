package main

import (
	"fmt"
	"image"
	"image/jpeg"

	// "io/ioutil"
	"os"
	"path/filepath"

	"github.com/gen2brain/go-fitz"
)

func pdfToImage(filePath string, savePath string, fromPageNumber int, toPageNumber int) []image.Image {
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
		// fmt.Printf("img: %s\n", reflect.TypeOf(img))
		f, err := os.Create(filepath.Join(savePath, fmt.Sprintf("page%03d.jpg", n)))
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

func pdfToText(filePath string, savePath string, fromPageNumber int, toPageNumber int) []string {
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
		fmt.Println(text)
		f, err := os.Create(filepath.Join(savePath, fmt.Sprintf("page%03d.text", n)))
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

func pdfToHtml(filePath string, savePath string, fromPageNumber int, toPageNumber int) []string {
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

		f, err := os.Create(filepath.Join(savePath, fmt.Sprintf("test%03d.html", n)))
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
func main() {
	// doc, err := fitz.New("test.pdf")
	// if err != nil {
	// 	panic(err)
	// }
	result := pdfToHtml("test.pdf", "assets/html", 1, 4)
	fmt.Println(result)
}
