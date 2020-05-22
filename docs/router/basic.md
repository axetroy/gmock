最简单的路由例子

```bash
$ tree
./example/
├── hello.get.json
$ cat ./example/hello.get.json
{
  "body": "hello world!"
}
$ gmock ./
$ curl http://localhost:8080/hello
hello world!
```

如果设置跟路径的路由 `/`，则在目录中创建 `.{method}.json` 的文件

例如

```bash
$ tree
./example/
├── .get.json
$ cat ./example/.get.json
{
  "body": "root path"
}
$ gmock ./
$ curl http://localhost:8080
root path
```