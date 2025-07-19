package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
)

// FileProcessor 文件处理器
type FileProcessor struct{}

// NewFileProcessor 创建文件处理器实例
func NewFileProcessor() *FileProcessor {
	return &FileProcessor{}
}

// ProcessFile 处理单个文件
func (fp *FileProcessor) ProcessFile(filePath string, processor interface {
	Patch(content []byte) ([]byte, error)
}) error {
	fileName := filepath.Base(filePath)

	// 检查文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return fmt.Errorf("文件不存在: %s", fileName)
	}

	// 读取文件内容
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("读取文件失败: %v", err)
	}

	// 创建备份
	backupPath := filePath + ".backup"
	if _, err := os.Stat(backupPath); os.IsNotExist(err) {
		if err := os.WriteFile(backupPath, content, 0644); err != nil {
			return fmt.Errorf("创建备份失败: %v", err)
		}
		fmt.Printf("  [INFO] 已创建备份: %s.backup\n", fileName)
	} else {
		fmt.Printf("  [INFO] 备份已存在: %s.backup\n", fileName)
	}

	// 处理文件内容
	newContent, err := processor.Patch(content)
	if err != nil {
		return err
	}

	// 写入新内容
	if err := os.WriteFile(filePath, newContent, 0644); err != nil {
		return fmt.Errorf("写入文件失败: %v", err)
	}

	fmt.Printf("  [OK] 成功修改: %s\n", fileName)
	return nil
}

// ParseMajorVersion 解析主版本号
func (fp *FileProcessor) ParseMajorVersion(dirName string) (int, error) {
	pattern := regexp.MustCompile(`eamodio\.gitlens-(\d+)\.`)
	matches := pattern.FindStringSubmatch(dirName)
	if len(matches) < 2 {
		return 0, fmt.Errorf("无法解析主版本号")
	}
	return strconv.Atoi(matches[1])
}

// GetGitLensPath 获取GitLens扩展路径
func (fp *FileProcessor) GetGitLensPath(extensionsDir string) (string, error) {
	entries, err := os.ReadDir(extensionsDir)
	if err != nil {
		return "", fmt.Errorf("读取扩展目录失败: %v", err)
	}

	pattern := regexp.MustCompile(`^eamodio\.gitlens-\d+\.\d+\.\d+`)
	var gitLensDirs []string

	for _, entry := range entries {
		if entry.IsDir() && pattern.MatchString(entry.Name()) {
			gitLensDirs = append(gitLensDirs, entry.Name())
		}
	}

	if len(gitLensDirs) == 0 {
		return "", fmt.Errorf("未找到 GitLens 扩展")
	}

	// 按版本排序（降序）
	sort.Sort(sort.Reverse(sort.StringSlice(gitLensDirs)))

	if len(gitLensDirs) == 1 {
		return filepath.Join(extensionsDir, gitLensDirs[0]), nil
	}

	// 显示多个版本供用户选择
	fmt.Println("[INFO] 发现多个 GitLens 版本:")
	for i, dir := range gitLensDirs {
		fmt.Printf("  [%d] %s\n", i+1, dir)
	}

	choice := PromptForSelection(len(gitLensDirs), "请选择要激活的 GitLens 版本 (输入数字): ")
	return filepath.Join(extensionsDir, gitLensDirs[choice-1]), nil
}
