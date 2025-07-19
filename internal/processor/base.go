package processor

import (
	"regexp"

	"github.com/sunerpy/gitlens-patch/internal/config"
)

// BaseProcessor 基础处理器，包含所有处理器共用的字段
type BaseProcessor struct {
	pattern      *regexp.Regexp
	insertCode   string
	replaceStyle string
}

// NewBaseProcessor 创建基础处理器
func NewBaseProcessor(pattern *regexp.Regexp, replaceStyle string) *BaseProcessor {
	return &BaseProcessor{
		pattern:      pattern,
		insertCode:   config.InsertCode,
		replaceStyle: replaceStyle,
	}
}

// GetPattern 获取正则表达式模式
func (p *BaseProcessor) GetPattern() *regexp.Regexp {
	return p.pattern
}

// GetInsertCode 获取插入代码
func (p *BaseProcessor) GetInsertCode() string {
	return p.insertCode
}

// GetReplaceStyle 获取替换样式
func (p *BaseProcessor) GetReplaceStyle() string {
	return p.replaceStyle
}
