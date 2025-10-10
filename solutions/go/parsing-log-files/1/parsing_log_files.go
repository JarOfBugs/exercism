package parsinglogfiles

import (
    "regexp"
    "fmt"
)

func IsValidLine(text string) bool {
	re, err := regexp.Compile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
    if err != nil {
        fmt.Printf("Error Parsing Regexp %+v\n", err)
        return false
    }
    return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`\s*<[~*=\-]*>\s*`)
    return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
    count := 0
    r := regexp.MustCompile(`(?i)".*password.*"`)
	for _, line := range lines {
        if r.MatchString(line) {
            count++
        }
    }
    return count
}

func RemoveEndOfLineText(text string) string {
	r := regexp.MustCompile(`end-of-line\S*`)
    return r.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
    substrRegexp := regexp.MustCompile(`User\s+(\S+)\b`)
	result := make([]string, len(lines))

	for i, line := range lines {
		match := substrRegexp.FindStringSubmatch(line)
		if len(match) == 2 {
			userName := match[1]
			
			result[i] = fmt.Sprintf("[USR] %s %s", userName, line)
		} else {
			result[i] = line
		}
	}

	return result
}
