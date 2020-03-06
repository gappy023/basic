package basic

import (
	"github.com/microservice/part1/basic/config"
	"github.com/microservice/part1/basic/db"
	"github.com/microservice/part1/basic/redis"
)

func Init() {
	config.Init()
	db.Init()
	redis.Init()
}
