## Why

需要一个命令行工具来解析 Go 模块信息，方便在指定路径下运行 `go list -json -m all` 并将输出解析到内存中供后续处理使用。

## What Changes

- 新增一个命令行工具，支持传入路径参数
- 在指定路径下执行 `go list -json -m all` 命令
- 解析命令的 JSON 输出到内存结构中

## Capabilities

### New Capabilities
- `go-list-mod-parser`: 命令行工具，支持传入路径，执行 `go list -json -m all` 并解析输出到内存

### Modified Capabilities
（无）

## Impact

- 新增命令行工具代码
- 依赖 Go 环境（需要 `go` 命令可用）