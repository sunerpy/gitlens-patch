package processor

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/sunerpy/gitlens-patch/internal/config"
)

// ComplexReplaceProcessor 复杂替换处理器
type ComplexReplaceProcessor struct {
	*BaseProcessor
}

// NewComplexReplaceProcessor 创建复杂替换处理器
func NewComplexReplaceProcessor(pattern *regexp.Regexp, replaceStyle string) *ComplexReplaceProcessor {
	return &ComplexReplaceProcessor{
		BaseProcessor: NewBaseProcessor(pattern, replaceStyle),
	}
}

// Patch 执行复杂替换操作
func (p *ComplexReplaceProcessor) Patch(content []byte) ([]byte, error) {
	src := string(content)
	indices := p.GetPattern().FindStringSubmatchIndex(src)
	if indices == nil {
		return nil, fmt.Errorf(config.ErrNoPatternMatch)
	}

	vars := strings.TrimSpace(src[indices[2]:indices[3]])
	objStart := src[indices[4]:indices[5]]

	var builder strings.Builder
	builder.WriteString("let ")
	builder.WriteString(vars)
	builder.WriteString(";")

	clean := strings.TrimSpace(p.GetInsertCode())
	builder.WriteString(clean)
	if !strings.HasSuffix(clean, ";") {
		builder.WriteString(";")
	}
	builder.WriteString("let n=")
	builder.WriteString(objStart)

	newSrc := src[:indices[0]] + builder.String() + src[indices[5]:]
	return []byte(newSrc), nil
}
