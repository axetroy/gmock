package function

import (
	"encoding/base64"
)

func Base64Encoding(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func Base64Decoding(str string) string {
	b, err := base64.StdEncoding.DecodeString(str)

	if err != nil {
		panic(err)
	}

	return string(b)
}
