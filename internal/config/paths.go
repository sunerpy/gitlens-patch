package config

import (
	"path/filepath"
)

// ExtensionPath 扩展路径结构
type ExtensionPath struct {
	Name        string
	Path        string
	Description string
}

// GetPredefinedPaths 获取预定义的扩展路径
func GetPredefinedPaths() []ExtensionPath {
	return []ExtensionPath{
		{"VSCode", filepath.Join(".vscode", "extensions"), "标准 VSCode"},
		{"Cursor", filepath.Join(".cursor", "extensions"), "Cursor 编辑器"},
		{"VSCode Insiders", filepath.Join(".vscode-insiders", "extensions"), "VSCode 预览版"},
		{"Windsurf", filepath.Join(".windsurf", "extensions"), "Windsurf 编辑器"},
	}
}
