package resistorcolor

// Colors returns the list of all colors.
func Colors() []string {
	return []string {
        "black",
        "brown",
        "red",
        "orange",
        "yellow",
        "green",
        "blue",
        "violet",
        "grey",
        "white",
    }
}

// ColorCode returns the resistance value of the given color.
// returns the color code as integer or -1 if it is not found
func ColorCode(color string) int {
	for i, c := range Colors() {
        if c == color {
            return i
        }
    }
    return -1
}
