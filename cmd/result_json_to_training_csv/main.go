package main

import (

	// "bluewave/pdfconverter"
	"bluewave/json_to_csv"
	"fmt"
	"os"

	// "bluewave/"

	"github.com/akamensky/argparse"
)

func main() {

	//  Create new parser object
	parser := argparse.NewParser("Input", "Input args")

	// Create string flag
	fileName := parser.String("f", "filename", &argparse.Options{Required: true, Help: "json file to turn into editable csv (in excel) for manual labeling"})
	// Parser input
	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Println(parser.Usage(err))
	}

	json_to_csv.JsonToCsv(*fileName)
}
