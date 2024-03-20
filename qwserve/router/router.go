package router

import (
	"qwserve/models"
	"qwserve/pkg/app"
	webcomponents "qwserve/service/webComponents"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	RouterComponent(r)
	r.GET("/component/:name", func(c *gin.Context) {
		res := app.Gin{c}
		name := c.Param("name")
		s1 := ""
		if name == "myComponent.vue" {
			s1 = `
			<template>
			   <h2>{{a}}</h2>
			   <qlcustom content="content.vue"/>
			   <el-button>啊啊啊</el-button>
			</template>
			<script setup>
			import {ref, watch,shallowRef, defineProps,defineAsyncComponent} from 'vue';
			const a=ref("测试")
			</script>
			`
		} else if name == "content.vue" {
			s1 = `
			<template>
			   <el-button @click="aaaa">啊啊啊</el-button>
			   <el-input v-model="aa">啊啊啊</el-input>
			</template>
			<script setup>
			import {ref, watch,shallowRef, defineProps,getCurrentInstance} from 'vue';
			const {proxy} =getCurrentInstance();

			const a=ref("测试22")
			const aa=ref("测试22")

			const aaaa=()=>{
				proxy.$get("aaa/bbb")
			}
			</script>
			`
			//s1 = input.Inputjs
		}

		res.Response(200, "哈哈", s1)
		//c.String(200, "<template><h2>aaa</h2></template>")
	})

	//<template><h2>aaa</h2></template>
	return r
}

func RouterComponent(r *gin.Engine) {
	r.POST("components/add", func(c *gin.Context) {
		res := app.Gin{c}
		data := &models.WebComponents{}
		err := c.ShouldBind(data)
		if err != nil {
			res.Response(1001, err.Error(), "")
			return
		}
		id, err := webcomponents.Add(data)
		if err != nil {
			res.Response(1001, err.Error(), "")
			return
		}
		res.Response(200, "成功", id)
	})

}

//字段->转换成元素

type 元素 struct {
	name      string
	component string //组件名称 版本?
	data      string
	childrens []*元素
	attr      map[string]interface{} //包含tid id
}
