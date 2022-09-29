package web_client

import (
	"IM-Service/web_client/router"
	"embed"

	"golang.org/x/net/context"
)

//go:embed templates
var htmlEmbed embed.FS

//go:embed assets
var assetsEmbed embed.FS

func Boot(ctx context.Context) {
	router.Start(":8484", htmlEmbed, assetsEmbed)
}
