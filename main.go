/*
	名称：赏金任务查看系统
	公司：柏链项目学院
	作者：叶开
*/
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type RespMsg struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func login(c *gin.Context) {
	resp := &RespMsg{
		"0",
		"OK",
		nil,
	}
	defer c.JSON(200, resp)
	userinfo := make(map[string]string)
	c.Bind(&userinfo)
	fmt.Println(userinfo)
	if userinfo["user"] == "yekai" && userinfo["pass"] == "admin1" {
		return
	} else {
		resp.Code = "1"
		resp.Msg = "user or password err"
	}

}
func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.StaticFile("/", "static/index.html")
	r.Static("js", "static/js")
	r.Static("css", "static/css")

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.POST("login", login)
	r.Run(":8080")
}
