package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"bio-algos/algos"
)

// Genome to analyse should be provided as command line argument
func main() {
	// Requires command line argument
	//if len(os.Args) < 2 {
	//	fmt.Println("Usage: go run . <genome file location>")
	//	os.Exit(1)
	//}
	
	//filename := os.Args[1]
	
	genomeFiles, err := findGenomeFiles("genomes")
	if err != nil {
		log.Fatal("Error accessing genomes folder:", err)
	}

	if len(genomeFiles) == 0 {
		fmt.Println("No genome files found in the 'genomes' folder. Add some '.txt' files to get started!")
		os.Exit(1)
	}

	// Scanner - used throughout main 
	scanner := bufio.NewScanner(os.Stdin)
	
	//selectedFile := selectGenomeFile(genomeFiles, scanner)
	//if selectedFile == "" {
	//	fmt.Println("no file selected. Exiting.")
	//	return
	//}

	selectedFile, genome := selectAndLoadGenome(genomeFiles, scanner)
	if selectedFile == "" {
		fmt.Println("No file selected. Exiting.")
		return
	}

	// Read genome from a text file
	//data, err := os.ReadFile(filename)
	//if err != nil {
	//	log.Fatal("Error reading file:", err)
	//}

	//genome := string(data)


	//genome, err := loadGenome(selectedFile)
	//if err != nil {
	//	log.Fatal("Error loading genome:", err)
	//}

	//fmt.Printf("Loaded genome with %d base pairs\n\n", len(genome))


	for {
		showMenu()
		fmt.Print("Enter your choice ('q' to quit): ")

		if !scanner.Scan() {
			break
		}

		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			newFile, newGenome := selectAndLoadGenome(genomeFiles, scanner)
			if newFile != "" && newFile != selectedFile {
				genome = newGenome
				selectedFile = newFile
			}
		case "2":
			patternCountAnalysis(genome, scanner)
		case "3":
			frequentSubstringAnalysis(genome, scanner)
		case "4":
			showGenomeInfo(genome)
		case "5":
			patternMatchingAnalysis(genome, scanner)
		case "q", "Q":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice, try again.\n")
		}
	}
}
		

// Show menu options - update as new options added
func showMenu() {
	fmt.Println("=== Genome Analysis Options ===")
	fmt.Println("1. Switch Genome - Load a different genome file")
	fmt.Println("2. Pattern Count - Count occurrences of a speific pattern")
	fmt.Println("3. Frequent Substrigns - Find most frequent k-mers")
	fmt.Println("4. Genome Info - Show basic genome stats")
	fmt.Println("5. Pattern Matching - Find location of patterns in a genome")
	fmt.Println("q. Quit")
	fmt.Println()
}


func findGenomeFiles(dir string) ([]string, error) {
	var files []string

	// Check directory exists
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return nil, fmt.Errorf("genomes folder does not exist")
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(strings.ToLower(entry.Name()), ".txt") {
			files = append(files, filepath.Join(dir, entry.Name()))
		}
	}
	return files, nil
}


func selectAndLoadGenome(files []string, scanner *bufio.Scanner) (string, string) {
	selectedFile := selectGenomeFile(files, scanner)
	if selectedFile == "" {
		return "", ""
	}

	genome, err := loadGenome(selectedFile)
	if err != nil {
		fmt.Printf("Error loading genome: %v\n\n", err)
		return "", ""
	}

	fmt.Printf("Loaded genome from '%s' with %d nucleotides\n\n", filepath.Base(selectedFile), len(genome))
	return selectedFile, genome
}

func selectGenomeFile(files []string, scanner *bufio.Scanner) string {
	fmt.Println("=== Available Genome Files ===")
	for i, file := range files {
		filename := filepath.Base(file)
		fmt.Printf("%d. %s\n", i+1, filename)
	}
	fmt.Println()

	for {
		fmt.Printf("Select a genome file (1-%d): ", len(files))

		if !scanner.Scan() {
			return ""
		}

		choice := strings.TrimSpace(scanner.Text())
		num, err := strconv.Atoi(choice)

		if err != nil || num < 1 || num > len(files) {
			fmt.Println("Invalid choice. Please enter a number between 1 and %d.\n", len(files))
			continue
		}

		return files[num-1]
	}
}

func loadGenome(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	// Clean up sequence - add FASTA cleaning later if required
	genome := strings.ToUpper(strings.ReplaceAll(string(data), " ", ""))
	genome = strings.ReplaceAll(genome, "\n", "")
	genome = strings.ReplaceAll(genome, "\r", "")
	genome = strings.ReplaceAll(genome, "\t", "")

	// Validate nucleotides - update if RNA added at later date
	validNucleotides := "ATCG"
	for _, char := range genome {
		if !strings.ContainsRune(validNucleotides, char) {
			return "", fmt.Errorf("invalid nucleotide '%c' found in genome.", char)
		}
	}
	return genome, nil
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
	fmt.Println("***UNDER CONSTRUCTION***")
	// work on this section

}

// Pattern matching analysis
func patternMatchingAnalysis(genome string, scanner *bufio.Scanner) {
	fmt.Println("Enter the pattern to search for: ")
	if !scanner.Scan() {
		return
	}

	pattern := strings.TrimSpace(scanner.Text())
	if pattern == "" {
		fmt.Println("Pattern cannot be empty. \n")
		return
	}

	result := algos.PatternMatching(pattern, genome)
	fmt.Printf("Pattern '%s' appears at the following positions in the genome\n\n", pattern)

	if len(result) == 0 {
		fmt.Println("No matches found.\n")
	} else {
		positions := make([]string, len(result))
		for i, pos := range result {
			positions[i] = strconv.Itoa(pos)
		}
		fmt.Println(strings.Join(positions, " "))
		fmt.Printf("\nTotal matches: %d\n\n", len(result))
	}
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
