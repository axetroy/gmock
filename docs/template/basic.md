这里我们使用一个简单的模版功能

```bash
$ tree
./example/
├── template
│   └── basic.get.json
$ cat ./example/template/basic.get.json
{
  "body": "{{ $x := 100 -}} {{- $x -}}"
}
$ ./gmock ./example
$ curl http://localhost:8080/example/template/basic
100
```