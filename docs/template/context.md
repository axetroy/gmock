模版的渲染都有上下文

上下文主要是当前的环境下一些变量

其中上下文包含三个部分

### 1. Request

[Request](https://golang.org/pkg/net/http/#Request) 是 Golang 的原生对象，里面包含了本次请求的相关信息

使用示例

```bash
$ tree ./example
./example/
├── template
│   └── context
│       └── request.get.json
$ cat ./example/template/context/request.get.json
{
  "body": "{{- .Request.URL.Path -}}"
}
$ gmock ./example
$ curl http://localhost:8080/template/context/request
/template/context/request
```

### 2. Query

Query 是请求的 URL 上带的参数, 例如 `?foo=bar`

类型为 `map[string]string` 或者 `map[string][]string`

### 3. Params

Params 是动态路由的参数, 例如 `/user/:id` 中的 `id` 则为 Params 的一部分

使用示例

```bash
$ tree ./example
./example/
├── template
│   └── context
│       └── params
│           └── [id].get.json
$ cat ./example/template/context/params/[id].get.json
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

### 4. Body

Body 是请求发送过来的数据体

这里提供几种字段

```go
var Body []byte
var BodyString string
var BodyMap map[string]interface{}
```

```json
{
  "body": {
    "string": "{{ .BodyString }}",
    "bytes": "{{ .Body }}",
    "map": "{{ .BodyMap }}",
    "map.name": "{{ index .BodyMap "name"}}"
  }
}
```

```bash
curl -X POST -d '{"name": "axetroy"}' http://localhost:8080

{"bytes":"[123 34 110 97 109 101 34 58 32 34 97 120 101 116 114 111 121 34 125]","map":"map[name:axetroy]","map.name":"axetroy","string":"{\"name\": \"axetroy\"}"}
```

### 5. Faker

Faker 提供了模拟假数据, 引用了 [faker 库](https://pkg.go.dev/github.com/bxcodec/faker/v3?tab=doc) 的实现

使用示例

```bash
$ tree ./example
./example/
├── template
│   └── context
│       ├── faker.get.json
$ cat ./example/template/context/faker.get.json
{
  "body": {
    "AmountWithCurrency": "{{ .Faker.AmountWithCurrency }}",
    "Currency": "{{ .Faker.Currency }}",
    "CreditCardNumber": "{{ .Faker.CreditCardNumber }}",
    "CreditCardType": "{{ .Faker.CreditCardType }}",
    "Century": "{{ .Faker.Century }}",
    "Date": "{{ .Faker.Date }}",
    "DayOfMonth": "{{ .Faker.DayOfMonth }}",
    "DayOfWeek": "{{ .Faker.DayOfWeek }}",
    "YearString": "{{ .Faker.YearString }}",
    "E164PhoneNumber": "{{ .Faker.E164PhoneNumber }}",
    "Name": "{{ .Faker.Name }}",
    "Username": "{{ .Faker.Username }}",
    "LastName": "{{ .Faker.LastName }}",
    "FirstName": "{{ .Faker.FirstName }}",
    "FirstNameFemale": "{{ .Faker.FirstNameFemale }}",
    "FirstNameMale": "{{ .Faker.FirstNameMale }}",
    "DomainName": "{{ .Faker.DomainName }}",
    "Email": "{{ .Faker.Email }}",
    "IPv4": "{{ .Faker.IPv4 }}",
    "IPv6": "{{ .Faker.IPv6 }}",
    "MacAddress": "{{ .Faker.MacAddress }}",
    "URL": "{{ .Faker.URL }}",
    "Latitude": {{ .Faker.Latitude }},
    "Longitude": {{ .Faker.Longitude }},
    "MonthName": "{{ .Faker.MonthName }}",
    "Paragraph": "{{ .Faker.Paragraph }}",
    "Password": "{{ .Faker.Password }}",
    "PhoneNumber": "{{ .Faker.PhoneNumber }}",
    "RandomUnixTime": {{ .Faker.RandomUnixTime }},
    "TimeString": "{{ .Faker.TimeString }}",
    "Timeperiod": "{{ .Faker.Timeperiod }}",
    "Timestamp": "{{ .Faker.Timestamp }}",
    "Timestamp": "{{ .Faker.Timestamp }}",
    "Timezone": "{{ .Faker.Timezone }}",
    "Timezone": "{{ .Faker.Timezone }}",
    "TitleFemale": "{{ .Faker.TitleFemale }}",
    "TitleMale": "{{ .Faker.TitleMale }}",
    "TollFreePhoneNumber": "{{ .Faker.TollFreePhoneNumber }}",
    "UUIDDigit": "{{ .Faker.UUIDDigit }}",
    "UUIDHyphenated": "{{ .Faker.UUIDHyphenated }}",
    "UnixTime": {{ .Faker.UnixTime }},
    "Word": "{{ .Faker.UnixTime }}",
    "Sentence": "{{ .Faker.Sentence }}"
  }
}
$ gmock ./example
$ curl http://localhost:8080/template/context/faker
{"AmountWithCurrency":"UAH 5.400000","Century":"XVI","CreditCardNumber":"342969764998659","CreditCardType":"American Express","Currency":"GMD","Date":"2003-02-18","DayOfMonth":"18","DayOfWeek":"Friday","DomainName":"ixObWPw.ru","E164PhoneNumber":"+710216493587","Email":"fPkoeLV@VNOoK.biz","FirstName":"Sarah","FirstNameFemale":"Oma","FirstNameMale":"Lane","IPv4":"204.189.118.229","IPv6":"f660:17f4:77e:3a37:8f93:7af4:47ca:97ca","LastName":"Hermann","Latitude":64.03468322753906,"Longitude":-58.44327163696289,"MacAddress":"64:f6:e9:1b:6e:6a","MonthName":"July","Name":"Prof. Hildegard Turner","Paragraph":"Accusantium consequatur aut sit perferendis voluptatem. Aut voluptatem perferendis accusantium sit consequatur. Accusantium consequatur voluptatem sit perferendis aut. Voluptatem consequatur perferendis accusantium sit aut. Accusantium consequatur sit aut perferendis voluptatem. Accusantium perferendis voluptatem sit consequatur aut.","Password":"gLmwcocnUpoScYHoJaLoeCdvOtDuMBuNsGbediIvgHNfrOTrJa","PhoneNumber":"891-624-3107","RandomUnixTime":413350644,"Sentence":"Consequatur sit voluptatem perferendis accusantium aut.","TimeString":"19:50:05","Timeperiod":"AM","Timestamp":"2001-07-10 23:48:45","Timezone":"Africa/Kampala","TitleFemale":"Miss","TitleMale":"King","TollFreePhoneNumber":"(777) 264-195873","URL":"http://www.LObidmm.net/","UUIDDigit":"72ad24422dc04e9787645f11a4f34833","UUIDHyphenated":"28f139ed-ae53-49af-b1b7-c2728e219521","UnixTime":474374152,"Username":"EQspIFS","Word":"709107104","YearString":"1990"}
```
