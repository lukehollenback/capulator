package ledger

import (
	"os"

	"github.com/gocarina/gocsv"
)

//
// Ledger holds a collection of time-series "buy" and "sell" transactions and their attributes.
//
type Ledger struct {
	Lines []*Line
}

//
// Load creates a ledger from a specified CSV file.
//
func Load(csvPath string) (*Ledger, error) {
	file, err := os.OpenFile(csvPath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	o := &Ledger{
		Lines: []*Line{},
	}

	if err := gocsv.UnmarshalFile(file, &o.Lines); err != nil {
		return nil, err
	}

	return o, nil
}

//
// Save writes out the ledger to the specified CSV file.
//
func (o *Ledger) Save(path string) error {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := gocsv.MarshalFile(&o.Lines, file); err != nil {
		return err
	}

	return nil
}
