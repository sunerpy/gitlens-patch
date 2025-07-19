package config

import (
	"os"
	"path/filepath"
	"runtime"
)

// VSCodeEdition VSCode版本类型
type VSCodeEdition string

const (
	EditionStandard         VSCodeEdition = "standard"
	EditionInsiders         VSCodeEdition = "insiders"
	EditionExploration      VSCodeEdition = "exploration"
	EditionVSCodium         VSCodeEdition = "vscodium"
	EditionVSCodiumInsiders VSCodeEdition = "vscodium-insiders"
	EditionOSS              VSCodeEdition = "oss"
	EditionCoder            VSCodeEdition = "coder"
	EditionCodeServer       VSCodeEdition = "codeserver"
	EditionCursor           VSCodeEdition = "cursor"
	EditionWindSurf         VSCodeEdition = "windsurf"
	EditionTrae             VSCodeEdition = "trae"
	EditionTraeCN           VSCodeEdition = "trae-cn"
	EditionRemoteSSH        VSCodeEdition = "remote-ssh"
	EditionCursorServer     VSCodeEdition = "cursor-server"
)

// Platform 平台类型
type Platform string

const (
	PlatformLinux   Platform = "linux"
	PlatformMac     Platform = "mac"
	PlatformWindows Platform = "windows"
)

// VSCodeEnvironment VSCode环境配置
type VSCodeEnvironment struct {
	DataDirectoryName       string
	ExtensionsDirectoryName string
	DisplayName             string
	SortOrder               int // 排序顺序
}

// VSCODE_BUILTIN_ENVIRONMENTS 内置环境配置
var VSCODE_BUILTIN_ENVIRONMENTS = map[VSCodeEdition]VSCodeEnvironment{
	EditionStandard: {
		DataDirectoryName:       "Code",
		ExtensionsDirectoryName: ".vscode",
		DisplayName:             "Visual Studio Code",
		SortOrder:               1,
	},
	EditionInsiders: {
		DataDirectoryName:       "Code - Insiders",
		ExtensionsDirectoryName: ".vscode-insiders",
		DisplayName:             "Visual Studio Code - Insiders",
		SortOrder:               2,
	},
	EditionExploration: {
		DataDirectoryName:       "Code - Exploration",
		ExtensionsDirectoryName: ".vscode-exploration",
		DisplayName:             "Visual Studio Code - Exploration",
		SortOrder:               3,
	},
	EditionVSCodium: {
		DataDirectoryName:       "VSCodium",
		ExtensionsDirectoryName: ".vscode-oss",
		DisplayName:             "VSCodium",
		SortOrder:               4,
	},
	EditionVSCodiumInsiders: {
		DataDirectoryName:       "VSCodium - Insiders",
		ExtensionsDirectoryName: ".vscodium-insiders",
		DisplayName:             "VSCodium - Insiders",
		SortOrder:               5,
	},
	EditionOSS: {
		DataDirectoryName:       "Code - OSS",
		ExtensionsDirectoryName: ".vscode-oss",
		DisplayName:             "Code - OSS",
		SortOrder:               6,
	},
	EditionCursor: {
		DataDirectoryName:       "Cursor",
		ExtensionsDirectoryName: ".cursor",
		DisplayName:             "Cursor",
		SortOrder:               7,
	},
	EditionCursorServer: {
		DataDirectoryName:       ".cursor-server",
		ExtensionsDirectoryName: ".cursor-server",
		DisplayName:             "Cursor Server",
		SortOrder:               8,
	},
	EditionWindSurf: {
		DataDirectoryName:       "WindSurf",
		ExtensionsDirectoryName: ".windsurf",
		DisplayName:             "WindSurf",
		SortOrder:               8,
	},
	EditionTrae: {
		DataDirectoryName:       "Trae",
		ExtensionsDirectoryName: ".trae",
		DisplayName:             "Trae",
		SortOrder:               9,
	},
	EditionTraeCN: {
		DataDirectoryName:       "Trae CN",
		ExtensionsDirectoryName: ".trae-cn",
		DisplayName:             "Trae CN",
		SortOrder:               10,
	},
	EditionCoder: {
		DataDirectoryName:       "Code",
		ExtensionsDirectoryName: "vscode",
		DisplayName:             "Coder",
		SortOrder:               11,
	},
	EditionCodeServer: {
		DataDirectoryName:       "../.local/share/code-server",
		ExtensionsDirectoryName: ".local/share/code-server",
		DisplayName:             "Code Server",
		SortOrder:               12,
	},
	EditionRemoteSSH: {
		DataDirectoryName:       ".vscode-server",
		ExtensionsDirectoryName: ".vscode-server",
		DisplayName:             "Remote SSH",
		SortOrder:               13,
	},
}

