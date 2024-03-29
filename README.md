![gmock](https://socialify.git.ci/axetroy/gmock/image?description=1&font=Inter&forks=1&issues=1&language=1&owner=1&pattern=Floating%20Cogs&pulls=1&stargazers=1&theme=Light)

[![Build Status](https://github.com/axetroy/gmock/workflows/ci/badge.svg)](https://github.com/axetroy/gmock/actions)
[![Coverage Status](https://coveralls.io/repos/github/axetroy/gmock/badge.svg?branch=master)](https://coveralls.io/github/axetroy/gmock?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/axetroy/gmock)](https://goreportcard.com/report/github.com/axetroy/gmock)
![Latest Version](https://img.shields.io/github/v/release/axetroy/gmock.svg)
![License](https://img.shields.io/github/license/axetroy/gmock.svg)
![Repo Size](https://img.shields.io/github/repo-size/axetroy/gmock.svg)

以文件为基准，最简单的 APIs 模拟工具

> 接口未至，开发先行

这里有一些 [example](/example) 已经部署到 [heroku](https://g-mock.herokuapp.com/template/context/faker)

特性:

- [x] 支持所有 HTTP 方法
- [x] 以最低的成本兼容 JSON 接口
- [x] 实时更改接口，无需重启服务
- [x] 自定义模版，满足复杂需求
- [x] 支持文件流，例如返回图片，下载等
- [x] 支持渲染 HTML/Markdown 文件

## 安装

1. Shell (Mac/Linux)

```bash
curl -fsSL https://github.com/release-lab/install/raw/v1/install.sh | bash -s -- -r=axetroy/gmock
```

2. PowerShell (Windows):

```powershell
$r="axetroy/gmock";iwr https://github.com/release-lab/install/raw/v1/install.ps1 -useb | iex
```

3. 从 [Github Release Page](https://github.com/axetroy/gmock/releases) 下载 (全平台支持)

> 下载可执行文件，并且把它加入到`$PATH` 环境变量中

4. 使用 [Golang](https://golang.org) 从源码中构建并安装 (全平台支持)

```bash
go install github.com/axetroy/gmock/cmd/whatchanged@latest
```

## 快速开始

首先先创建一个存放 APIs 文件的目录

```bash
$ mkdir -p ./apis/v1
```

然后在创建对应的 APIs 文件

```bash
$ touch ./apis/v1/ping.get.json
$ echo "{\"body\": \"tong\"}" > ./apis/v1/ping.get.json
```

运行命令

```bash
# gmock <接口文件目录>
$ gmock ./apis
$ curl http://localhost:8080/v1/ping
tong
```

## [文档](https://axetroy.github.io/gmock)

## 开源许可

The [MIT License](LICENSE)
