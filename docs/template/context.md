模版的渲染都有上下文

上下文主要是当前的环境下一些变量

其中上下文包含两个部分

1. Request

[Request](https://golang.org/pkg/net/http/#Request) 是 Golang 的原生对象，里面包含了本次请求的相关信息

使用示例

```bash
$ tree ./example
./example/
├── template
│   └── context
│       └── path.get.json
$ cat ./example/template/context/path.get.json
{
  "body": "{{- .Request.URL.Path -}}"
}
$ gmock ./example
$ curl http://localhost:8080/template/context/path
/template/context/path
```

2. Params

Params 是动态路由的参数, 例如 `/user/:id` 中的 `id` 则为 Params 的一部分

使用示例

```bash
$ tree ./example
./example/
├── template
│   └── context
│       └── user
│           └── [id].get.json
$ cat ./example/template/context/user/[id].get.json
{
  "body": {
    "uid": "{{- .Params.id -}}"
  }
}
$ gmock ./example
$ curl http://localhost:8080/template/context/user/123
{"uid": 123}
$ curl http://localhost:8080/template/context/user/321
{"uid": 321}
```
