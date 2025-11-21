package strand

// ToRNA convert dna string to rna string
func ToRNA(dna string) string {
    var DNA_TRANSCODE_MAP = map[string]string{
        "G": "C",
        "C": "G",
        "T": "A",
        "A": "U",
    }
	var result string
    for _,char := range dna {
        result += DNA_TRANSCODE_MAP[string(char)]
    }
    return result
}
