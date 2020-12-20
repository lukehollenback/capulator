package tests

import (
	"log"
	"testing"

	"github.com/lukehollenback/capulator/calculator"
	"github.com/lukehollenback/capulator/ledger"
	"github.com/shopspring/decimal"
)

func TestMultiSymbolGains(t *testing.T) {
	//
	// Load the specified CSV ledger file.
	//
	ledger, err := ledger.Load("./data/ledger-one.csv")
	if err != nil {
		t.Errorf("Failed to load the ledger. (Error: %s)", err)
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
	// Assert that calculated capital gains from "sell" transactions are as expected.
	//
	expected, err := decimal.NewFromString("635")
	if err != nil {
		t.Errorf("Failed to created expected decimal value. (Error: %s)", err)
	}

	if !ledger.Lines[4].Gains.Equal(expected) {
		t.Errorf(
			"Expected ledger line 4 (0-indexed) to have a capital gain of %s, but was instead %s.",
			expected, ledger.Lines[4].Gains,
		)
	}

	expected, err = decimal.NewFromString("238.2")
	if err != nil {
		t.Errorf("Failed to created expected decimal value. (Error: %s)", err)
	}

	if !ledger.Lines[5].Gains.Equal(expected) {
		t.Errorf(
			"Expected ledger line 5 (0-indexed) to have a capital gain of %s, but was instead %s.",
			expected, ledger.Lines[5].Gains,
		)
	}

	expected, err = decimal.NewFromString("49")
	if err != nil {
		t.Errorf("Failed to created expected decimal value. (Error: %s)", err)
	}

	if !ledger.Lines[6].Gains.Equal(expected) {
		t.Errorf(
			"Expected ledger line 6 (0-indexed) to have a capital gain of %s, but was instead %s.",
			expected, ledger.Lines[6].Gains,
		)
	}
}

func TestMultiSymbolLosses(t *testing.T) {
	//
	// Load the specified CSV ledger file.
	//
	ledger, err := ledger.Load("./data/ledger-two.csv")
	if err != nil {
		t.Errorf("Failed to load the ledger. (Error: %s)", err)
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
	// Assert that calculated capital gains from "sell" transactions are as expected.
	//
	expected, err := decimal.NewFromString("-77.5")
	if err != nil {
		t.Errorf("Failed to created expected decimal value. (Error: %s)", err)
	}

	if !ledger.Lines[4].Gains.Equal(expected) {
		t.Errorf(
			"Expected ledger line 4 (0-indexed) to have a capital gain of %s, but was instead %s.",
			expected, ledger.Lines[4].Gains,
		)
	}

	expected, err = decimal.NewFromString("-31.8")
	if err != nil {
		t.Errorf("Failed to created expected decimal value. (Error: %s)", err)
	}

	if !ledger.Lines[5].Gains.Equal(expected) {
		t.Errorf(
			"Expected ledger line 5 (0-indexed) to have a capital gain of %s, but was instead %s.",
			expected, ledger.Lines[5].Gains,
		)
	}

	expected, err = decimal.NewFromString("-98.5")
	if err != nil {
		t.Errorf("Failed to created expected decimal value. (Error: %s)", err)
	}

	if !ledger.Lines[6].Gains.Equal(expected) {
		t.Errorf(
			"Expected ledger line 6 (0-indexed) to have a capital gain of %s, but was instead %s.",
			expected, ledger.Lines[6].Gains,
		)
	}
}
