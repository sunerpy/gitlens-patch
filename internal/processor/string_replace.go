package processor

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/sunerpy/gitlens-patch/internal/config"
)

// StringReplaceProcessor 字符串替换处理器
type StringReplaceProcessor struct {
	*BaseProcessor
}

// NewStringReplaceProcessor 创建字符串替换处理器
func NewStringReplaceProcessor(pattern *regexp.Regexp, replaceStyle string) *StringReplaceProcessor {
	return &StringReplaceProcessor{
		BaseProcessor: NewBaseProcessor(pattern, replaceStyle),
	}
}

// Patch 执行字符串替换操作
func (p *StringReplaceProcessor) Patch(content []byte) ([]byte, error) {
	contentStr := string(content)
	matches := p.GetPattern().FindStringSubmatch(contentStr)

	if len(matches) == 0 {
		return nil, fmt.Errorf(config.ErrNoPatternMatch)
	}

	switch p.GetReplaceStyle() {
	case "prefix":
		exactMatch := matches[0]
		return []byte(strings.Replace(contentStr, exactMatch, p.GetInsertCode()+exactMatch, 1)), nil
	case "replace":
		entireMatch := matches[0]
		return []byte(strings.Replace(contentStr, entireMatch, p.GetInsertCode(), 1)), nil
	default:
		return nil, fmt.Errorf(config.ErrUnknownReplaceStyle)
	}
}
