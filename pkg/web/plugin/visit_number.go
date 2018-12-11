package plugin

import (
	"software_experiment/pkg/comm/database"
	"strconv"
	"strings"
	"time"
)

func ChangeVisitNumRedis(formName string, dataId uint) error {
	value, err := database.RedisClient.Exists(strings.Join([]string{formName, strconv.Itoa(int(dataId))}, "_")).Result()
	if err != nil {
		return err
	}
	if value == 0 {
		database.RedisClient.Incr(strings.Join([]string{formName, strconv.Itoa(int(dataId))}, "_"))
	} else {
		database.RedisClient.Set(strings.Join([]string{formName, strconv.Itoa(int(dataId))}, "_"), 0, 2*time.Minute)
	}
}
