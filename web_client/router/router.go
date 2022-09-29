package router

import (
	"IM-Service/web_client/controller"
	"embed"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Start(port string, htmlEmbed embed.FS, assetsEmbed embed.FS) {
	engine := gin.Default()
	templ := template.Must(template.New("").ParseFS(htmlEmbed, "templates/*.html"))
	engine.SetHTMLTemplate(templ)
	// 加载静态资源(比如图片,文件等)
	engine.StaticFS("/assets", http.FS(assetsEmbed))

	engine.GET("/test", controller.Test)
	engine.GET("/", controller.Index)
	engine.GET("/chat", controller.Chat)

	engine.Run(port)
}
