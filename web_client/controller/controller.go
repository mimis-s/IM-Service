package controller

import "github.com/gin-gonic/gin"

func Index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

// 登录界面
func Login(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

// 注册界面
func Register(c *gin.Context) {
	c.HTML(200, "register.html", nil)
}

// 主界面
func Home(c *gin.Context) {
	c.HTML(200, "home.html", nil)
}
