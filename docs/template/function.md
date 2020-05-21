除了上下文，还提供了各种灵活的函数以便生成各种模版

除了 Golang 的[内置函数](https://golang.org/pkg/text/template/#hdr-Functions)之外

还添加了其他额外的函数功能

| 函数名            | 描述                    | 参数                   | 示例                                                       |
| ----------------- | ----------------------- | ---------------------- | ---------------------------------------------------------- |
| 运算              | --                      | --                     | --                                                         |
| plusInt           | 累加函数                | `plusInt ...int`       | `plusInt 1 2` 等于 `3`                                     |
| minusInt          | 累减函数                | `minusInt ...int`      | `minusInt 2 1` 等于 `1`                                    |
| timesInt          | 累乘函数                | `timesInt ...int`      | `timesInt 2 3` 等于 `6`                                    |
| divInt            | 累除函数                | `divInt ...int`        | `divInt 6 2` 等于 `3`                                      |
| 生成器            | --                      | --                     | --                                                         |
| randomStr         | 生成 n 个长度随机字符串 | `randomStr int string` | `randomStr 6` 等于 `生成长度为 6 的字符串`，第二个参数可选 |
| rangeInt          | 生成 n - m 范围内的整数 | `rangeInt int int`     | `rangeInt 1 10` 等于 `生成 1 - 10 范围内的整数`            |
| 数组              | --                      | --                     | --                                                         |
| makeSlice         | 创建数组                | `interface{}`          | `makeSlice 6 2 3 5 1` 等于 `[]interface{}{6, 2, 3, 5, 1}`  |
| makeSliceByLength | 创建长度为 n 的数组     | `int`                  | `makeSliceByLength 5` 等于 `[]int{1, 2, 3, 4, 5}`          |
