package algos

import "strings"

// Count the number of times the provided sequence occurs in the genome
func PatternCount(genome string, pattern string) int {
	if len(pattern) == 0 {
		return 0
	}

	var count int 
	for i := 0; i <= len(genome)-len(pattern); i++ {
		if genome[i:i+len(pattern)] == pattern {
			count++
		}
	}
	return count
}

// Build a frequency table of all k-mers in genome
func FrequencyTable(genome string, k int) map[string]int {
	freqMap := make(map[string]int)

	// Validate k
	if k <= 0 || k > len(genome) {
		return freqMap
	}

	for i := 0; i <= len(genome)-k; i++ {
		pattern := genome[i:i+k]
		freqMap[pattern]++
	}
	return freqMap
}

// Get the most frequent value from the frequency table
func MaxMap(table map[string]int) int {
	result := 0
	for _, value := range table {
		if value > result {
			result = value
		}
	}
	return result
}

// Returns a list of the most frquent substrings
func FrequentSubstrings(genome string, k int) []string {
	var freqPatterns []string
	freqMap := FrequencyTable(genome, k)
	maximum := MaxMap(freqMap)
	for key, value := range freqMap {
		if value == maximum {
			freqPatterns = append(freqPatterns, key)
		}
	}
	return freqPatterns
}

// Reverse complement function
func ReverseComplement(pattern string) string {
	var builder strings.Builder
	builder.Grow(len(pattern))

	for i := len(pattern) - 1; i >= 0; i-- {
		switch pattern[i] {
		case 'A':
			builder.WriteByte('T')
		case 'T':
			builder.WriteByte('A')
		case 'C':
			builder.WriteByte('G')
		case 'G':
			builder.WriteByte('C')
		default:
			builder.WriteByte('N')
		}
	}
	return builder.String()
}

// Pattern matching - return the positions in a sequence where a pattern is found
func PatternMatching(pattern string, genome string) []int {
	// Pattern validation
	if len(pattern) == 0 || len(pattern) > len(genome) {
		return []int{}
	}

	var result []int
	for i := 0; i <= len(genome) - len(pattern); i++ {
		if genome[i:i+len(pattern)] == pattern {
			result = append(result, i)
		}
	}
	return result
}

