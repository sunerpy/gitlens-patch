package config

// 激活代码常量
const (
	InsertCode = `e={user:{id:"88888888-8888-8888-8888-888888888888",name:"Neo",email:"x@x.com",status:"activated",createdDate:"2000-01-01T00:00:00.000Z"},licenses:{paidLicenses:{},effectiveLicenses:{"gitlens-pro":{organizationId:"Linux",latestStatus:"active",latestStartDate:"2024-01-01",latestEndDate:"2999-01-01",reactivationCount:99,nextOptInDate:"2999-01-01"}}},nextOptInDate:"2999-01-01"};`
)

// 版本支持配置
const (
	MinSupportedVersion = 15
	MaxSupportedVersion = 17
)

// 文件路径配置
const (
	GitLensJSPath      = "dist/gitlens.js"
	GitLensBrowserPath = "dist/browser/gitlens.js"
)

// 错误消息
const (
	ErrVersionNotSupported = "不支持的版本，仅支持 GitLens 15、16、17 版本"
	ErrGitLensNotFound     = "未找到 GitLens 扩展"
	ErrNoContentToReplace  = "未找到需要替换的内容"
	ErrNoPatternMatch      = "未找到匹配模式"
	ErrUnknownReplaceStyle = "未知的替换方式"
)

// 平台兼容的显示符号
const (
	// 状态符号
	StatusSuccess = "[OK]"
	StatusError   = "[ERROR]"
	StatusWarning = "[WARN]"
	StatusInfo    = "[INFO]"

	// 操作符号
	ActionSelect = ">>"
	ActionInput  = ">"
	ActionWait   = "..."

	// 分隔符
	SeparatorLine = "="
	SeparatorDash = "-"

	// 标记符号
	MarkerDetected = "[DETECTED]"
	MarkerRemote   = "[REMOTE]"
	MarkerCustom   = "[CUSTOM]"
)
