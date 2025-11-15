package isogram


import "strings"

// IsIsogram checks if the given word string is an isogram
// return true if the word is an isogram, false otherwise
func IsIsogram(word string) bool {
	memory := make(map[rune]bool)
    word = strings.ToLower(word)
    for _, char := range word {
        if char == ' ' || char == '-' {
            continue
        }
        if memory[char] == true {
            return false
        }
        memory[char] = true
    }
    return true
}
