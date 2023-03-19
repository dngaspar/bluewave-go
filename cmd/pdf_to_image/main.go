package main

import (
	"bluewave/pdf_converter"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/akamensky/argparse"
)

func removeDuplicates(s []int) []int {
	bucket := make(map[int]bool)
	var result []int
	for _, str := range s {
		if _, ok := bucket[str]; !ok {
			bucket[str] = true
			result = append(result, str)
		}
	}
	return result
}

func getPagelist(pages string) []int {
	segments := strings.Split(pages, ",")
	var pagelist []int
	for _, s := range segments {
		subSegments := strings.Split(s, "-")
		if len(subSegments) == 1 {
			temp, _ := strconv.Atoi(subSegments[0])
			pagelist = append(pagelist, temp)
		} else {
			fromNo, _ := strconv.Atoi(subSegments[0])
			toNo, _ := strconv.Atoi(subSegments[1])
			for ss := fromNo; ss <= toNo; ss++ {
				pagelist = append(pagelist, ss)
			}
		}
		// temp, _ := strconv.Atoi(s)
	}
	pagelist = removeDuplicates(pagelist)
	sort.Ints(pagelist[:])
	return pagelist
}

func main() {
	//  Create new parser object
	parser := argparse.NewParser("Input", "Input args")

	// Create string flag
	fileName := parser.String("f", "filename", &argparse.Options{Required: true, Help: "PDF filename to create thumbnails of"})

	pages := parser.String("p", "pages", &argparse.Options{Required: true, Help: "Pages to create thumbnails of (e.g. '1,2,3' or '3,5-10')"})

	outPath := parser.String("o", "outpath", &argparse.Options{Required: true, Help: "path where to save resulting images"})

	// Parser input
	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Println(parser.Usage(err))
	}
	pagelist := getPagelist(*pages)
	// fmt.Println(*fileName, *outPath, pagelist)
	_ = pdf_converter.PdfToImage(*fileName, *outPath, pagelist)
	// fmt.Println(result)

}
