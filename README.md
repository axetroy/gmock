[![Build Status](https://github.com/axetroy/gmock/workflows/ci/badge.svg)](https://github.com/axetroy/gmock/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/axetroy/gmock)](https://goreportcard.com/report/github.com/axetroy/gmock)
![Latest Version](https://img.shields.io/github/v/release/axetroy/gmock.svg)
![License](https://img.shields.io/github/license/axetroy/gmock.svg)
![Repo Size](https://img.shields.io/github/repo-size/axetroy/gmock.svg)

以文件为基准，最简单的 APIs 模拟工具

> 接口未至，开发先行

特性:

- [x] 支持所有 HTTP 方法
- [x] 以最低的成本兼容 JSON 接口
- [x] 自定义模版

## 安装方式

如果你使用 `Linux` 或 `macOS`， 你可以输入一下命令安装:

```shell
# 安装最新版本
curl -fsSL https://raw.githubusercontent.com/axetroy/gmock/master/install.sh | bash
# 安装指定的版本
curl -fsSL https://raw.githubusercontent.com/axetroy/gmock/master/install.sh | bash -s v0.1.0
```

## 快速开始

首先先创建一个用户存放 APIs 文件的目录

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

## 文档

所有的 APIs 文件命名规则遵循 `{name}.{method}.json`. `name` 为文件名，`method` 为 http 的小写方法

并且 JSON 的文件内容遵循一下格式, 其中 `body` 为必填

```go
type Schema struct {
	Status  *int                 `json:"status"`  // 返回的状态码
	Body    interface{}          `json:"body"`    // 请求体
	Headers *map[string][]string `json:"headers"` // 返回头
}
```

文件名 `{name}` 也可以是动态参数，以 `[name]` 为格式

例如文件 `$root/v1/user/[id].get.json` 则匹配路由 `GET /v1/user/{id}`

如果你觉得简单的模版无法满足复杂的 mock 需求，那么还可以使用模版功能

## 从源码中构建

确保你已安装 `Golang@v1.14.2` 或者更高版本.

```shell
$ git clone https://github.com/axetroy/gmock.git $GOPATH/src/github.com/axetroy/gmock
$ cd $GOPATH/src/github.com/axetroy/gmock
$ make build
```

## 测试

```bash
$ make test
```

## 开源许可

The [MIT License](LICENSE)
