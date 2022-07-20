package main

import (
	"blog-go/dao/mysql"
	"blog-go/dao/redisd"
	"blog-go/routers"
)

func main() {
	//连接mysql数据库
	err := mysql.InitMysql()
	if err != nil {
		panic(err)
	}
	err = redisd.SetupRedisDb()
	if err != nil {
		panic(err)
	}
	r := routers.SetupRouter()

	//启动http服务
	err = r.Run(":9090")
	if err != nil {
		return
	}
}
