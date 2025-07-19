package processor

// PatchProcessor 补丁处理器接口
type PatchProcessor interface {
	Patch(content []byte) ([]byte, error)
	GetVersion() int
}
