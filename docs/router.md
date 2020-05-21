### 文件名规范

所有有效的路由名字都应遵循 `{name}.{method}.json` 规范

`{name}`: 路由名称
`{method}`: HTTP 方法名. 全部为小写. 例如 `get`/`post`/`put`

### 文件内容规范

文件内容应该包含以下字段

```go
type Schema struct {
	Status  *int                 `json:"status"`  // 返回的状态码
	Body    interface{}          `json:"body"`    // 请求体
	Headers *map[string][]string `json:"headers"` // 返回头
}
```

其中 `body` 定义了返回体，为必填项
