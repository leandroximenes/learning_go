package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\] .+`)
	rpl := re.MatchString(text)
	return rpl
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<>|<[~\-*=]+>`)
	slp := re.Split(text, -1)
	return slp
}

func CountQuotedPasswords(lines []string) int {
	qtde := 0
	for _, line := range lines {
		re := regexp.MustCompile(`(?i)(\\?['"]\bpassword\b\\?['"])`)
		matches := re.FindStringSubmatch(line)
		qtde += len(matches)
	}
	return qtde
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	result := re.ReplaceAllString(text, "")
	return result
}

func TagWithUserName(lines []string) []string {
	for k, line := range lines {
		re := regexp.MustCompile(`User\s+(\w+)`)
		user := re.FindStringSubmatch(line)
		if len(user) > 0 {
			lines[k] = fmt.Sprintf("[USR] %s %s", user[1], line)
		}
	}
	return lines
}
