package processor

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringReplaceProcessor_Patch_Prefix(t *testing.T) {
	pattern := regexp.MustCompile(`let [a-zA-Z]={id:e\.user\.id,name:`)
	p := NewStringReplaceProcessor(pattern, "prefix")
	p.insertCode = "//inserted;"

	input := []byte("let a={id:e.user.id,name:")
	output, err := p.Patch(input)
	assert.NoError(t, err)
	assert.Contains(t, string(output), "//inserted;")
}

func TestStringReplaceProcessor_Patch_Replace(t *testing.T) {
	pattern := regexp.MustCompile(`let [a-zA-Z]={id:e\.user\.id,name:`)
	p := NewStringReplaceProcessor(pattern, "replace")
	p.insertCode = "//only;"

	input := []byte("let a={id:e.user.id,name:xxx}")
	output, err := p.Patch(input)
	assert.NoError(t, err)
	assert.Equal(t, "//only;xxx}", string(output))
}

func TestStringReplaceProcessor_Patch_NoMatch(t *testing.T) {
	pattern := regexp.MustCompile(`let [a-zA-Z]={id:e\.user\.id,name:`)
	p := NewStringReplaceProcessor(pattern, "prefix")
	input := []byte("no match string")
	output, err := p.Patch(input)
	assert.Error(t, err)
	assert.Nil(t, output)
}
