// Pangram is a sentence that contains all the letter of the english alphabet
package pangram

import "strings"

// IsPangram checks if a given string is a pangram
// It returns true if the sentence contains all 26 letter of the english alphabet,
// ignoring case, space and punctuation
func IsPangram(input string) bool {
    input = strings.ToLower(input)
    for letter := 'a'; letter <= 'z'; letter++ {
    	if !strings.ContainsRune(input, letter) {
            return false
        }  
    }
    return true
}
