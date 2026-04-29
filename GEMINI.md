# Hub-Proxy-Go 项目指南

Hub-Proxy-Go 是一个高性能的 GitHub 和 Hugging Face 透明代理服务器，旨在加速 Git 克隆、Release 下载和大文件获取。

## 项目概览

- **核心功能**: 转发请求至 GitHub (包括 raw, gist, releases) 和 Hugging Face (模型, 数据集, Spaces)，支持自动将 GitHub `blob` 链接转换为 `raw` 链接。
- **技术栈**: Go 1.22+, 基于流式转发 (Streaming) 减少内存占用，利用 `sync.Pool` 复用字节缓冲区。
- **架构设计**: 
    - `cmd/hub-proxy-go/`: 包含主程序入口。
    - `internal/config/`: 处理基于环境变量的配置加载。
    - `internal/proxy/`: 实现反向代理逻辑、流式传输及内置网页 UI。
    - `internal/matcher/`: 负责 URL 路径的正则匹配与解析。

## 构建与运行

项目使用 `Makefile` 管理常用任务：

- **本地构建**: `make build` (生成 `hub-proxy-go` 二进制文件)
- **运行程序**: `make run`
- **运行测试**: `make test`
- **清理产物**: `make clean`
- **发布构建**: `make dist` (交叉编译 Linux x64 和 ARM64)

### 部署方式
- **Docker**: 支持从 Docker Hub 和 GHCR 拉取镜像。
- **Systemd**: 提供一键安装脚本 `scripts/install.sh` 自动配置服务。

## 开发规范

- **编程语言**: Go (Golang)，遵循 Google Go 编码规范。
- **文档语言**: **所有代码注释、README 及文档必须使用中文**。
- **Git 规范**: 
    - 远程仓库: `git@github.com:cololi/hub-proxy-go.git`
    - `.gitignore` 和 `.dockerignore` **不应**被提交至远程仓库（已通过 `git rm --cached` 处理）。
- **CI/CD**:
    - 使用 GitHub Actions 进行 Docker 镜像自动构建 (`docker.yml`)。
    - 标签推送 (`v*.*.*`) 会触发二进制文件自动发布 (`release.yml`)。

## 配置参考

| 环境变量 | 默认值 | 说明 |
| --- | --- | --- |
| `LISTEN` | `:8080` | 服务监听地址 |
| `SIZE_LIMIT` | `1072668082176` | 最大代理文件大小（字节），超出则重定向到原始地址 |
| `UPSTREAM_TIMEOUT` | `30s` | 请求上游的超时时间 |
| `SHUTDOWN_TIMEOUT` | `10s` | 优雅停机的超时时间 |
| `BUFFER_SIZE` | `32768` | 代理传输使用的缓冲区大小 (32KB) |
