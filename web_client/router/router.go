package router

import (
	"embed"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mimis-s/IM-Service/web_client/controller"
)

func Start(port string, htmlEmbed embed.FS, assetsEmbed embed.FS) {
	engine := gin.Default()
	templ := template.Must(template.New("").ParseFS(htmlEmbed, "templates/*.html"))
	engine.SetHTMLTemplate(templ)
	// 加载静态资源(比如图片,文件等)
	engine.StaticFS("/assets", http.FS(assetsEmbed))

	engine.GET("/", controller.Index)
	engine.GET("/login", controller.Login)
	engine.GET("/register", controller.Register)
	engine.GET("/home", controller.Home)

	engine.Run(port)
}
