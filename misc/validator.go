package misc

import (
	"time"

	"git.iguiyu.com/park/api/db"
	"github.com/labstack/echo"
)

func Validator(key string, c echo.Context) bool {
	userId, err := db.Redis().Get(db.REDIS_KEY_PREFIX + "-" + key).Result()
	if err == nil {
		db.Redis().Expire(db.REDIS_KEY_PREFIX+"-"+key, 7*24*time.Hour)
		c.Set("userId", userId)
		return true
	} else {
		return false
	}
}
