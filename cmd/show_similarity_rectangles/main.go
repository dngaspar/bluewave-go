package main

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
)

func main() {
	//  Create new parser object
	parser := argparse.NewParser("Input", "Input args")

	// Create string flag
	fileName := parser.String("f", "filename", &argparse.Options{Required: true, Help: "PDF comparison result filename to create thumbnails of"})
	// Parser input
	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Println(parser.Usage(err))
	}
	fmt.Println(*fileName)
}
