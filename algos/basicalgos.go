package algos

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

