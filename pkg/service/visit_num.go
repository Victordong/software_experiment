package service

import (
	"context"
	"fmt"
	"software_experiment/pkg/comm/database"
	"software_experiment/pkg/comm/manager"
	"strconv"
	"strings"
)

func ChangeMysqlData() {
	ctx := context.Background()
	keys, err := database.RedisClient.Keys("*_*").Result()
	if err != nil {
		fmt.Printf(err.Error())
	}
	for _, key := range keys {
		value, err := database.RedisClient.Get(key).Result()
		if err != nil {
			fmt.Printf(err.Error())
		}
		args := strings.Split(key, "_")
		if len(args) == 2 {
			switch args[0] {
			case "exhibition":
				id, err := strconv.Atoi(args[1])
				if err != nil {
					fmt.Printf(err.Error())
				}
				valueInt, err := strconv.Atoi(value)
				if err != nil {
					fmt.Printf(err.Error())
				}
				err = manager.IncreaseExhibitionVistNum(ctx, uint(id), valueInt)
				if err != nil {
					fmt.Printf(err.Error())
				}
			case "information":
				id, _ := strconv.Atoi(args[1])
				if err != nil {
					fmt.Printf(err.Error())
				}
				valueInt, _ := strconv.Atoi(value)
				if err != nil {
					fmt.Printf(err.Error())
				}
				err = manager.IncreaseExhibitionVistNum(ctx, uint(id), valueInt)
				if err != nil {
					fmt.Printf(err.Error())
				}
			case "supply":
				id, err := strconv.Atoi(args[1])
				if err != nil {
					fmt.Printf(err.Error())
				}
				valueInt, err := strconv.Atoi(value)
				if err != nil {
					fmt.Printf(err.Error())
				}
				err = manager.IncreaseExhibitionVistNum(ctx, uint(id), valueInt)
				if err != nil {
					fmt.Printf(err.Error())
				}
			}
			_, err = database.RedisClient.Del(key).Result()
			if err != nil {
				fmt.Printf(err.Error())
			}
		}
	}
}
