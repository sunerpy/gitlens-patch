# GitLens Patch

一个用于激活 GitLens Pro 功能的工具。

## 功能特性

- 支持 GitLens 15、16、17 版本
- 自动检测 GitLens 安装路径
- 支持多种编辑器（VSCode、Cursor、VSCode Insiders、Windsurf）
- 自动备份原文件

## 目录结构

```
gitlens-patch/
├── cmd/
│   └── main.go              # 主程序入口
├── internal/
│   ├── app/
│   │   └── app.go           # 主应用逻辑
│   ├── config/
│   │   ├── constants.go      # 常量配置
│   │   └── paths.go          # 路径配置
│   ├── processor/
│   │   ├── processor.go      # 处理器接口
│   │   ├── factory.go        # 处理器工厂
│   │   ├── v15.go           # v15版本处理器
│   │   └── v16plus.go       # v16+版本处理器
│   └── utils/
│       ├── file.go           # 文件操作工具
│       ├── input.go          # 用户输入工具
│       └── path.go           # 路径处理工具
├── go.mod                    # Go模块文件
└── README.md                 # 项目说明
```

## 使用方法

### 编译

```bash
go build -o gitlens-patch cmd/main.go
```

### 运行

```bash
# 激活 GitLens Pro
./gitlens-patch

# 恢复原文件
./gitlens-patch restore

# 指定扩展目录
./gitlens-patch --ext-dir /path/to/extensions
```

### 环境变量

- `VSCODE_EXTENSIONS_DIR`: 指定 VSCode 扩展目录

## 版本支持

- ✅ GitLens 15.x
- ✅ GitLens 16.x  
- ✅ GitLens 17.x
- ❌ 其他版本（会提示不支持）

## 注意事项

1. 使用前请确保已安装 GitLens 扩展
2. 程序会自动备份原文件（.backup 后缀）
3. 修改后需要重启编辑器才能生效
4. 仅支持指定的版本，其他版本会提示不支持

