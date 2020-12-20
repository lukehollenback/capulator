package main

import (
	"flag"
	"log"

	"github.com/lukehollenback/capulator/calculator"
	"github.com/lukehollenback/capulator/ledger"
)

func main() {
	//
	// Register and parse program arguments.
	//
	inputPath := flag.String("in", "", "Path to CSV ledger file to process.")
	outputPath := flag.String("out", "", "Path to CSV ledger file to write out calculations to.")

	flag.Parse()

	//
	// Load the specified CSV ledger file.
	//
	ledger, err := ledger.Load(*inputPath)
	if err != nil {
		panic(err)
	}

	//
	// Create a logger.
	//
	logger := log.New(log.Writer(), "", log.Ldate|log.Ltime|log.Lmsgprefix)

	//
	// Scroll through the loaded ledger and calculate capital gains for each sale.
	//
	calculator := calculator.NewCalculator(ledger, logger)

	calculator.Calculate()

	//
	// Save an updated CSV ledger file with capital gains included to the specified location.
	//
	if err := ledger.Save(*outputPath); err != nil {
		panic(err)
	}
}
