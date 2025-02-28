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
			input:               "Example llc Oy GmbH & Co. KG",
			expectedCompanyName: "Example llc Oy",
			expectedLegalForm:   "GmbH & Co. KG",
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
		{
			input:               "Example",
			expectedCompanyName: "Example",
			expectedLegalForm:   "",
		},
		{
			input:               "Example GmbH (Foobar)",
			expectedCompanyName: "Example GmbH (Foobar)",
			expectedLegalForm:   "",
		},
		{
			input:               "Example GmbH & Co. KG Foobar",
			expectedCompanyName: "Example GmbH & Co. KG Foobar",
			expectedLegalForm:   "",
		},
		{
			input:               "LLC Example",
			expectedCompanyName: "LLC Example",
			expectedLegalForm:   "",
		},
		{
			input:               "Foo A Example",
			expectedCompanyName: "Foo A Example",
			expectedLegalForm:   "",
		},
		{
			input:               "LLC",
			expectedCompanyName: "LLC",
			expectedLegalForm:   "",
		},
		{
			input:               "",
			expectedCompanyName: "",
			expectedLegalForm:   "",
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

func TestStripMiddle(t *testing.T) {
	cases := []struct {
		input               string
		expectedCompanyName string
		expectedLegalForm   string
		expectedOther       string
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
		{
			input:               "Example",
			expectedCompanyName: "Example",
			expectedLegalForm:   "",
		},
		{
			input:               "Example GmbH (Foobar)",
			expectedCompanyName: "Example",
			expectedLegalForm:   "GmbH",
			expectedOther:       "(Foobar)",
		},
		{
			input:               "Example GmbH & Co. KG Foobar",
			expectedCompanyName: "Example",
			expectedLegalForm:   "GmbH & Co. KG",
			expectedOther:       "Foobar",
		},
		{
			input:               "Example LLC GmbH & Co. KG Some Street Name No 1",
			expectedCompanyName: "Example LLC",
			expectedLegalForm:   "GmbH & Co. KG",
			expectedOther:       "Some Street Name No 1",
		},
		{
			input:               "LLC Example",
			expectedCompanyName: "LLC Example",
			expectedLegalForm:   "",
		},
		{
			input:               "Foo A Example",
			expectedCompanyName: "Foo A Example",
			expectedLegalForm:   "",
		},
		{
			input:               "",
			expectedCompanyName: "",
			expectedLegalForm:   "",
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("#%v", i), func(t *testing.T) {
			actualCompany, actualLegalForm, actualOther := legalform.Default.StripMiddle(c.input)
			assert.Equal(t, c.expectedCompanyName, actualCompany)
			assert.Equal(t, c.expectedLegalForm, actualLegalForm)
			assert.Equal(t, c.expectedOther, actualOther)
		})
	}
}

func TestStripThenAlias(t *testing.T) {
	cases := []struct {
		input                  string
		country                string
		expectedCompanyName    string
		expectedLegalForm      string
		expectedAliasLegalForm string
	}{
		{
			input:                  "Example Shipping Private Limited",
			country:                "XX",
			expectedCompanyName:    "Example Shipping",
			expectedLegalForm:      "Private Limited",
			expectedAliasLegalForm: "pvtltd",
		},
		{
			input:                  "Example Shipping Pvt Limited",
			country:                "XX",
			expectedCompanyName:    "Example Shipping",
			expectedLegalForm:      "Pvt Limited",
			expectedAliasLegalForm: "pvtltd",
		},
		{
			input:                  "Example Shipping Pvt Ltd",
			country:                "XX",
			expectedCompanyName:    "Example Shipping",
			expectedLegalForm:      "Pvt Ltd",
			expectedAliasLegalForm: "pvtltd",
		},
		{
			input:                  "Example Shipping Private Ltd",
			country:                "XX",
			expectedCompanyName:    "Example Shipping",
			expectedLegalForm:      "Private Ltd",
			expectedAliasLegalForm: "pvtltd",
		},
		{
			input:                  "The example Corp. kk",
			country:                "JP",
			expectedCompanyName:    "The example",
			expectedLegalForm:      "Corp. kk",
			expectedAliasLegalForm: "kk",
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("#%v", i), func(t *testing.T) {
			actualCompany, actualLegalForm := legalform.Default.Strip(c.input)
			actualAliasLegalForm := legalform.DefaultAliases.Find(c.country, actualLegalForm)
			assert.Equal(t, c.expectedCompanyName, actualCompany)
			assert.Equal(t, c.expectedLegalForm, actualLegalForm)
			assert.Equal(t, c.expectedAliasLegalForm, actualAliasLegalForm)
		})
	}
}

func BenchmarkStrip(b *testing.B) {
	input := "Some Example S. A. de C. V., F. I. en I. D."
	for i := 0; i < b.N; i++ {
		_, _ = legalform.Default.Strip(input)
	}
}

func BenchmarkStripMiddle(b *testing.B) {
	input := "Example LLC GmbH & Co. KG Some Street Name No 1"
	for i := 0; i < b.N; i++ {
		_, _, _ = legalform.Default.StripMiddle(input)
	}
}
