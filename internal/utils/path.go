package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/sunerpy/gitlens-patch/internal/config"
)

// PathHandler 路径处理器
type PathHandler struct{}

// NewPathHandler 创建路径处理器实例
func NewPathHandler() *PathHandler {
	return &PathHandler{}
}

// GetExtensionsDir 获取扩展目录
func (ph *PathHandler) GetExtensionsDir() string {
	// 检查命令行参数
	if len(os.Args) > 2 && os.Args[1] == "--ext-dir" {
		return os.Args[2]
	}

	// 检查环境变量
	if envDir := os.Getenv("VSCODE_EXTENSIONS_DIR"); envDir != "" {
		return envDir
	}

	// 获取所有VSCode环境
	environments := config.GetAllVSCodeEnvironments()

	// 检测当前环境
	detectedEdition := config.DetectVSCodeEdition()

	// 显示环境选择界面
	ph.displayEnvironmentSelection(environments, detectedEdition)

	// 获取用户选择
	choice := PromptForSelection(len(environments)+1, "请选择编辑器环境 (输入数字): ")

	if choice == len(environments)+1 {
		return PromptCustomPath()
	}

	return environments[choice-1].ExtensionsDir
}

// displayEnvironmentSelection 显示环境选择界面
func (ph *PathHandler) displayEnvironmentSelection(environments []config.VSCodeEnvironmentInfo, detectedEdition config.VSCodeEdition) {
	fmt.Println("\n" + strings.Repeat(config.SeparatorLine, 80))
	fmt.Println("GitLens Patch - 编辑器环境选择")
	fmt.Println(strings.Repeat(config.SeparatorLine, 80))

	// 显示当前平台信息
	platform := config.GetPlatform()
	fmt.Printf("[INFO] 当前平台: %s\n", ph.getPlatformDisplayName(platform))

	// 显示检测到的环境
	if detectedEdition != "" {
		fmt.Printf("[INFO] 主要检测环境: %s\n", config.VSCODE_BUILTIN_ENVIRONMENTS[detectedEdition].DisplayName)
	}
	fmt.Println()

	// 显示可用环境
	fmt.Println("可用的编辑器环境:")
	fmt.Println(strings.Repeat(config.SeparatorDash, 80))

	for i, env := range environments {
		status := "[ERROR]"
		if env.Exists {
			status = "[OK]"
		}

		// 标记检测到的环境
		marker := ""
		if env.Edition == detectedEdition {
			marker = config.MarkerDetected + " "
		}

		fmt.Printf("[%2d] %s %s%s\n", i+1, status, marker, env.Environment.DisplayName)
		fmt.Printf("     扩展目录: %s\n", env.ExtensionsDir)

		if env.Exists {
			fmt.Printf("     状态: 已安装\n")
		} else {
			fmt.Printf("     状态: 未安装\n")
		}

		// 特殊标记Remote-SSH
		if env.Edition == config.EditionRemoteSSH {
			fmt.Printf("     类型: 远程SSH连接\n")
		}

		fmt.Println()
	}

	fmt.Printf("[%2d] %s 自定义路径\n", len(environments)+1, config.MarkerCustom)
	fmt.Println(strings.Repeat(config.SeparatorDash, 80))

	// 在底部显示检测到的平台信息
	ph.displayDetectionSummary(environments, detectedEdition)
}

