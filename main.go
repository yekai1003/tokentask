package main

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/yekai1003/tokentask/routes"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	store := sessions.NewCookieStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	// r.StaticFile("/", "static/index.html")
	// r.StaticFile("/tasklist.html", "static/tasklist.html")
	// r.Static("js", "static/js")
	// r.Static("css", "static/css")
	// r.Static("bootstrap", "static/bootstrap")

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.POST("login", routes.Login)
	r.POST("register", routes.Register)
	r.POST("issue", routes.Issue)
	r.POST("update", routes.Modify)
	r.GET("tasklist", routes.TaskList)
	r.Run(":8080")
}
