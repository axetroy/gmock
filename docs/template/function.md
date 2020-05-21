除了上下文，还提供了各种灵活的函数以便生成各种模版

除了 Golang 的[内置函数](https://golang.org/pkg/text/template/#hdr-Functions)之外

还添加了其他额外的函数功能

| 函数名            | 描述                | 参数          | 示例                                                      |
| ----------------- | ------------------- | ------------- | --------------------------------------------------------- |
| 运算              | --                  | --            | --                                                        |
| plusInt           | 累加函数            | `int`         | `plusInt 1 2` 等于 `3`                                    |
| minusInt          | 累减函数            | `int`         | `minusInt 2 1` 等于 `1`                                   |
| timesInt          | 累乘函数            | `int`         | `timesInt 2 3` 等于 `6`                                   |
| divInt            | 累除函数            | `int`         | `divInt 6 2` 等于 `3`                                     |
| 数组              | --                  | --            | --                                                        |
| makeSlice         | 创建数组            | `interface{}` | `makeSlice 6 2 3 5 1` 等于 `[]interface{}{6, 2, 3, 5, 1}` |
| makeSliceByLength | 创建长度为 n 的数组 | `int`         | `makeSliceByLength 5` 等于 `[]int{1, 2, 3, 4, 5}`         |
