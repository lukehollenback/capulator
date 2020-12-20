# Capulator

Capulator is a basic capital gains calculator for cryptocurrency traders. Given a CSV-formatted
ledger file of cryptocurrency purchases and sales, it performs FIFO cost basis and capital gains
calculations. Results are output back to a new CSV-formatted ledger file. Multiple symbols are
supported at once.

## Usage

Usage is documented in the program's help, which can be viewed with the `--help` program argument.

## Ledger File Format

Supported ledger files are basic CSV files. For examples of the expected structure, take a look at
the integration test data in the `tests` package.

## Testing

Both unit tests and integration tests can be run for Capulator. Unit tests for each consumed package
can be run with a simple `go test` command issued from the root of the project. Integration tests
reside in the `tests` directory and package, and can thus be run with a special `go test ./tests`
command issued from the root of the project.

## Disclaimer

Capulator was created solely for the author's own personal use. Use it for real-world calculations
at your own risk – it is not gauranteed to produce accurate results.
