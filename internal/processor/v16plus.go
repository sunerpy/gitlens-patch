package processor

import (
	"regexp"
)

// V16PlusProcessor v16版本处理器
type V16PlusProcessor struct {
	*StringReplaceProcessor
}

// NewV16PlusProcessor 创建v16处理器实例
func NewV16PlusProcessor() *V16PlusProcessor {
	return &V16PlusProcessor{
		StringReplaceProcessor: NewStringReplaceProcessor(
			regexp.MustCompile(`let [a-zA-Z]={id:e\.user\.id,name:`),
			"prefix",
		),
	}
}

// GetVersion 获取处理器版本
func (p *V16PlusProcessor) GetVersion() int {
	return 16
}
