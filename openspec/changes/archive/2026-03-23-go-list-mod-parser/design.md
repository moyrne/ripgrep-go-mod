## Context

这是一个 Go 语言编写的命令行工具，用于在指定路径下执行 `go list -json -m all` 并将输出解析到内存中。项目当前位于 `/home/kali/develop/src/github.com/moyrne/ripgrep-go-mod`。

## Goals / Non-Goals

**Goals:**
- 提供简单的命令行接口，支持传入路径参数
- 执行 `go list -json -m all` 命令
- 将 JSON 输出解析为 Go 结构体

**Non-Goals:**
- 修改 Go 模块信息
- 提供图形用户界面
- 支持除 `go list -json -m all` 之外的其他 go 命令

## Decisions

### 语言选择: Go
- **决定**: 使用 Go 语言实现
- **理由**: 工具本身用于处理 Go 模块，使用 Go 可以无缝集成；Go 的标准库提供了强大的命令执行和 JSON 解析能力
- **替代方案**: Python, Rust

### 命令行参数解析: flag 包
- **决定**: 使用 Go 标准库的 flag 包
- **理由**: 简单够用，无需引入第三方依赖
- **替代方案**: cobra, pflag

### JSON 解析: encoding/json
- **决定**: 使用 Go 标准库的 encoding/json 包
- **理由**: 标准库提供了完善的 JSON 解析功能
- **替代方案**: easyjson, jsonparser

### 目录切换: os.Chdir
- **决定**: 在执行 go 命令前使用 os.Chdir 切换到指定目录
- **理由**: 简单直接，符合 `go list` 命令的工作方式
- **替代方案**: 使用 exec.Command 的 Dir 字段

## Risks / Trade-offs

- **Risk**: 不同 Go 版本的 `go list -json -m all` 输出格式可能有差异 → Mitigation: 解析时使用灵活的结构体，保持向后兼容
- **Risk**: 用户没有安装 Go 环境 → Mitigation: 提供清晰的错误提示