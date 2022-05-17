package legalform_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	legalform "github.com/tilotech/go-company-legal-form"
)

func TestStrip(t *testing.T) {
	cases := []struct {
		input               string
		expectedCompanyName string
		expectedLegalForm   string
	}{
		{
			input:               "Example LLC",
			expectedCompanyName: "Example",
			expectedLegalForm:   "LLC",
		},
		{
			input:               "Example GmbH & Co. KG",
			expectedCompanyName: "Example",
			expectedLegalForm:   "GmbH & Co. KG",
		},
		{
			input:               "llc Oy AG",
			expectedCompanyName: "llc Oy",
			expectedLegalForm:   "AG",
		},
		{
			input:               "Example KG",
			expectedCompanyName: "Example",
			expectedLegalForm:   "KG",
		},
		{
			input:               "Example e.V.",
			expectedCompanyName: "Example",
			expectedLegalForm:   "e.V.",
		},
		{
			input:               "Some Example S. A. de C. V., F. I. en I. D.",
			expectedCompanyName: "Some Example",
			expectedLegalForm:   "S. A. de C. V., F. I. en I. D.",
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("#%v", i), func(t *testing.T) {
			actualCompany, actualLegalForm := legalform.Default.Strip(c.input)
			assert.Equal(t, c.expectedCompanyName, actualCompany)
			assert.Equal(t, c.expectedLegalForm, actualLegalForm)
		})
	}
}
