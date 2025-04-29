package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"text/tabwriter"
)

func formatOutput(data interface{}, format string) (string, error) {
	switch format {
	case "json":
		return formatJSON(data)
	case "csv":
		return formatCSV(data)
	case "table":
		return formatTable(data), nil
	default:
		return "", fmt.Errorf("unsupported format: %s", format)
	}
}

func formatJSON(data interface{}) (string, error) {
	var jsonData interface{}

	switch v := data.(type) {
	case []string:
		jsonData = map[string]interface{}{"products": v}
	case []Cycle:
		jsonData = map[string]interface{}{"cycles": v}
	case *Cycle:
		jsonData = v
	default:
		return "", fmt.Errorf("unsupported type for JSON formatting")
	}

	bytes, err := json.MarshalIndent(jsonData, "", "  ")
	if err != nil {
		return "", fmt.Errorf("JSON marshaling failed: %w", err)
	}
	return string(bytes), nil
}

func formatCSV(data interface{}) (string, error) {
	writer := csv.NewWriter(os.Stdout)
	defer writer.Flush()

	var records [][]string

	switch v := data.(type) {
	case []string:
		records = append(records, []string{"Products"})
		for _, product := range v {
			records = append(records, []string{product})
		}
	case []Cycle:
		records = append(records, []string{"Cycle", "Release Date", "EOL", "Latest", "LTS"})
		for _, cycle := range v {
			records = append(records, []string{
				fmt.Sprintf("%v", cycle.Cycle),
				cycle.ReleaseDate,
				fmt.Sprintf("%v", cycle.EOL),
				cycle.Latest,
				fmt.Sprintf("%v", cycle.LTS),
			})
		}
	case *Cycle:
		records = append(records, []string{"Field", "Value"})
		val := reflect.ValueOf(*v)
		typ := val.Type()
		for i := 0; i < val.NumField(); i++ {
			records = append(records, []string{
				typ.Field(i).Name,
				fmt.Sprintf("%v", val.Field(i).Interface()),
			})
		}
	default:
		return "", fmt.Errorf("unsupported data type for CSV")
	}

	var output string
	for _, record := range records {
		line := ""
		for i, field := range record {
			if i > 0 {
				line += ","
			}
			line += field
		}
		output += line + "\n"
	}

	return output, nil
}

func formatTable(data interface{}) string {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	defer w.Flush()

	var output string

	switch v := data.(type) {
	case []string:
		output += "PRODUCTS\n"
		for _, product := range v {
			output += fmt.Sprintf("- %s\n", product)
		}
	case []Cycle:
		output += "CYCLE\tRELEASE DATE\tEOL\tLATEST\tLTS\n"
		for _, cycle := range v {
			output += fmt.Sprintf(
				"%v\t%s\t%v\t%s\t%v\n",
				cycle.Cycle,
				cycle.ReleaseDate,
				cycle.EOL,
				cycle.Latest,
				cycle.LTS,
			)
		}
	case *Cycle:
		output += "FIELD\tVALUE\n"
		val := reflect.ValueOf(*v)
		typ := val.Type()
		for i := 0; i < val.NumField(); i++ {
			output += fmt.Sprintf(
				"%s\t%v\n",
				typ.Field(i).Name,
				val.Field(i).Interface(),
			)
		}
	}

	return output
}

func handleOutput(content string, outputFile string) error {
	if outputFile != "" {
		file, err := os.Create(outputFile)
		if err != nil {
			return fmt.Errorf("failed to create output file: %w", err)
		}
		defer file.Close()

		_, err = file.WriteString(content)
		if err != nil {
			return fmt.Errorf("failed to write to output file: %w", err)
		}
		return nil
	}

	fmt.Println(content)
	return nil
}
