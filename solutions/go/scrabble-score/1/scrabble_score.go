package scrabble

import "strings"

// Score calculates the scrabble score of a given word
// returns the score as an integer
func Score(word string) int {
	total :=0
    word = strings.ToUpper(word)
    for _, char := range word {
        total += getLetterScore(char)
    }
    return total
}

// getLetterScore calculates the value of a single letter (it handle only upper case letter)
// returns the score as integer
func getLetterScore(letter rune) int {
    switch(letter) {
		case 'D','G':
        	return 2
        case 'B','C','M','P':
        	return 3
        case 'F','H','V','W','Y':
        	return 4
        case 'K':
        	return 5
        case 'J','X':
        	return 8
        case 'Q', 'Z':
        	return 10
        default:
        	return 1
    }
}
