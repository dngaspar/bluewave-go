package main

import (
	"fmt"

	"github.com/gen2brain/go-fitz"
)

func main() {
	// fmt.Println(src.FindCommonSubstrings("abcde", "aaaabcbcd", 2))
	// src.GetFileData("tmp/sample_file_1.pdf", 1, false, "1.6.1")

	// src.ComparePdfsText()
	// src.GetFileData()
	// src.TrainResultClassifiers()
	doc, _ := fitz.New("smaple_files/sample_file_1.pdf")
	fmt.Println(doc)

}
