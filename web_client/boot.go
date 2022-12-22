package web_client

import (
	"embed"

	"github.com/mimis-s/IM-Service/src/common/boot_config"
	"github.com/mimis-s/IM-Service/web_client/router"
	"golang.org/x/net/context"
)

//go:embed templates
var htmlEmbed embed.FS

//go:embed assets
var assetsEmbed embed.FS

func Boot(ctx context.Context, configOptions *boot_config.ConfigOptions) {
	router.Start(":8484", htmlEmbed, assetsEmbed)
}
