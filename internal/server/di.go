package server

import (
	"github.com/eldelto/plant-guilds/internal/server/api"
	"github.com/eldelto/plant-guilds/internal/web"
)

type Container struct {
	AssetController    *web.Controller
	TemplateController *web.Controller
}

func Init() *Container {
	// Services

	// API
	assetControler := web.NewAssetController(api.AssetsFS)
	templateControler := web.NewTemplateController(api.TemplatesFS, &api.TemplateData{})

	return &Container{
		AssetController:    assetControler,
		TemplateController: templateControler,
	}
}

func (c *Container) Close() {
}
