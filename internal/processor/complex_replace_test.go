package processor

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComplexReplaceProcessor_Patch(t *testing.T) {
	pattern := regexp.MustCompile(`let\s+([a-zA-Z0-9_$,\s]*)\,\s*n\s*=\s*(\{id:e\.user\.id,name:)`)
	p := NewComplexReplaceProcessor(pattern, "prefix")
	p.insertCode = "//inserted;"

	input := []byte("let a, n={id:e.user.id,name:")
	output, err := p.Patch(input)
	assert.NoError(t, err)
	assert.Contains(t, string(output), "//inserted;")
	assert.Contains(t, string(output), "let n={id:e.user.id,name:")
}

func TestComplexReplaceProcessor_Patch_NoMatch(t *testing.T) {
	pattern := regexp.MustCompile(`let\s+([a-zA-Z0-9_$,\s]*)\,\s*n\s*=\s*(\{id:e\.user\.id,name:)`)
	p := NewComplexReplaceProcessor(pattern, "prefix")
	input := []byte("no match string")
	output, err := p.Patch(input)
	assert.Error(t, err)
	assert.Nil(t, output)
}
