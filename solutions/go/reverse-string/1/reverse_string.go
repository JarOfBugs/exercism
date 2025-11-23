package reverse

// Reverse take a string in input and return its reversed version
func Reverse(input string) string {
	runes := []rune(input)
    for l, r := 0, len(runes) - 1; l < r; l, r = l + 1, r - 1 {
        runes[l], runes[r] = runes[r], runes[l]
    }
    return string(runes)
}
