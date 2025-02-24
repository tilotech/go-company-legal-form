package legalform

import (
	"strings"
)

// Strip strips the legal form from the end of the full company name and returns the plain
// name as well as the legal form independently.
func (f LegalForms) Strip(fullName string) (string, string) {
	tokens := strings.Fields(fullName)

	cleanTokens := make([]string, len(tokens))
	currentTokenLength := 0
	legalFormTokenLength := 0

	for i := len(tokens) - 1; i > 0; i-- {
		currentTokenLength++
		cleanTokens[i] = clean(tokens[i])
		tokenSearch := strings.Join(cleanTokens[len(tokens)-currentTokenLength-legalFormTokenLength:], "")
		if _, ok := f[tokenSearch]; ok {
			legalFormTokenLength += currentTokenLength
			currentTokenLength = 0
		}
	}

	legalFormTokenStart := len(tokens) - legalFormTokenLength
	company := strings.Join(tokens[0:legalFormTokenStart], " ")
	legalForm := strings.Join(tokens[legalFormTokenStart:], " ")
	return company, legalForm
}

// StripMiddle strips the legal form from anywhere in the full company name and
// returns the name as well as the legal form.
//
// If the legal form was not at the end, then everything after the legal form
// will be returned as the third response value.
func (f LegalForms) StripMiddle(fullName string) (string, string, string) {
	tokens := strings.Fields(fullName)
	cleanTokens := make([]string, len(tokens))

	for i := range tokens {
		cleanTokens[i] = clean(tokens[i])
	}

	for i := len(tokens) - 1; i > 0; i-- {
		searchEndIdx := i + 1
		obsoleteTokenLength := len(tokens) - searchEndIdx
		currentTokenLength := 0
		legalFormTokenLength := 0
		for j := i; j > 0; j-- {
			currentTokenLength++
			searchStartIdx := len(tokens) - currentTokenLength - legalFormTokenLength - obsoleteTokenLength
			tokenSearch := strings.Join(cleanTokens[searchStartIdx:searchEndIdx], "")
			if _, ok := f[tokenSearch]; ok {
				legalFormTokenLength += currentTokenLength
				currentTokenLength = 0
			}
		}

		legalFormTokenStart := len(tokens) - legalFormTokenLength - obsoleteTokenLength
		foundLegalForm := legalFormTokenStart <= i
		if foundLegalForm {
			legalForm := strings.Join(tokens[legalFormTokenStart:searchEndIdx], " ")
			if len(legalForm) > 1 {
				company := strings.Join(tokens[0:legalFormTokenStart], " ")
				other := ""
				if obsoleteTokenLength > 0 {
					other = strings.Join(tokens[i+1:], " ")
				}
				return company, legalForm, other
			}
		}
	}
	return strings.Join(tokens, " "), "", ""
}

func clean(s string) string {
	var sb strings.Builder
	sb.Grow(len(s))

	for _, r := range s {
		switch r {
		case '.', '-', '/', '"', '’', '(', ')', '&', '\'', ',', ':', ' ':
			continue
		default:
			sb.WriteRune(r)
		}
	}

	return strings.ToLower(sb.String())
}
