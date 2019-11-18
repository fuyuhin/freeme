package controllers

import (
	"freeme/business/services"
	"freeme/models"

	"github.com/8treenet/freedom"
	"github.com/kataras/iris"
)

func init() {
	freedom.Booting(func(initiator freedom.Initiator) {
		serFunc := func(ctx iris.Context) (m *services.AlbumService) {
			initiator.GetService(ctx, &m)
			return
		}
		initiator.BindControllerByParty(initiator.CreateParty("/albums"), &AlbumController{}, serFunc)
	})
}

// AlbumController .
type AlbumController struct {
	Sev     *services.AlbumService
	Runtime freedom.Runtime
}

// Get handles the GET: /{id:int} route.
func (c *AlbumController) GetBy(id int) (models.Album, error) {
	return c.Sev.GetAlbum(id)
}
