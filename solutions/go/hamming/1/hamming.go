package hamming

import "errors"

// Distance calcuate the Hamming distance between 2 DNA strand
// returns the Hamming distance as integer 
// returns error in case the 2 strands have different lengths
func Distance(a, b string) (int, error) {
    count := 0
    if len(a) != len(b) {
        return count, errors.New("The DNA strands sample must be of the same length!")
    }
    aRune := []rune(a)
	for i,char := range b {
		if aRune[i] != char {
            count++
        }
    }
    return count, nil
}
