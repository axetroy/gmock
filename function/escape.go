package function

import "strings"

func Escape(str string) string {
	return strings.Replace(str, `"`, `\"`, -1)
}
