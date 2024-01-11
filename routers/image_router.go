package routers

import (
	"gvd_server/api"
	"gvd_server/middleware"
)

func (router RouterGroup) ImageRouter() {
	app := api.App.ImageApi
	router.POST("image", middleware.JwtAuth(), app.ImageUploadView)
}
