package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"bio-algos/algos"
)

// Genome to analyse should be provided as command line argument
func main() {
	// Requires command line argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <genome file location>")
		os.Exit(1)
	}
	
	filename := os.Args[1]
	
	// Read genome from a text file
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	genome := string(data)
	fmt.Printf("Loaded genome with %d base pairs\n\n", len(genome))

	scanner := bufio.NewScanner(os.Stdin)

	for {
		showMenu()
		fmt.Print("Enter your choice ('q' to quit): ")

		if !scanner.Scan() {
			break
		}

		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			patternCountAnalysis(genome, scanner)
		case "2":
			frequentSubstringAnalysis(genome, scanner)
		case "3":
			showGenomeInfo(genome)
		case "q", "Q":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice, try again.\n")
		}
	}
}
		

// Show menu options
func showMenu() {
	fmt.Println("=== Genome Analysis Options ===")
	fmt.Println("1. Pattern Count - Count occurrences of a speific pattern")
	fmt.Println("2. Frequent Substrigns - Find most frequent k-mers")
	fmt.Println("3. Genome Info - Show basic genome stats")
	fmt.Println("q. Quit")
	fmt.Println()
}

// Pattern count analysis
func patternCountAnalysis(genome string, scanner *bufio.Scanner) {
	fmt.Println("Enter the pattern to search for: ")
	if !scanner.Scan() {
		return
	}

	pattern := strings.TrimSpace(scanner.Text())
	if pattern == "" {
		fmt.Println("Pattern cannot be empty. \n")
		return
	}

	result := algos.PatternCount(genome, pattern)
	fmt.Printf("Pattern '%s' appears %d times in the genome\n\n", pattern, result)
}

// Frequent words analysis
func frequentSubstringAnalysis(genome string, scanner *bufio.Scanner) {
	k := getKValue(scanner)
	if k == -1 {
		return
	}

	frequentSubstrings := algos.FrequentSubstrings(genome, k)
	freqTable := algos.FrequencyTable(genome, k)
	maxCount := algos.MaxMap(freqTable)

	fmt.Printf("\nMost frequent %d-mers (appearing %d times):\n", k, maxCount)
	for _, word := range frequentSubstrings {
		fmt.Printf("- %s\n", word)
	}
	fmt.Printf("\nFound %d most frequent %d-mers\n\n", len(frequentSubstrings), k)
}

// Not yet implemented
func showGenomeInfo(genome string) {
	fmt.Printf("\n=== Genome Statistics ===\n")
	// work on this section

}

// Get k value from user
func getKValue(scanner *bufio.Scanner) int {
	fmt.Print("Enter k-mer size (k): ")
	if !scanner.Scan() {
		return -1
	}

	kStr := strings.TrimSpace(scanner.Text())
	k, err := strconv.Atoi(kStr)
	if err != nil {
		fmt.Println("Invalid number. Please enter a valid integer.\n")
		return -1
	}

	if k <= 0 {
		fmt.Println("k must be greater than 0.\n")
		return -1
	}

	return k
}
