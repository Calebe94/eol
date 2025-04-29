package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "list":
		handleListCommand()
	case "product":
		handleProductCommand()
	case "cycle":
		handleCycleCommand()
	case "help":
		printUsage()
	default:
		fmt.Printf("Unknown command: %s\n\n", os.Args[1])
		printUsage()
		os.Exit(1)
	}
}

func handleListCommand() {
	fs := flag.NewFlagSet("list", flag.ExitOnError)
	cli := parseCommonFlags(fs)
	fs.Parse(os.Args[2:])

	fmt.Printf("cli.format: %s\n", cli.format)
	products, err := ListProducts()
	if err != nil {
		log.Fatalf("Error fetching products: %v", err)
	}

	output, err := formatOutput(products, cli.format)
	if err != nil {
		log.Fatal(err)
	}

	handleOutput(output, cli.output)
}

func handleProductCommand() {
	fs := flag.NewFlagSet("product", flag.ExitOnError)
	cli := parseCommonFlags(fs)

	err := fs.Parse(os.Args[2:])
	if err != nil {
		log.Fatal(err)
	}

	if fs.NArg() == 0 {
		log.Fatal("Product name required")
	}
	product := fs.Arg(0)

	cycles, err := GetProductCycles(product)
	if err != nil {
		log.Fatalf("Error fetching cycles: %v", err)
	}

	output, err := formatOutput(cycles, cli.format)
	if err != nil {
		log.Fatal(err)
	}

	handleOutput(output, cli.output)
}

func handleCycleCommand() {
	fs := flag.NewFlagSet("cycle", flag.ExitOnError)
	cli := parseCommonFlags(fs)
	fs.Parse(os.Args[2:])

	if fs.NArg() < 2 {
		log.Fatal("Product and cycle name required")
	}
	product := fs.Arg(0)
	cycle := fs.Arg(1)

	details, err := GetCycleDetails(product, cycle)
	if err != nil {
		log.Fatalf("Error fetching cycle details: %v", err)
	}

	output, err := formatOutput(details, cli.format)
	if err != nil {
		log.Fatal(err)
	}

	handleOutput(output, cli.output)
}

func parseCommonFlags(fs *flag.FlagSet) CLI {
	var cli CLI
	fs.StringVar(&cli.format, "format", "table", "Output format (table|json|csv)")
	fs.StringVar(&cli.format, "f", "table", "Output format (shorthand)")
	fs.StringVar(&cli.output, "output", "", "Output file")
	fs.DurationVar(&cli.watch, "watch", 0, "Refresh interval")
	fs.BoolVar(&cli.noColor, "no-color", false, "Disable color output")
	fs.StringVar(&cli.filter, "filter", "", "Filter expression")
	return cli
}
