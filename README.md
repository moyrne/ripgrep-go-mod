# ripgrep-go-mod

生成 ripgrep 配置文件，自动排除 Go 模块和 GOROOT 目录。

## 功能

- 自动获取当前项目的 Go 模块列表
- 生成 `.ripgreprc` 配置文件，包含 `--glob` Go 模块目录
- 支持跨平台路径转换（Windows/Linux/macOS）

## 安装

```bash
go install github.com/moyrne/ripgrep-go-mod@latest
```

## 使用

```bash
# 在当前目录生成配置
ripgrep-go-mod

# 指定项目路径和输出文件
ripgrep-go-mod --project-path /path/to/project --output .ripgreprc
```

### 参数说明

- `--project-path`: 执行 `go list -json -m all` 的目录路径（默认：`.`）
- `--output`: 输出的 ripgrep 配置文件路径（默认：`.ripgreprc`）

## 工作原理

1. 执行 `go env -json` 获取 GOROOT 和 GOMODCACHE 信息
2. 执行 `go list -json -m all` 获取项目所有依赖模块
3. 为每个模块生成对应的 `--glob` 规则，排除模块目录
4. 将规则写入配置文件

## 示例输出

生成的 `.ripgreprc` 内容示例：

```
--glob=**/*/usr/local/go/src*/**
--glob=**/*github.com/example/module@v1.0.0*/**
--glob=**/*golang.org/x/sys@v0.1.0*/**
```