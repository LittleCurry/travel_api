package misc

import "strings"

func GetAccessToken(s string) string {
	return strings.Split(s, " ")[1]
}
