支持渲染 HTML 文件

如果想返回 HTML 模块，并且支持上下文/函数等，可以使用 Template Protocol (template:///path/to/your/file) 格式的字符串，它会编译模版，然后再返回文件流

例如：

```bash
$ tree
./example/
├── template.get.json
├── template.html
$ cat ./example/template.get.json
{
  "body": "template://./template.html"
}
$ gmock ./
$ curl http://localhost:8080/template
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Template</title>
</head>
<body>
  <p>Your request URL path: /template</p>
</body>
</html>
```