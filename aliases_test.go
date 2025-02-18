package legalform_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	legalform "github.com/tilotech/go-company-legal-form"
)

func TestFindAlias(t *testing.T) {
	actual := legalform.DefaultAliases.Find("DE", "Gesellschaft mit beschränkter Haftung")
	assert.Equal(t, "gmbh", actual)
	actual = legalform.DefaultAliases.Find("XX", "Gesellschaft mit beschränkter Haftung")
	assert.Equal(t, "gesellschaftmitbeschränkterhaftung", actual)
	actual = legalform.DefaultAliases.Find("US", "Incorporated")
	assert.Equal(t, "inc", actual)
	actual = legalform.DefaultAliases.Find("US", "Gesellschaft mit beschränkter Haftung")
	assert.Equal(t, "gesellschaftmitbeschränkterhaftung", actual)
	actual = legalform.DefaultAliases.Find("XX", "Incorporated")
	assert.Equal(t, "inc", actual)
	actual = legalform.DefaultAliases.Find("DE", "Incorporated")
	assert.Equal(t, "inc", actual)
}
