package main

// Csepreg

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/net/idna"
)

var (
	inputPath  = flag.String("inputpath", "", "Path to your DNSmasq-style file directory")
	outputName = flag.String("outputname", "", "Name of the generated domain-list format file")
	outputDir  = flag.String("outputdir", "./", "Directory to place all generated files")
)

func sanitize(line string) string {
	if strings.HasPrefix(line, `"`) {
		return strings.TrimPrefix(line, `"`)
	}

	puny, err:= idna.ToASCII(line)
	if err != nil {
		return ""
	}

	return puny
}

func main() {
	flag.Parse()

	if err := os.MkdirAll(*outputDir, 0755); err != nil {
		fmt.Errorf("failed to create output directory: %w", err)
		os.Exit(1)
	}
	outputPath := filepath.Join(*outputDir, *outputName)

	// Open the input file.
	inFile, err := os.Open(*inputPath)
	if err != nil {
		fmt.Errorf("failed to open file: %w", err)
		os.Exit(1)
	}
	defer inFile.Close()

	// Create or truncate the output file.
	outFile, err := os.Create(outputPath)
	if err != nil {
		fmt.Errorf("failed to create output file: %w", err)
		os.Exit(1)
	}
	defer outFile.Close()

	scanner := bufio.NewScanner(inFile)
	writer := bufio.NewWriter(outFile)
	defer writer.Flush()

	var results []string

	// Read and process each line.
	for scanner.Scan() {
		line := scanner.Text()

		result := sanitize(line)
		if result != "" {
			results = append(results, result)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Errorf("failed to read file: %w", err)
		os.Exit(1)
	}

	// Write the collected results, one per line.
	for _, result := range results {
		if _, err := fmt.Fprintln(writer, result); err != nil {
			fmt.Errorf("failed to write cache: %w", err)
			return
		}
	}

	if err := writer.Flush(); err != nil {
		fmt.Errorf("failed to write file: %w", err)
		os.Exit(1)
	}
}
