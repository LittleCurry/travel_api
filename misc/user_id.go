package misc

import "strconv"

func FormatUserIdToInt(userId string) int {
	id, _ := strconv.Atoi(userId)
	return id
}