// GetPlatform 获取当前平台
func GetPlatform() Platform {
	switch runtime.GOOS {
	case "linux":
		return PlatformLinux
	case "darwin":
		return PlatformMac
	case "windows":
		return PlatformWindows
	default:
		return PlatformLinux
	}
}

// GetVSCodeDataDirectory 获取VSCode数据目录
func GetVSCodeDataDirectory(edition VSCodeEdition) string {
	// 检查便携模式
	if portable := os.Getenv("VSCODE_PORTABLE"); portable != "" {
		return filepath.Join(portable, "user-data")
	}

	platform := GetPlatform()
	env := VSCODE_BUILTIN_ENVIRONMENTS[edition]

	// 特殊处理Remote-SSH
	if edition == EditionRemoteSSH {
		home, _ := os.UserHomeDir()
		remoteSSHPaths := []string{
			filepath.Join(home, ".vscode-server", "data", "User"),
			filepath.Join(home, ".vscode-server", "User"),
		}

		for _, sshPath := range remoteSSHPaths {
			if _, err := os.Stat(filepath.Join(sshPath, "settings.json")); err == nil {
				return filepath.Dir(sshPath)
			}
		}
		return filepath.Join(home, ".vscode-server", "data")
	}

	// 特殊处理code-server
	if edition == EditionCodeServer {
		codeServerPaths := []string{
			"/config/data/User",
			filepath.Join(os.Getenv("HOME"), ".local/share/code-server/User"),
			filepath.Join(os.Getenv("HOME"), ".config/code-server/User"),
		}

		for _, serverPath := range codeServerPaths {
			if _, err := os.Stat(filepath.Join(serverPath, "settings.json")); err == nil {
				return filepath.Dir(serverPath)
			}
		}
		return "/config/data"
	}

	switch platform {
	case PlatformWindows:
		appData := os.Getenv("APPDATA")
		if appData == "" {
			appData = filepath.Join(os.Getenv("USERPROFILE"), "AppData", "Roaming")
		}
		return filepath.Join(appData, env.DataDirectoryName)
	case PlatformMac:
		home, _ := os.UserHomeDir()
		return filepath.Join(home, "Library", "Application Support", env.DataDirectoryName)
	case PlatformLinux:
		fallthrough
	default:
		home, _ := os.UserHomeDir()
		return filepath.Join(home, ".config", env.DataDirectoryName)
	}
}

