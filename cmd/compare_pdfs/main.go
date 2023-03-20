package main

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
)

func main() {
	parser := argparse.NewParser("Input", "Input args")

	fileName := parser.String("f", "filenames", &argparse.Options{Required: true, Help: "PDF filenames to compare"})

	method := parser.String("m", "methods", &argparse.Options{Required: true, Help: "Which of the three comparison methods to use: text, digits, images"})

	prettyPrint := parser.String("p", "pretty_print", &argparse.Options{Required: true, Help: "Pretty print output"})

	regenCache := parser.String("c", "regen_cache", &argparse.Options{
		Required: true,
		Help:     "Ignore and overwrite cached data",
	})

	sidecarOnly := parser.String("n", "no_importance", &argparse.Options{
		Required: true,
		Help:     "Do not generate importance scores",
	})

	verbose := parser.String("v", "verbose", &argparse.Options{
		Required: true,
		Help:     "Print things while running",
	})

	version := parser.String("vs", "version", &argparse.Options{
		Required: true,
		Help:     "Print version",
	})

	// Parser input
	err := parser.Parse(os.Args)

	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Println(parser.Usage(err))
	}
	fmt.Println(fileName, method, prettyPrint, regenCache, sidecarOnly, verbose, version)
}
