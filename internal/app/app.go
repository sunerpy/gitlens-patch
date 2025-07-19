package app

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/sunerpy/gitlens-patch/internal/config"
	"github.com/sunerpy/gitlens-patch/internal/processor"
	"github.com/sunerpy/gitlens-patch/internal/utils"
)

// App 主应用结构
type App struct {
	fileProcessor    *utils.FileProcessor
	pathHandler      *utils.PathHandler
	processorFactory *processor.Factory
}

// NewApp 创建应用实例
func NewApp() *App {
	return &App{
		fileProcessor:    utils.NewFileProcessor(),
		pathHandler:      utils.NewPathHandler(),
		processorFactory: processor.NewFactory(),
	}
}

// Run 运行应用
func (app *App) Run() error {
	// 显示欢迎界面
	app.displayWelcome()

	// 获取扩展目录
	extensionsDir := app.pathHandler.GetExtensionsDir()
	app.displayExtensionsDir(extensionsDir)

	// 获取GitLens路径
	extensionPath, err := app.fileProcessor.GetGitLensPath(extensionsDir)
	if err != nil {
		return fmt.Errorf("查找 GitLens 扩展失败: %v", err)
	}

	// 解析版本号
	majorVersion, err := app.fileProcessor.ParseMajorVersion(filepath.Base(extensionPath))
	if err != nil {
		return fmt.Errorf("无法解析版本号: %v", err)
	}

	// 显示版本信息
	app.displayVersionInfo(majorVersion, extensionPath)

	// 创建处理器
	patchProcessor, err := app.processorFactory.CreateProcessor(majorVersion)
	if err != nil {
		return err
	}

	app.displayProcessorInfo(patchProcessor.GetVersion())

	// 定义需要修改的文件
	filesToModify := []string{
		filepath.Join(extensionPath, config.GitLensJSPath),
		filepath.Join(extensionPath, config.GitLensBrowserPath),
	}

	// 处理文件
	app.displayProcessingStart()
	successCount := 0
	for _, file := range filesToModify {
		if err := app.fileProcessor.ProcessFile(file, patchProcessor); err != nil {
			fmt.Printf("[ERROR] 处理文件 %s 失败: %v\n", filepath.Base(file), err)
		} else {
			successCount++
		}
	}

	// 显示完成信息
	app.displayCompletion(successCount, len(filesToModify))
	return nil
}

// displayWelcome 显示欢迎界面
func (app *App) displayWelcome() {
	fmt.Println(strings.Repeat(config.SeparatorLine, 80))
	fmt.Println("GitLens Patch - GitLens Pro 激活工具")
	fmt.Println(strings.Repeat(config.SeparatorLine, 80))
	fmt.Println("功能特性:")
	fmt.Println("  [OK] 支持 GitLens 15、16、17 版本")
	fmt.Println("  [OK] 自动检测编辑器环境")
	fmt.Println("  [OK] 自动备份原文件")
	fmt.Println("  [OK] 版本兼容性检查")
	fmt.Println("  [OK] 支持多种编辑器")
	fmt.Println(strings.Repeat(config.SeparatorLine, 80))
	fmt.Println()
}

// displayExtensionsDir 显示扩展目录信息
func (app *App) displayExtensionsDir(extensionsDir string) {
	fmt.Println("[INFO] 正在检测扩展目录...")
	fmt.Printf("[INFO] 扩展目录: %s\n", extensionsDir)
	fmt.Println()
}

// displayVersionInfo 显示版本信息
func (app *App) displayVersionInfo(majorVersion int, extensionPath string) {
	fmt.Println("[INFO] 检测到 GitLens 扩展:")
	fmt.Printf("  版本: %d.x\n", majorVersion)
	fmt.Printf("  路径: %s\n", extensionPath)

	// 检查版本支持
	if majorVersion >= config.MinSupportedVersion && majorVersion <= config.MaxSupportedVersion {
		fmt.Printf("  版本支持: [OK] (支持范围: %d-%d)\n", config.MinSupportedVersion, config.MaxSupportedVersion)
	} else {
		fmt.Printf("  版本支持: [ERROR] (支持范围: %d-%d)\n", config.MinSupportedVersion, config.MaxSupportedVersion)
	}
	fmt.Println()
}

// displayProcessorInfo 显示处理器信息
func (app *App) displayProcessorInfo(version int) {
	fmt.Printf("[INFO] 使用处理器: v%d 版本处理器\n", version)
	fmt.Println()
}

// displayProcessingStart 显示处理开始信息
func (app *App) displayProcessingStart() {
	fmt.Println("[INFO] 开始处理文件...")
	fmt.Println(strings.Repeat(config.SeparatorDash, 80))
}

// displayCompletion 显示完成信息
func (app *App) displayCompletion(successCount, totalCount int) {
	fmt.Println(strings.Repeat(config.SeparatorDash, 80))
	fmt.Println("[INFO] 处理完成!")
	fmt.Printf("  成功处理: %d/%d 个文件\n", successCount, totalCount)

	if successCount == totalCount {
		fmt.Println("  状态: [OK] 所有文件处理成功!")
	} else if successCount > 0 {
		fmt.Println("  状态: [WARN] 部分文件处理成功，请检查失败的文件")
	} else {
		fmt.Println("  状态: [ERROR] 所有文件处理失败")
	}

	fmt.Println()
	fmt.Println("下一步操作:")
	fmt.Println("  1. 重启您的编辑器")
	fmt.Println("  2. 检查 GitLens Pro 功能是否已激活")
	fmt.Println("  3. 如有问题，可使用 'restore' 命令恢复原文件")
	fmt.Println()
}

// Restore 恢复功能
func Restore() error {
	// TODO: 实现恢复逻辑
	return nil
}

// WaitForKeyPress 等待用户按键
func WaitForKeyPress() {
	utils.WaitForKeyPress()
}
