package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/fatih/color"
)

// Function to print the banner
func printBanner() {
	banner := `
  _     ____  _   _____  ____  _  ____ 
 / \ /\/  __\/ \ /__ __\/  _ \/ \/  __\
 | | |||  \/|| |   / \  | / \|| ||  \/|
 | \_/||    /| |_/\| |  | \_/|| ||  __/
 \____/\_/\_\\____/\_/  \____/\_/\_/   
                                      
`
	// Print the banner with color
	color.Cyan(banner)
	color.Green("Version 1.0 - Created by AlulCode45")
	fmt.Println()
}

func main() {
	// Print the banner at the beginning
	printBanner()

	// Set up flags for short (-f, -o, -d) and long versions (--file, --output, --domain)
	inputFile := flag.String("f", "", "Input file containing domain list (optional)")
	outputFile := flag.String("o", "", "Output file to save resolved IPs")
	domain := flag.String("d", "", "Resolve a single domain (optional)")

	// Long versions of the flags
	flag.StringVar(inputFile, "file", "", "Input file containing domain list (optional)")
	flag.StringVar(outputFile, "output", "", "Output file to save resolved IPs")
	flag.StringVar(domain, "domain", "", "Resolve a single domain (optional)")

	showHelp := flag.Bool("help", false, "Show usage instructions")

	// Parse the command-line flags
	flag.Parse()

	// Show help if --help is passed
	if *showHelp {
		color.Yellow("Usage instructions:")
		fmt.Println("  -f, --file <file>      : Read domain list from a file.")
		fmt.Println("  -o, --output <file>    : Save the resolved IPs to an output file.")
		fmt.Println("  -d, --domain <domain>  : Resolve a single domain.")
		fmt.Println("Examples:")
		fmt.Println("  go run main.go -f domains.txt -o ips.txt")
		fmt.Println("  go run main.go -d google.com -o result.txt")
		return
	}

	// Ensure either input file or domain is provided
	if *inputFile == "" && *domain == "" {
		color.Red("Error: Please provide either a file input or a domain. Use --help for instructions.")
		return
	}

	// Open output file if specified
	var writer *bufio.Writer
	if *outputFile != "" {
		outFile, err := os.Create(*outputFile)
		if err != nil {
			color.Red("Error creating output file: %v", err)
			return
		}
		defer outFile.Close()
		writer = bufio.NewWriter(outFile)
	}

	// Function to write results to file or stdout
	writeResult := func(result string) {
		if writer != nil {
			_, err := writer.WriteString(result + "\n")
			if err != nil {
				color.Red("Error writing to file: %v", err)
				return
			}
		} else {
			fmt.Println(result)
		}
	}

	// Function to resolve domain to IP
	resolveDomain := func(domain string) {
		ips, err := net.LookupIP(domain)
		if err != nil {
			writeResult(fmt.Sprintf("Error resolving domain %s: %v", domain, err))
			return
		}
		for _, ip := range ips {
			writeResult(ip.String())
		}
	}

	// If input file is provided, read domains from file
	if *inputFile != "" {
		file, err := os.Open(*inputFile)
		if err != nil {
			color.Red("Error opening input file: %v", err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			domain := scanner.Text()
			resolveDomain(domain)
		}

		if err := scanner.Err(); err != nil {
			color.Red("Error reading input file: %v", err)
		}
	}

	// If domain is provided, resolve it directly
	if *domain != "" {
		resolveDomain(*domain)
	}

	// Flush the output writer if used
	if writer != nil {
		writer.Flush()
	}
}
