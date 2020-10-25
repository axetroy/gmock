除了上下文，还提供了各种灵活的函数以便生成各种模版

除了 Golang 的[内置函数](https://golang.org/pkg/text/template/#hdr-Functions)之外

还添加了其他额外的函数功能

| 函数名            | 描述                      | 参数                   | 示例                                                       |
| ----------------- | ------------------------- | ---------------------- | ---------------------------------------------------------- |
| 运算              | --                        | --                     | --                                                         |
| Plus              | 累加函数                  | `Plus ...int`          | `Plus 1 2` 等于 `3`                                        |
| Minus             | 累减函数                  | `Minus ...int`         | `Minus 2 1` 等于 `1`                                       |
| Times             | 累乘函数                  | `Times ...int`         | `Times 2 3` 等于 `6`                                       |
| Div               | 累除函数                  | `Div ...int`           | `Div 6 2` 等于 `3`                                         |
| 随机数            | --                        | --                     | --                                                         |
| RandomStr         | 生成 n 个长度随机字符串   | `RandomStr int string` | `RandomStr 6` 等于 `生成长度为 6 的字符串`，第二个参数可选 |
| RangeInt          | 生成 n - m 范围内的整数   | `RangeInt int int`     | `RangeInt 1 10` 等于 `生成 1 - 10 范围内的整数`            |
| RangeFloat        | 生成 n - m 范围内的浮点数 | `RangeFloat int int`   | `RangeInt 0 1` 等于 `生成 0 - 1 范围内的浮点数`            |
| 数组              | --                        | --                     | --                                                         |
| MakeSlice         | 创建数组                  | `interface{}`          | `MakeSlice 6 2 3 5 1` 等于 `[]interface{}{6, 2, 3, 5, 1}`  |
| MakeSliceByLength | 创建长度为 n 的数组       | `int`                  | `MakeSliceByLength 5` 等于 `[]int{1, 2, 3, 4, 5}`          |
