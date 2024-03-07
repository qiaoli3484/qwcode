package router

import (
	"qwserve/pkg/app"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/myComponent.vue", func(c *gin.Context) {
		res := app.Gin{c}
		res.Response(200, "哈哈", "<template><h2>aaa</h2></template>")
		//c.String(200, "<template><h2>aaa</h2></template>")
	})

	//<template><h2>aaa</h2></template>
	return r
}
