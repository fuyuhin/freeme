package controllers

import (
	"freeme/business/services"
	"freeme/models"
	"net/http"

	"github.com/8treenet/freedom"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
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

// func (c *AlbumController) Post() (*models.Album, error) {
// 	requestBody := &models.Album{}
// 	if err := c.Runtime.Ctx().ReadJSON(requestBody); err != nil {
// 		return nil, err
// 	}
// 	return c.Sev.Create(requestBody)
// }

func (c *AlbumController) Post() (*models.Album, error) {
	a := &models.Album{}
	if err := c.Runtime.Ctx().ReadJSON(a); err != nil {
		return nil, err
	}
	if err := c.Sev.Create(a); err != nil {
		return nil, err
	}
	return a, nil
}

// PutBy handles the PUT: /{id:int} route.
func (c *AlbumController) PutBy(id int) (mvc.Result, error) {
	a := models.Album{}
	if err := c.Runtime.Ctx().ReadJSON(&a); err != nil {
		return nil, err
	}
	if err := c.Sev.Update(id, a); err != nil {
		return nil, err
	}
	return mvc.Response{Code: http.StatusNoContent}, nil
}

// DeleteBy handles the PUT: /{id:int} route.
func (c *AlbumController) DeleteBy(id int) (mvc.Result, error) {
	err := c.Sev.Delete(id)
	if err != nil {
		return nil, err
	}
	return mvc.Response{Code: http.StatusNoContent}, nil
}
