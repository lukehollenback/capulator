package ledger

import (
	"time"

	"github.com/shopspring/decimal"
)

//
// Line represents a single transaction in a ledger.
//
type Line struct {
	Date   time.Time       `csv:"Date"`
	Action string          `csv:"Action"`
	Symbol string          `csv:"Symbol"`
	PPU    decimal.Decimal `csv:"Price per Unit"`
	Units  decimal.Decimal `csv:"Units"`
	Fees   decimal.Decimal `csv:"Fees"`
	Gains  decimal.Decimal `csv:"Capital Gain"`
}

//
// Price returns the product of the line's price-per-unit and its number of units.
//
func (o *Line) Price() decimal.Decimal {
	return o.PPU.Mul(o.Units)
}
