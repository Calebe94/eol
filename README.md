# EOL CLI

A command-line interface for checking product lifecycle information from [endoflife.date](https://endoflife.date).

[![Go Report Card](https://img.shields.io/github/go-mod/go-version/lnutimura/eol)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-apache-blue.svg)](https://opensource.org/licenses/apache-2-0)

## Features

- List all supported products
- View detailed lifecycle information for any product
- Watch mode for monitoring changes
- Multiple output formats (JSON, Table, CSV)
- Filtering and sorting capabilities
- Cross-platform support

## Installation

### Using Go (requires Go 1.20+)
```bash
go install github.com/lnutimura/eol@latest
```

### From Source
```bash
git clone https://github.com/lnutimura/eol.git
cd eol
sudo make install
```

## Usage

```
eol [command] [arguments] [flags]

Commands:
  list        List all available products
  product     Show lifecycle details for a product
  cycle       Get details for a specific product cycle
  help        Show help information

Flags:
  -f, --format string   Output format (table|json|csv) (default "table")
  -o, --output string   Write output to file
  -w, --watch duration  Auto-refresh interval (e.g., 5m, 1h)
  --no-color            Disable color output
  -h, --help            Show help
```

## Examples

### List all products
```bash
eol list
```

### Get product details (JSON format)
```bash
eol product ubuntu --format json
```

### Check specific cycle
```bash
eol cycle python 3.7
```

### Monitor RHEL cycles
```bash
eol product rhel --watch 15m
```

### Filter active support cycles
```bash
eol product dotnet --filter "support = true"
```

### Generate CSV report
```bash
eol product android --format csv --output android_report.csv
```

## Supported Output Formats

| Format  | Description                     | Example Use Case          |
|---------|---------------------------------|---------------------------|
| Table   | Human-readable formatted table  | Quick terminal inspection |
| JSON    | Raw JSON output                 | Script integration        |
| CSV     | Comma-separated values          | Spreadsheet import        |

## Contributing

Contributions are welcome! Please follow these steps:
1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Acknowledgements

- Data provided by [endoflife.date](https://endoflife.date)
- Inspired by [endoflife.date/docs](https://endoflife.date/docs)

## License

Apache License - See [LICENSE](LICENSE) for details.
