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