// GetVSCodeExtensionsDirectory 获取VSCode扩展目录
func GetVSCodeExtensionsDirectory(edition VSCodeEdition) string {
	// 检查便携模式
	if portable := os.Getenv("VSCODE_PORTABLE"); portable != "" {
		return filepath.Join(portable, "extensions")
	}

	// 特殊处理Cursor-Server
	if edition == EditionCursorServer {
		return "/config/.cursor-server/extensions"
	}

	// 特殊处理Remote-SSH
	if edition == EditionRemoteSSH {
		home, _ := os.UserHomeDir()
		remoteSSHExtPaths := []string{
			filepath.Join(home, ".vscode-server", "extensions"),
			filepath.Join(home, ".vscode-server", "data", "extensions"),
		}

		for _, extPath := range remoteSSHExtPaths {
			if _, err := os.Stat(extPath); err == nil {
				return extPath
			}
		}
		return filepath.Join(home, ".vscode-server", "extensions")
	}

	// 特殊处理code-server
	if edition == EditionCodeServer {
		codeServerExtPaths := []string{
			"/config/extensions",
			filepath.Join(os.Getenv("HOME"), ".local/share/code-server/extensions"),
			filepath.Join(os.Getenv("HOME"), ".config/code-server/extensions"),
		}

		for _, extPath := range codeServerExtPaths {
			if _, err := os.Stat(extPath); err == nil {
				return extPath
			}
		}
		return "/config/extensions"
	}

	env := VSCODE_BUILTIN_ENVIRONMENTS[edition]
	home, _ := os.UserHomeDir()
	return filepath.Join(home, env.ExtensionsDirectoryName, "extensions")
}

// DetectVSCodeEdition 检测当前VSCode版本
func DetectVSCodeEdition() VSCodeEdition {
	// 检查环境变量来检测code-server（优先级最高）
	if os.Getenv("VSCODE_SERVER_PATH") != "" ||
		os.Getenv("CODE_SERVER_VERSION") != "" {
		return EditionCodeServer
	}

	// 检查环境变量来检测Remote-SSH
	if os.Getenv("VSCODE_AGENT_FOLDER") != "" ||
		os.Getenv("REMOTE_SSH_EXTENSION") != "" ||
		os.Getenv("VSCODE_SSH_HOST") != "" {
		return EditionRemoteSSH
	}

	// 检查常见的扩展目录来推断版本
	home, _ := os.UserHomeDir()

	// 检查Remote-SSH目录
	if _, err := os.Stat(filepath.Join(home, ".vscode-server")); err == nil {
		return EditionRemoteSSH
	}

	// 检查Cursor
	if _, err := os.Stat(filepath.Join(home, ".cursor")); err == nil {
		return EditionCursor
	}

	// 检查Cursor Server
	if _, err := os.Stat(filepath.Join(home, ".cursor-server")); err == nil {
		return EditionCursorServer
	}

	// 检查WindSurf
	if _, err := os.Stat(filepath.Join(home, ".windsurf")); err == nil {
		return EditionWindSurf
	}

	// 检查Trae
	if _, err := os.Stat(filepath.Join(home, ".trae")); err == nil {
		return EditionTrae
	}

	// 检查Trae CN
	if _, err := os.Stat(filepath.Join(home, ".trae-cn")); err == nil {
		return EditionTraeCN
	}

	// 检查VSCode Insiders
	if _, err := os.Stat(filepath.Join(home, ".vscode-insiders")); err == nil {
		return EditionInsiders
	}

	// 检查VSCodium
	if _, err := os.Stat(filepath.Join(home, ".vscode-oss")); err == nil {
		return EditionVSCodium
	}

	// 默认返回标准版本
	return EditionStandard
}

// contains 检查切片是否包含指定元素
func contains(slice []VSCodeEdition, item VSCodeEdition) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// GetAllVSCodeEnvironments 获取所有可用的VSCode环境（按固定顺序排序）
func GetAllVSCodeEnvironments() []VSCodeEnvironmentInfo {
	var environments []VSCodeEnvironmentInfo

	// 按SortOrder排序的版本列表
	orderedEditions := []VSCodeEdition{
		EditionStandard,
		EditionInsiders,
		EditionExploration,
		EditionVSCodium,
		EditionVSCodiumInsiders,
		EditionOSS,
		EditionCursor,
		EditionCursorServer, // 新增
		EditionWindSurf,
		EditionTrae,
		EditionTraeCN,
		EditionCoder,
		EditionCodeServer,
		EditionRemoteSSH,
	}

	for _, edition := range orderedEditions {
		env := VSCODE_BUILTIN_ENVIRONMENTS[edition]
		extDir := GetVSCodeExtensionsDirectory(edition)
		exists := false

		// 检查扩展目录是否存在
		if _, err := os.Stat(extDir); err == nil {
			exists = true
		}

		environments = append(environments, VSCodeEnvironmentInfo{
			Edition:       edition,
			Environment:   env,
			ExtensionsDir: extDir,
			Exists:        exists,
		})
	}

	return environments
}

