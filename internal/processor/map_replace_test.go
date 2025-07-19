package processor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapReplaceProcessor_Patch(t *testing.T) {
	replacements := map[string]string{
		"foo":   "bar",
		"hello": "world",
	}
	p := NewMapReplaceProcessor(replacements)

	// 正常替换
	input := []byte("foo and hello")
	output, err := p.Patch(input)
	assert.NoError(t, err)
	assert.Equal(t, "bar and world", string(output))

	// 无需替换
	input2 := []byte("no match string")
	output2, err2 := p.Patch(input2)
	assert.Error(t, err2)
	assert.Nil(t, output2)

	// 空输入
	input3 := []byte("")
	output3, err3 := p.Patch(input3)
	assert.Error(t, err3)
	assert.Nil(t, output3)
}
