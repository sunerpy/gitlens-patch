package processor

import (
	"regexp"
)

type V17Processor struct {
	*ComplexReplaceProcessor
}

func NewV17Processor() *V17Processor {
	pat := regexp.MustCompile(
		`let\s+([a-zA-Z0-9_$,\s]*)\,\s*n\s*=\s*(\{id:e\.user\.id,name:)`)

	return &V17Processor{
		ComplexReplaceProcessor: NewComplexReplaceProcessor(pat, "prefix"),
	}
}

func (*V17Processor) GetVersion() int { return 17 }