// DetectAllVSCodeEnvironments 检测所有可用的VSCode环境
func DetectAllVSCodeEnvironments() []VSCodeEdition {
	var detectedEditions []VSCodeEdition

	// 检查环境变量
	if os.Getenv("VSCODE_SERVER_PATH") != "" ||
		os.Getenv("CODE_SERVER_VERSION") != "" {
		detectedEditions = append(detectedEditions, EditionCodeServer)
	}

	if os.Getenv("VSCODE_AGENT_FOLDER") != "" ||
		os.Getenv("REMOTE_SSH_EXTENSION") != "" ||
		os.Getenv("VSCODE_SSH_HOST") != "" {
		detectedEditions = append(detectedEditions, EditionRemoteSSH)
	}

	// 检查扩展目录是否存在（与GetAllVSCodeEnvironments保持一致）
	home, _ := os.UserHomeDir()
	checkPaths := []struct {
		edition VSCodeEdition
		paths   []string
	}{
		{EditionRemoteSSH, []string{
			filepath.Join(home, ".vscode-server", "extensions"),
			filepath.Join(home, ".vscode-server", "data", "extensions"),
		}},
		{EditionCodeServer, []string{
			"/config/extensions",
			filepath.Join(home, ".local/share/code-server/extensions"),
			filepath.Join(home, ".config/code-server/extensions"),
		}},
		{EditionCursor, []string{
			filepath.Join(home, ".cursor", "extensions"),
		}},
		{EditionCursorServer, []string{
			"/config/.cursor-server/extensions",
		}},
		{EditionWindSurf, []string{
			filepath.Join(home, ".windsurf", "extensions"),
		}},
		{EditionTrae, []string{
			filepath.Join(home, ".trae", "extensions"),
		}},
		{EditionTraeCN, []string{
			filepath.Join(home, ".trae-cn", "extensions"),
		}},
		{EditionInsiders, []string{
			filepath.Join(home, ".vscode-insiders", "extensions"),
		}},
		{EditionVSCodium, []string{
			filepath.Join(home, ".vscode-oss", "extensions"),
		}},
		{EditionCoder, []string{
			filepath.Join(home, "vscode", "extensions"),
		}},
	}
	for _, item := range checkPaths {
		for _, extPath := range item.paths {
			if _, err := os.Stat(extPath); err == nil {
				if !contains(detectedEditions, item.edition) {
					detectedEditions = append(detectedEditions, item.edition)
				}
				break
			}
		}
	}

	// 如果没有检测到任何特定环境，添加标准版本
	if len(detectedEditions) == 0 {
		detectedEditions = append(detectedEditions, EditionStandard)
	}

	return detectedEditions
}

// GetDetectedEnvironmentsWithStatus 获取检测到的环境及其状态
func GetDetectedEnvironmentsWithStatus() []VSCodeEnvironmentInfo {
	var detectedEnvironments []VSCodeEnvironmentInfo

	// 获取所有环境
	allEnvironments := GetAllVSCodeEnvironments()

	// 获取检测到的环境
	detectedEditions := DetectAllVSCodeEnvironments()

	// 只返回检测到的环境
	for _, edition := range detectedEditions {
		for _, env := range allEnvironments {
			if env.Edition == edition {
				detectedEnvironments = append(detectedEnvironments, env)
				break
			}
		}
	}

	return detectedEnvironments
}

// VSCodeEnvironmentInfo VSCode环境信息
type VSCodeEnvironmentInfo struct {
	Edition       VSCodeEdition
	Environment   VSCodeEnvironment
	ExtensionsDir string
	Exists        bool
}
