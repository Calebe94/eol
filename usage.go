package main

import (
	"fmt"
	"os"
)

func printUsage() {
	fmt.Fprintf(os.Stderr, `EOL CLI - Product lifecycle information tool

Usage:
  eol [command] [arguments] [flags]

Commands:
  list        List all available products
  product     Show lifecycle details for a product
  cycle       Get details for a specific product cycle
  help        Show this help message

Flags:
  -f, --format string    Output format (table|json|csv) (default "table")
  -o, --output string    Write output to file
  -w, --watch duration   Auto-refresh interval (e.g., 5m, 1h)
  --no-color             Disable color output
  -h, --help             Show help

Examples:
  eol list
  eol product ubuntu --format json
  eol cycle python 3.7 --output cycle.csv
  eol product rhel --watch 30m
`)
}
