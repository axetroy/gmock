[![Build Status](https://github.com/axetroy/gmock/workflows/ci/badge.svg)](https://github.com/axetroy/gmock/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/axetroy/gmock)](https://goreportcard.com/report/github.com/axetroy/gmock)
![Latest Version](https://img.shields.io/github/v/release/axetroy/gmock.svg)
![License](https://img.shields.io/github/license/axetroy/gmock.svg)
![Repo Size](https://img.shields.io/github/repo-size/axetroy/gmock.svg)

Mock the APIs in the simplest way

### Usage

```bash
gmock ./
```

### Installation

If you are using Linux/macOS. you can install it with the following command:

```shell
# install latest version
curl -fsSL https://raw.githubusercontent.com/axetroy/gmock/master/install.sh | bash
# or install specified version
curl -fsSL https://raw.githubusercontent.com/axetroy/gmock/master/install.sh | bash -s v0.1.0
```

### Build from source code

Make sure you have `Golang@v1.14.2` installed.

```shell
$ git clone https://github.com/axetroy/gmock.git $GOPATH/src/github.com/axetroy/gmock
$ cd $GOPATH/src/github.com/axetroy/gmock
$ make build
```

### Test

```bash
$ make test
```

### License

The [MIT License](LICENSE)