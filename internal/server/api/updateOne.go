
package api

import (
	"b3_ux_backend/internal/db"
	"b3_ux_backend/internal/entities"
	"b3_ux_backend/internal/server/core"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)


func UpdateOneColor(ctx *gin.Context) {
	type parameters struct {
		Id uuid.UUID `json:"id"`
		entities.ColorUpdateOptions
	}
	var params parameters
	if err := ctx.ShouldBindJSON(&params); err != nil {
		core.HandleError(ctx, 400, "Could not bind parameters", err)
		return
	}

	update := db.Client.Color.UpdateOneID(params.Id)
	err := update.ApplyUpdateOptions(&params.ColorUpdateOptions)
	if err != nil {
		core.HandleError(ctx, 400, "Could not apply update", err)
		return
	}

	result, err := update.Save(ctx)
	if err != nil {
		core.HandleError(ctx, 400, "Could not update Color", err)
		return
	}

	ctx.JSON(200, result.ToData())
}

func UpdateOneImage(ctx *gin.Context) {
	type parameters struct {
		Id uuid.UUID `json:"id"`
		entities.ImageUpdateOptions
	}
	var params parameters
	if err := ctx.ShouldBindJSON(&params); err != nil {
		core.HandleError(ctx, 400, "Could not bind parameters", err)
		return
	}

	update := db.Client.Image.UpdateOneID(params.Id)
	err := update.ApplyUpdateOptions(&params.ImageUpdateOptions)
	if err != nil {
		core.HandleError(ctx, 400, "Could not apply update", err)
		return
	}

	result, err := update.Save(ctx)
	if err != nil {
		core.HandleError(ctx, 400, "Could not update Image", err)
		return
	}

	ctx.JSON(200, result.ToData())
}

func UpdateOneSession(ctx *gin.Context) {
	type parameters struct {
		Id uuid.UUID `json:"id"`
		entities.SessionUpdateOptions
	}
	var params parameters
	if err := ctx.ShouldBindJSON(&params); err != nil {
		core.HandleError(ctx, 400, "Could not bind parameters", err)
		return
	}

	update := db.Client.Session.UpdateOneID(params.Id)
	err := update.ApplyUpdateOptions(&params.SessionUpdateOptions)
	if err != nil {
		core.HandleError(ctx, 400, "Could not apply update", err)
		return
	}

	result, err := update.Save(ctx)
	if err != nil {
		core.HandleError(ctx, 400, "Could not update Session", err)
		return
	}

	ctx.JSON(200, result.ToData())
}

func UpdateOneStoreApp(ctx *gin.Context) {
	type parameters struct {
		Id uuid.UUID `json:"id"`
		entities.StoreAppUpdateOptions
	}
	var params parameters
	if err := ctx.ShouldBindJSON(&params); err != nil {
		core.HandleError(ctx, 400, "Could not bind parameters", err)
		return
	}

	update := db.Client.StoreApp.UpdateOneID(params.Id)
	err := update.ApplyUpdateOptions(&params.StoreAppUpdateOptions)
	if err != nil {
		core.HandleError(ctx, 400, "Could not apply update", err)
		return
	}

	result, err := update.Save(ctx)
	if err != nil {
		core.HandleError(ctx, 400, "Could not update StoreApp", err)
		return
	}

	ctx.JSON(200, result.ToData())
}

func UpdateOneStoreGenre(ctx *gin.Context) {
	type parameters struct {
		Id uuid.UUID `json:"id"`
		entities.StoreGenreUpdateOptions
	}
	var params parameters
	if err := ctx.ShouldBindJSON(&params); err != nil {
		core.HandleError(ctx, 400, "Could not bind parameters", err)
		return
	}

	update := db.Client.StoreGenre.UpdateOneID(params.Id)
	err := update.ApplyUpdateOptions(&params.StoreGenreUpdateOptions)
	if err != nil {
		core.HandleError(ctx, 400, "Could not apply update", err)
		return
	}

	result, err := update.Save(ctx)
	if err != nil {
		core.HandleError(ctx, 400, "Could not update StoreGenre", err)
		return
	}

	ctx.JSON(200, result.ToData())
}

func UpdateOneUser(ctx *gin.Context) {
	type parameters struct {
		Id uuid.UUID `json:"id"`
		entities.UserUpdateOptions
	}
	var params parameters
	if err := ctx.ShouldBindJSON(&params); err != nil {
		core.HandleError(ctx, 400, "Could not bind parameters", err)
		return
	}

	update := db.Client.User.UpdateOneID(params.Id)
	err := update.ApplyUpdateOptions(&params.UserUpdateOptions)
	if err != nil {
		core.HandleError(ctx, 400, "Could not apply update", err)
		return
	}

	result, err := update.Save(ctx)
	if err != nil {
		core.HandleError(ctx, 400, "Could not update User", err)
		return
	}

	ctx.JSON(200, result.ToData())
}

func UpdateOneWywwGenre(ctx *gin.Context) {
	type parameters struct {
		Id uuid.UUID `json:"id"`
		entities.WywwGenreUpdateOptions
	}
	var params parameters
	if err := ctx.ShouldBindJSON(&params); err != nil {
		core.HandleError(ctx, 400, "Could not bind parameters", err)
		return
	}

	update := db.Client.WywwGenre.UpdateOneID(params.Id)
	err := update.ApplyUpdateOptions(&params.WywwGenreUpdateOptions)
	if err != nil {
		core.HandleError(ctx, 400, "Could not apply update", err)
		return
	}

	result, err := update.Save(ctx)
	if err != nil {
		core.HandleError(ctx, 400, "Could not update WywwGenre", err)
		return
	}

	ctx.JSON(200, result.ToData())
}

func UpdateOneWywwMovie(ctx *gin.Context) {
	type parameters struct {
		Id uuid.UUID `json:"id"`
		entities.WywwMovieUpdateOptions
	}
	var params parameters
	if err := ctx.ShouldBindJSON(&params); err != nil {
		core.HandleError(ctx, 400, "Could not bind parameters", err)
		return
	}

	update := db.Client.WywwMovie.UpdateOneID(params.Id)
	err := update.ApplyUpdateOptions(&params.WywwMovieUpdateOptions)
	if err != nil {
		core.HandleError(ctx, 400, "Could not apply update", err)
		return
	}

	result, err := update.Save(ctx)
	if err != nil {
		core.HandleError(ctx, 400, "Could not update WywwMovie", err)
		return
	}

	ctx.JSON(200, result.ToData())
}

