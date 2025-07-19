package processor

import (
	"fmt"
	"strings"

	"github.com/sunerpy/gitlens-patch/internal/config"
)

// MapReplaceProcessor 批量字符串替换处理器
type MapReplaceProcessor struct {
	*BaseProcessor
	replacements map[string]string
	order        []string // 新增字段，支持有序替换
}

// NewMapReplaceProcessor 创建批量字符串替换处理器
func NewMapReplaceProcessor(replacements map[string]string) *MapReplaceProcessor {
	return &MapReplaceProcessor{
		BaseProcessor: NewBaseProcessor(nil, "map"),
		replacements:  replacements,
	}
}

// NewMapReplaceProcessorOrdered 支持有序替换
func NewMapReplaceProcessorOrdered(replacements []struct{ old, new string }) *MapReplaceProcessor {
	m := make(map[string]string)
	order := make([]string, 0, len(replacements))
	for _, kv := range replacements {
		m[kv.old] = kv.new
		order = append(order, kv.old)
	}
	return &MapReplaceProcessor{
		BaseProcessor: NewBaseProcessor(nil, "map"),
		replacements:  m,
		order:         order,
	}
}

// Patch 执行批量字符串替换
func (p *MapReplaceProcessor) Patch(content []byte) ([]byte, error) {
	contentStr := string(content)
	modified := false

	// 有序替换
	if len(p.order) > 0 {
		for _, old := range p.order {
			newVal := p.replacements[old]
			if strings.Contains(contentStr, old) {
				contentStr = strings.ReplaceAll(contentStr, old, newVal)
				modified = true
			}
		}
	} else {
		for old, newVal := range p.replacements {
			if strings.Contains(contentStr, old) {
				contentStr = strings.ReplaceAll(contentStr, old, newVal)
				modified = true
			}
		}
	}

	if !modified {
		return nil, fmt.Errorf(config.ErrNoContentToReplace)
	}

	return []byte(contentStr), nil
}
