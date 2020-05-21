输入以下命令进行快速开始

```bash
$ mkdir -p ./apis/v1
$ touch ./apis/v1/ping.get.json
$ echo "{\"body\": \"tong\"}" > ./apis/v1/ping.get.json
$ gmock ./apis
$ curl http://localhost:8080/v1/ping
tong
```
