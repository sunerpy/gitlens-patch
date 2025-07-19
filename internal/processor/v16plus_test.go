package processor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestV16PlusProcessor_Patch_Prefix(t *testing.T) {
	p := NewV16PlusProcessor()
	p.replaceStyle = "prefix"
	p.insertCode = "//inserted;"

	input := []byte("let a={id:e.user.id,name:xxx}")
	output, err := p.Patch(input)
	assert.NoError(t, err)
	assert.Contains(t, string(output), "//inserted;")
}

func TestV16PlusProcessor_Patch_Replace(t *testing.T) {
	p := NewV16PlusProcessor()
	p.replaceStyle = "replace"
	p.insertCode = "//only;"

	input := []byte("let a={id:e.user.id,name:xxx}")
	output, err := p.Patch(input)
	assert.NoError(t, err)
	assert.Equal(t, "//only;xxx}", string(output))
}

func TestV16PlusProcessor_Patch_NoMatch(t *testing.T) {
	p := NewV16PlusProcessor()
	input := []byte("no match string")
	output, err := p.Patch(input)
	assert.Error(t, err)
	assert.Nil(t, output)
}
