package legalform

import (
	"regexp"
	"strings"
)

var tokenRegexp *regexp.Regexp

func init() {
	tokenRegexp = regexp.MustCompile(`[^\s]+`)
}

// Strip strips the legal form from the full company name and returns the plain
// name as well as the legal form independently.
func (f LegalForms) Strip(fullName string) (string, string) {
	tokens := tokenRegexp.FindAllString(fullName, -1)

	currentTokens := []string{}
	currentCleanTokens := []string{}
	legalFormTokens := []string{}

	for i := len(tokens) - 1; i > 0; i-- {
		token := tokens[i]
		cleanToken := clean(token)

		currentTokens = append([]string{token}, currentTokens...)
		currentCleanTokens = append([]string{cleanToken}, currentCleanTokens...)
		tokenSearch := strings.Join(currentCleanTokens, "")

		if _, ok := f[tokenSearch]; ok {
			legalFormTokens = append(currentTokens, legalFormTokens...)
			currentTokens = []string{}
		}
	}

	company := strings.Join(tokens[0:len(tokens)-len(legalFormTokens)], " ")
	legalForm := strings.Join(legalFormTokens, " ")
	return company, legalForm
}

func clean(s string) string {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, ".", "")
	s = strings.ReplaceAll(s, "-", "")
	s = strings.ReplaceAll(s, "/", "")
	s = strings.ReplaceAll(s, `"`, "")
	s = strings.ReplaceAll(s, "’", "")
	s = strings.ReplaceAll(s, "(", "")
	s = strings.ReplaceAll(s, ")", "")
	s = strings.ReplaceAll(s, "&", "")
	s = strings.ReplaceAll(s, "'", "")
	s = strings.ReplaceAll(s, ",", "")
	s = strings.ReplaceAll(s, ":", "")
	s = strings.ReplaceAll(s, " ", "")
	return s
}
