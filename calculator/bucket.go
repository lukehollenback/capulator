package calculator

import (
	"log"

	"github.com/lukehollenback/capulator/ledger"
	"github.com/shopspring/decimal"
)

//
// Bucket holds cryptocurrency units that were acquired in a "buy" transaction. It provides easy
// ways to subsequently drain weighted units during "sell" transactions.
//
type Bucket struct {
	CurrentUnits decimal.Decimal
	Line         *ledger.Line
	logger       *log.Logger
}

//
// NewBucket constructs a bucket instance.
//
func NewBucket(line *ledger.Line, logger *log.Logger) *Bucket {
	return &Bucket{
		CurrentUnits: line.Units,
		Line:         line,
		logger:       logger,
	}
}

//
// CalculateCostBasis calculates and returns the weighted cost basis given the regested number of
// units and what this bucket can provide. The bucket is updated accordingly, and the number of
// remaining units that could not be provided by this bucket are also returned.
//
func (o *Bucket) CalculateCostBasis(units decimal.Decimal) (remainingUnits decimal.Decimal, costBasis decimal.Decimal) {
	//
	// Determine how many units this bucket can actually provide.
	//
	var availableUnits decimal.Decimal

	if units.GreaterThan(o.CurrentUnits) {
		availableUnits = o.CurrentUnits
	} else {
		availableUnits = units
	}

	//
	// Determine the cost basis that this bucket can provide.
	//
	multiplier := availableUnits.Div(o.Line.Units)
	costBasis = multiplier.Mul(o.Line.Price().Add(o.Line.Fees))

	o.logger.Printf("requestedUnits: %s, currentUnits: %s, availableUnits: %s, costBasis: %s\n", units, o.CurrentUnits, availableUnits, costBasis)

	//
	// Determine the remaining units that will need to come from other buckets.
	//
	remainingUnits = units.Sub(availableUnits)

	//
	// Update the bucket to no longer have the units that were provided.
	//
	o.CurrentUnits = o.CurrentUnits.Sub(availableUnits)

	return
}
