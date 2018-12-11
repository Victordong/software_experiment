package main

import (
	"software_experiment/pkg/service"
	"time"
)

func main() {
	for {
		service.ChangeMysqlData()
		time.Sleep(time.Minute)
	}

}
