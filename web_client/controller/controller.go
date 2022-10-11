package controller

import "github.com/gin-gonic/gin"

// 专门用于测试
func Test(c *gin.Context) {
	c.HTML(200, "test.html", nil)
}

// 主界面(可能成为登录界面)
func Index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

// 聊天界面
func Chat(c *gin.Context) {
	c.HTML(200, "chat.html", nil)
}

// 登录界面
func Login(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}