// displayDetectionSummary 显示检测摘要
func (ph *PathHandler) displayDetectionSummary(environments []config.VSCodeEnvironmentInfo, detectedEdition config.VSCodeEdition) {
	fmt.Println("\n" + strings.Repeat(config.SeparatorDash, 80))
	fmt.Println("检测摘要:")

	// 获取检测到的环境及其状态
	detectedEnvironments := config.GetDetectedEnvironmentsWithStatus()

	if len(detectedEnvironments) > 0 {
		fmt.Printf("[INFO] 检测到 %d 个环境:\n", len(detectedEnvironments))

		for _, env := range detectedEnvironments {
			// 找到对应的索引
			detectedIndex := -1
			for i, envInfo := range environments {
				if envInfo.Edition == env.Edition {
					detectedIndex = i + 1
					break
				}
			}

			fmt.Printf("  - %s\n", env.Environment.DisplayName)
			fmt.Printf("    扩展目录: %s\n", env.ExtensionsDir)
			fmt.Printf("    状态: %s\n", ph.getStatusText(env.Exists))
			if detectedIndex > 0 {
				fmt.Printf("    对应选项: [%d]\n", detectedIndex)
			}

			// 标记主要检测环境
			if env.Edition == detectedEdition {
				fmt.Printf("    标记: 主要检测环境\n")
			}
			fmt.Println()
		}
	} else {
		fmt.Printf("[WARN] 未检测到特定环境，请手动选择\n")
	}

	// 显示检测依据
	fmt.Println("检测依据:")

	// 检查环境变量
	envVars := []struct {
		name  string
		value string
	}{
		{"VSCODE_SERVER_PATH", os.Getenv("VSCODE_SERVER_PATH")},
		{"CODE_SERVER_VERSION", os.Getenv("CODE_SERVER_VERSION")},
		{"VSCODE_AGENT_FOLDER", os.Getenv("VSCODE_AGENT_FOLDER")},
		{"REMOTE_SSH_EXTENSION", os.Getenv("REMOTE_SSH_EXTENSION")},
		{"VSCODE_SSH_HOST", os.Getenv("VSCODE_SSH_HOST")},
		{"VSCODE_PORTABLE", os.Getenv("VSCODE_PORTABLE")},
	}

	hasEnvVars := false
	for _, envVar := range envVars {
		if envVar.value != "" {
			fmt.Printf("  [INFO] 环境变量: %s=%s\n", envVar.name, envVar.value)
			hasEnvVars = true
		}
	}

	if !hasEnvVars {
		fmt.Printf("  [INFO] 未发现相关环境变量\n")
	}

	// 检查目录
	home, _ := os.UserHomeDir()
	dirsToCheck := []struct {
		name string
		path string
	}{
		{".vscode-server/extensions", filepath.Join(home, ".vscode-server", "extensions")},
		{".vscode-server/data/extensions", filepath.Join(home, ".vscode-server", "data", "extensions")},
		{".cursor/extensions", filepath.Join(home, ".cursor", "extensions")},       // 本地 cursor 客户端
		{"/config/.cursor-server/extensions", "/config/.cursor-server/extensions"}, // Remote-SSH 场景
		{".windsurf/extensions", filepath.Join(home, ".windsurf", "extensions")},
		{".trae/extensions", filepath.Join(home, ".trae", "extensions")},
		{".trae-cn/extensions", filepath.Join(home, ".trae-cn", "extensions")},
		{".vscode-insiders/extensions", filepath.Join(home, ".vscode-insiders", "extensions")},
		{".vscode-oss/extensions", filepath.Join(home, ".vscode-oss", "extensions")},
		{"vscode/extensions", filepath.Join(home, "vscode", "extensions")},
		{"/config/extensions", "/config/extensions"},
		{".local/share/code-server/extensions", filepath.Join(home, ".local/share/code-server/extensions")},
		{".config/code-server/extensions", filepath.Join(home, ".config/code-server/extensions")},
	}

	hasDirs := false
	for _, dir := range dirsToCheck {
		if _, err := os.Stat(dir.path); err == nil {
			fmt.Printf("  [INFO] 扩展目录存在: %s\n", dir.path)
			hasDirs = true
		}
	}

	if !hasDirs {
		fmt.Printf("  [INFO] 未发现相关扩展目录\n")
	}

	fmt.Println(strings.Repeat(config.SeparatorDash, 80))
}

// getPlatformDisplayName 获取平台显示名称
func (ph *PathHandler) getPlatformDisplayName(platform config.Platform) string {
	switch platform {
	case config.PlatformLinux:
		return "Linux"
	case config.PlatformMac:
		return "macOS"
	case config.PlatformWindows:
		return "Windows"
	default:
		return "Unknown"
	}
}

// getStatusText 获取状态文本
func (ph *PathHandler) getStatusText(exists bool) string {
	if exists {
		return "已安装"
	}
	return "未安装"
}
