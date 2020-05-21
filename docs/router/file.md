同样支持返回一个文件流

在 `body` 处，以 File Protocol 格式的字符串，则会重定向到文件流当中

```bash
$ tree
./example/
├── home.get.json
├── home.html
$ cat ./example/home.get.json
{
  "body": "file://./home.html"
}
$ gmock ./
$ curl http://localhost:8080/home
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Hello</title>
</head>
<body>
  <p>Hello world! This is a HTML element</p>
</body>
</html>
```
