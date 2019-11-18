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

// GetBy handles the GET: /{id:int} route.
func (c *AlbumController) GetBy(id int) (models.Album, error) {
	return c.Sev.GetAlbum(id)
}

// Get .
func (c *AlbumController) Get() ([]models.Album, error) {
	ctx := c.Runtime.Ctx()
	page := ctx.URLParamIntDefault("page", 1)
	perPage := ctx.URLParamIntDefault("per_page", 10)
	filter := &models.Album{
		ArtistID: ctx.URLParamIntDefault("artist_id", 0),
		Title:    ctx.URLParamDefault("title", ""),
	}
	return c.Sev.GetAlbums(page, perPage, filter)
}
