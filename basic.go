package basic

import (
	"github.com/gappy023/basic/config"
	"github.com/gappy023/basic/db"
	"github.com/gappy023/basic/redis"
)

func Init() {
	config.Init()
	db.Init()
	redis.Init()
}
