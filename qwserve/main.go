package main

import (
	"qwserve/init/initDB"
	"qwserve/router"
)

func main() {
	initDB.InitMysql() //数据库连接

	r := router.InitRouter()

	r.Run(":18967")
}
