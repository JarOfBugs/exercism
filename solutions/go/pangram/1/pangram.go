// Pangram is a sentence that contains all the letter of the english alphabet
package pangram

import "strings"

// IsPangram checks if a given string is a pangram
// It returns true if the sentence contains all 26 letter of the english alphabet,
// ignoring case, space and punctuation
func IsPangram(input string) bool {
    if len(input) < 25 {
        return false
    }
    
    input = strings.ToLower(input)
    
    var count int
    var visited map[rune]bool = make(map[rune]bool)
    
    for _, char := range input {
        if char < 'a' || char > 'z' {
            continue
        }
        
        if !visited[char] {
            visited[char] = true
            count++
        }
    }
    return count == 26
}
