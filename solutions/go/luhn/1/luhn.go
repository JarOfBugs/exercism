package luhn

import "strings"

// Valid verifies the given id numeric string according to the luhn formula
// returns true if the id is valid, false otherwise
func Valid(id string) bool {
    id = strings.ReplaceAll(id, " ", "")
    length := len(id)
    
    if length <= 1 {
        return false
    }
    
    var sum int
    var j int

    for i := length-1; i >= 0; i-- {
        char := id[i]
        if char < '0' || char > '9' {
             return false
        }
        digit := int(char - '0')
        if j % 2 == 1 {
            digit *= 2
            if digit > 9 {
                digit -= 9
            }
        }
        sum += digit
        j++
    }
    return sum % 10 == 0
}
