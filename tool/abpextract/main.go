package main

// Albertirsa

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	inputPath  = flag.String("inputpath", "", "Path to your ABP-style file directory")
	outputName = flag.String("outputname", "", "Name of the generated domain-list format file")
	outputDir  = flag.String("outputdir", "./", "Directory to place all generated files")
)

type RuleType int

const (
	RuleDomain     RuleType = 1
	RuleFullDomain RuleType = 2
	RuleKeyword    RuleType = 3
	RuleRegexp     RuleType = 4
	RuleUnknown    RuleType = -1
)

func rulePick(line string) (raw string, exception bool, skip bool) {
	line = strings.TrimSpace(line)

	if line == "" || strings.HasPrefix(line, "!") || strings.HasPrefix(line, "[") {
		return "", false, true
	}
	if strings.HasPrefix(line, "@@") {
		exception = true
		line = strings.TrimPrefix(line, "@@")
	}

	return line, exception, false
}

func isFastRegexp(s string) bool {
	return len(s) >= 2 && strings.HasPrefix(s, "/")
}

func convertRegexp(s string) (string, bool) {
	if !isFastRegexp(s) {
		return "", false
	}

	pattern := s[1 : len(s)-1]
	if after, ok := strings.CutPrefix(pattern, `^https?:\/\/`); ok {
		pattern = after
	}

	// Validate syntax using Go's regexp engine (RE2).
	if _, err := regexp.Compile(pattern); err != nil {
		return "", false
	}

	return pattern, true
}

func extractHost(s string) string {
	s = strings.TrimSpace(s)

	// Strip ABP anchors.
	s = strings.TrimPrefix(s, "||")
	s = strings.TrimPrefix(s, "|")

	// Handle explicit URL schemes.
	if strings.Contains(s, "://") {
		u, err := url.Parse(s)
		if err != nil {
			return ""
		}
		s = u.Hostname()
		return strings.TrimSuffix(s, ".")
	}

	// Remove the resource path and query/fragment.
	if i := strings.IndexAny(s, "/?#"); i >= 0 {
		s = s[:i]
	}

	// ABP's trailing separator marker.
	s = strings.TrimSuffix(strings.TrimSuffix(s, "^"), ".")

	ifIP := net.ParseIP(s)
	if ifIP != nil {
		return ""
	}

	return s
}

func formatRule(kind RuleType, value string) string {
	switch kind {
	case RuleFullDomain:
		return "full:" + value
	case RuleKeyword:
		return "keyword:" + value
	case RuleRegexp:
		return "regexp:" + value
	case RuleDomain:
		return value
	default:
		return ""
	}
}

func wildcardToKeyword(host string) (string, bool) {
	if !strings.Contains(host, "*") {
		return host, false
	}

	if strings.HasPrefix(host, "*") || strings.HasSuffix(host, "*") {
		keyword := strings.Trim(host, "*")
		if keyword == "" {
			return "", false
		}
		return keyword, true
	}

	return host, false
}

func middleWildcardToRegexp(host string) (string, bool) {
	if !strings.Contains(host, "*") {
		return host, false
	}

	pattern := strings.ReplaceAll(host, `*`, `(.*)`)

	return pattern, true
}

func processLine(line string) (rule string, exception bool, ok bool) {
	raw, exception, skip := rulePick(line)
	if skip {
		return "", false, false
	}

	// Handle explicit ABP regexp syntax first.
	if pattern, matched := convertRegexp(raw); matched {
		return formatRule(RuleRegexp, pattern), exception, true
	}

	if keyword, matched := wildcardToKeyword(extractHost(raw)); matched {
		return formatRule(RuleKeyword, keyword), exception, true
	}

	// Preserve the distinction between || and |.
	isDomain := strings.HasPrefix(raw, "||")
	isFull := !isDomain && strings.HasPrefix(raw, "|")

	host := extractHost(raw)
	if host == "" {
		return "", exception, false
	}

	if midword, matched := middleWildcardToRegexp(host); matched {
		return formatRule(RuleRegexp, midword), exception, true
	}

	if isFull {
		return formatRule(RuleFullDomain, host), exception, true
	}

	if isDomain {
		return formatRule(RuleDomain, host), exception, true
	}

	return "", exception, false
}

func main() {
	flag.Parse()

	if err := os.MkdirAll(*outputDir, 0755); err != nil {
		fmt.Errorf("failed to create output directory: %w", err)
		os.Exit(1)
	}
	outputPath := filepath.Join(*outputDir, *outputName)
	outputExcludePath := filepath.Join(*outputDir, *outputName+"-exclude")

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

	outExcludeFile, err := os.Create(outputExcludePath)
	if err != nil {
		fmt.Errorf("failed to create output file: %w", err)
		os.Exit(1)
	}
	defer outExcludeFile.Close()

	scanner := bufio.NewScanner(inFile)
	writer := bufio.NewWriter(outFile)
	writerExclude := bufio.NewWriter(outExcludeFile)
	defer writer.Flush()
	defer writerExclude.Flush()

	var results []string
	var resultsExclude []string

	// Read and process each line.
	for scanner.Scan() {
		line := scanner.Text()

		result, exclude, yes := processLine(line)
		if yes {
			if !exclude {
				results = append(results, result)
			} else {
				resultsExclude = append(resultsExclude, result)
			}
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

	for _, result := range resultsExclude {
		if _, err := fmt.Fprintln(writerExclude, result); err != nil {
			fmt.Errorf("failed to write cache: %w", err)
			return
		}
	}

	if err := writerExclude.Flush(); err != nil {
		fmt.Errorf("failed to write file: %w", err)
		os.Exit(1)
	}
}
