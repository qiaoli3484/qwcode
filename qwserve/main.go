package main

import (
	"qwserve/router"
)

func main() {
	r := router.InitRouter()

	r.Run(":18967")
}
