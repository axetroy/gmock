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
