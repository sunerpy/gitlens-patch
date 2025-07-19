package processor

import (
	"fmt"

	"github.com/sunerpy/gitlens-patch/internal/config"
)

// Factory 处理器工厂
type Factory struct{}

// NewFactory 创建处理器工厂实例
func NewFactory() *Factory {
	return &Factory{}
}

// CreateProcessor 根据版本创建对应的处理器
func (f *Factory) CreateProcessor(majorVersion int) (PatchProcessor, error) {
	// 检查版本是否支持
	if majorVersion < config.MinSupportedVersion || majorVersion > config.MaxSupportedVersion {
		return nil, fmt.Errorf("%s (当前版本: %d)", config.ErrVersionNotSupported, majorVersion)
	}

	switch majorVersion {
	case 15:
		return NewV15Processor(), nil
	case 16:
		return NewV16PlusProcessor(), nil
	case 17:
		return NewV17Processor(), nil
	default:
		return nil, fmt.Errorf("%s (当前版本: %d)", config.ErrVersionNotSupported, majorVersion)
	}
}
