package main

import (
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	logArr := []string{"TRC", "DBG", "INF", "WRN", "ERR", "FTL"}

	re := regexp.MustCompile(`[A-Z]`)
	for _, item := range logArr {
		if text[1:4] == item {
			matchString := re.MatchString(text[1:4])
			if matchString {
				return true
			}
		}
	}

	return false
}

func SplitLogLine(text string) []string {
	if len(text) == 0 {
		return []string{""}
	}

	re := regexp.MustCompile("(?:\\w+? \\d+|<start>|<end>)")

	findLines := re.FindAllString(text, -1)

	return findLines
}

func CountQuotedPasswords(lines []string) int {
	if len(lines) == 0 {
		return 0
	}

	count := 0

	for _, item := range lines {
		re := regexp.MustCompile("\"[^\"]*?password[^\"]*?\"")
		if re.MatchString(strings.ToLower(item)) {
			count++
		}
	}

	return count
}

func RemoveEndOfLineText(text string) string {
	if len(text) == 0 {
		return ""
	}

	re := regexp.MustCompile("end-of-line\\d+")

	findLines := re.ReplaceAllString(text, "")

	return findLines
}

func TagWithUserName(lines []string) []string {
	if len(lines) == 0 {
		return []string{""}
	}
	var out []string
	for _, item := range lines {
		re := regexp.MustCompile("^(.*?User\\s+?)([\\d\\w]+?)(\\s.*)$")
		out = append(out, re.ReplaceAllString(item, "[USR] $2 $1$2$3"))

	}
	return out
}

func main() {
	//fmt.Println(SplitLogLine("section 1<*>section 2<~~~>section 3"))
	//fmt.Println(SplitLogLine("<start> <end>"))

	result := TagWithUserName([]string{
		"[WRN] User James123 has exceeded storage space.",
		"[WRN] Host down. User   Michelle4 lost connection.",
		"[INF] Users can login again after 23:00.",
		"[DBG] We need to check that user names are at least 6 chars long.",
	})

	TagWithUserName(result)
}
