package main

import "fmt"

// Application identifies the application emitting the given log.
func Application(log string) string {
	strOfChar := map[rune]string{
		'❗': "recommendation",
		'🔍': "search",
		'☀': "weather",
	}

	runeLogs := []rune(log)

	for _, runeLog := range runeLogs {
		if _, ok := strOfChar[runeLog]; ok {
			return strOfChar[runeLog]
		}
	}

	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	result := ""
	runeLogs := []rune(log)
	for _, runeLog := range runeLogs {
		if runeLog == oldRune {
			runeLog = newRune
		}
		result += string(runeLog)
	}
	return result
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	runeLogs := []rune(log)
	if (len(runeLogs) - 1) < limit {
		return false
	}

	return true
}

func main() {
	log := "please replace '👎' with '👍'"

	fmt.Println(Replace(log, '👎', '👍'))
}
