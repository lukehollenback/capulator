package calculator

import (
	"log"

	"github.com/lukehollenback/capulator/ledger"
	"github.com/shopspring/decimal"
)

//
// Calculator holds references to ledger and bucket data that can be used to calculate capital gains
// on sales.
//
type Calculator struct {
	Ledger          *ledger.Ledger
	bucketsBySymbol map[string][]*Bucket
	logger          *log.Logger
}

//
// NewCalculator constructs a calculator instance.
//
func NewCalculator(ledger *ledger.Ledger, logger *log.Logger) *Calculator {
	return &Calculator{
		Ledger:          ledger,
		bucketsBySymbol: make(map[string][]*Bucket),
		logger:          logger,
	}
}

//
// Calculate iterates, in FIFO order, through each line in the calculator's ledger. "Buy"
// transactions are put into buckets, and "sell" transactions pull from said buckets in order
// to determine the appropriate weighted cost basis and ultimataley the associated capital gain (or
// loss).
//
func (o *Calculator) Calculate() {
	for _, line := range o.Ledger.Lines {
		o.logger.Printf("Processing %+v...", line)

		if line.Action == "Buy" {
			buckets := o.getOrInitBuckets(line.Symbol)
			bucket := NewBucket(line, o.logger)
			o.bucketsBySymbol[line.Symbol] = append(buckets, bucket)
		} else if line.Action == "Sell" {
			line.Gains = o.traverseBuckets(line)
		}

		o.logger.Printf("...Done at %+v.", line)
		o.logger.Printf("---")
	}
}

//
// getOrInitBuckets retrieves a slice of ordered buckets for the given symbol.
//
func (o *Calculator) getOrInitBuckets(symbol string) []*Bucket {
	buckets, exists := o.bucketsBySymbol[symbol]

	if !exists {
		o.bucketsBySymbol[symbol] = []*Bucket{}
		buckets = o.bucketsBySymbol[symbol]
	}

	return buckets
}

//
// traverseBuckets calculates the appropriate cost basis for the provided "sell" line by pulling
// weighted units out of previous buckets in FIFO order. Ultimately, the cost basis is used to
// calculate the "sell" line's capital gain and return it.
//
func (o *Calculator) traverseBuckets(line *ledger.Line) decimal.Decimal {
	//
	// Iterate through the relevant (to the line's symbol) buckets to aggregate the appropriate cost
	// basis until all units that were sold on the line are accounted for.
	//
	remainingUnits := line.Units
	runningCostBasis := line.Fees

	for _, bucket := range o.getOrInitBuckets(line.Symbol) {
		var costBasis decimal.Decimal

		remainingUnits, costBasis = bucket.CalculateCostBasis(remainingUnits)

		runningCostBasis = runningCostBasis.Add(costBasis)

		o.logger.Printf("remainingUnits: %s, runningCostBasis: %s\n", remainingUnits, runningCostBasis)

		if remainingUnits.Equal(decimal.Zero) {
			break
		}
	}

	//
	// Assert that we are not in a bad state.
	//
	if remainingUnits.GreaterThan(decimal.Zero) {
		o.logger.Fatalf("Remaining units for \"sell\" greater than available units from all previous \"buys\". How was more sold then bought?")
	}

	//
	// Return the calculated capital gain.
	//
	return line.Price().Sub(runningCostBasis)
}
