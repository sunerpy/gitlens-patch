package processor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestV17Processor_Patch(t *testing.T) {
	p := NewV17Processor()
	p.insertCode = "//inserted;"

	input := []byte("let a, n={id:e.user.id,name:xxx}")
	output, err := p.Patch(input)
	assert.NoError(t, err)
	assert.Contains(t, string(output), "//inserted;")
	assert.Contains(t, string(output), "let n={id:e.user.id,name:")
}

func TestV17Processor_Patch_NoMatch(t *testing.T) {
	p := NewV17Processor()
	input := []byte("no match string")
	output, err := p.Patch(input)
	assert.Error(t, err)
	assert.Nil(t, output)
}
