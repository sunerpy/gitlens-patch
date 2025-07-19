package processor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestV15Processor_Patch(t *testing.T) {
	p := NewV15Processor()

	// 正常替换
	input := []byte("qn.CommunityWithAccount + qn.Community + qn.Pro")
	output, err := p.Patch(input)
	assert.NoError(t, err)
	assert.Equal(t, "qn.Enterprise + qn.Enterprise + qn.Enterprise", string(output))

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
