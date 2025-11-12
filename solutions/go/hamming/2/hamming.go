package hamming

import "errors"

// Distance calcuate the Hamming distance between 2 DNA strand
// returns the Hamming distance as integer 
// returns error in case the 2 strands have different lengths
func Distance(a, b string) (int, error) {
    count := 0
    bLength := len(b)
    if len(a) != bLength {
        return count, errors.New("The DNA strands sample must be of the same length!")
    }
	for i := range bLength {
		if a[i] != b[i] {
            count++
        }
    }
    return count, nil
}
