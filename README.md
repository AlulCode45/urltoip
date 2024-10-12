# Domain to IP Resolver

This is a simple Go program that resolves domain names into their corresponding IP addresses. It supports reading domains from a file or resolving a single domain directly. The results can be saved to an output file. The program includes a colorful banner and version information for better presentation.

## Features
- **Resolve from File**: Input a list of domains from a file and resolve each domain to its IP address.
- **Single Domain Resolution**: Resolve a single domain directly via command-line arguments.
- **Save Results**: Save the resolved IPs to an output file or display them directly in the terminal.
- **Colorful Output**: The program includes a banner and uses colored output for enhanced readability in the terminal.
- **Help Command**: Get instructions on how to use the program via the `--help` flag.

## Usage

### Command-line Options
The program supports both short and long flag versions for flexibility:
- `-f, --file` : Specify the input file containing the list of domains (optional).
- `-d, --domain` : Specify a single domain to resolve.
- `-o, --output` : Specify the output file to save the resolved IP addresses.
- `--help` : Display the usage instructions.

### Examples

1. **Resolving a Single Domain**
   ```bash
   go run main.go -d google.com -o result.txt
   ```

   This command resolves the domain `google.com` and saves the IP address to `result.txt`.

2. **Resolving Domains from a File**
   ```bash
   go run main.go -f domains.txt -o ips.txt
   ```

   This command reads the domain names from `domains.txt`, resolves their IP addresses, and saves the results to `ips.txt`.

3. **Displaying Help**
   ```bash
   go run main.go --help
   ```

   This command displays the usage instructions with details about the available options.

### Output
If an output file is not specified, the resolved IPs are printed directly to the terminal. For example:
```bash
go run main.go -d google.com
```
Will display:
```
142.250.190.78
```

### Installation

1. **Clone the Repository**
   ```bash
   git clone https://github.com/AlulCode45/urltoip
   ```

2. **Install Dependencies**
   The program uses the `color` package for colorful output. Install it with:
   ```bash
   go get github.com/fatih/color
   ```

3. **Run the Program**
   Use `go run` to run the program:
   ```bash
   go run main.go -d example.com
   ```
## Author

This project was created by **AlulCode45**. You can find more of my projects on my [GitHub profile](https://github.com/AlulCode45/).

## License

This project is open-source and available under the [MIT License](LICENSE).
```