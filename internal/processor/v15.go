package processor

// V15Processor v15版本处理器
type V15Processor struct {
	*MapReplaceProcessor
}

// NewV15Processor 创建v15处理器实例
func NewV15Processor() *V15Processor {
	replacements := []struct{ old, new string }{
		{"qn.CommunityWithAccount", "qn.Enterprise"},
		{"qn.Community", "qn.Enterprise"},
		{"qn.Pro", "qn.Enterprise"},
	}
	return &V15Processor{
		MapReplaceProcessor: NewMapReplaceProcessorOrdered(replacements),
	}
}

// GetVersion 获取处理器版本
func (p *V15Processor) GetVersion() int {
	return 15
}
