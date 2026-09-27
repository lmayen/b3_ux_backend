
package api

import (
	"b3_ux_backend/internal/db"
	"b3_ux_backend/internal/entities"
	"b3_ux_backend/internal/server/core"

	"github.com/gin-gonic/gin"
	"github.com/olekukonko/errors"
)

type findOneParameters struct {
	Filters []entities.FilterItem `json:"filters"`
	Sorting *entities.SortOptions `json:"sorting,omitempty"`
}


func FindOneColor(ctx *gin.Context) {
	var params findOneParameters
	if err := ctx.ShouldBindJSON(&params); err != nil {
		core.HandleError(ctx, 400, "Could not bind parameters", err)
		return
	}

	query := db.Client.Color.Query()

	if params.Filters != nil {
		err := query.ApplyFilterOptions(params.Filters)
		if err != nil {
			core.HandleError(ctx, 400, "Could not apply filters", err)
			return
		}
	}

	if params.Sorting != nil {
		query.ApplySortingOptions(params.Sorting)
	}

	queryResult, err := query.First(ctx)
	if err != nil {
		core.HandleError(ctx, 500, "could not query artists", errors.Wrapf(err, "FindOneArtists ERR: could not query artist"))
		return
	}

	ctx.JSON(200, queryResult.ToData())
}

func FindOneImage(ctx *gin.Context) {
	var params findOneParameters
	if err := ctx.ShouldBindJSON(&params); err != nil {
		core.HandleError(ctx, 400, "Could not bind parameters", err)
		return
	}

	query := db.Client.Image.Query()

	if params.Filters != nil {
		err := query.ApplyFilterOptions(params.Filters)
		if err != nil {
			core.HandleError(ctx, 400, "Could not apply filters", err)
			return
		}
	}

	if params.Sorting != nil {
		query.ApplySortingOptions(params.Sorting)
	}

	queryResult, err := query.First(ctx)
	if err != nil {
		core.HandleError(ctx, 500, "could not query artists", errors.Wrapf(err, "FindOneArtists ERR: could not query artist"))
		return
	}

	ctx.JSON(200, queryResult.ToData())
}

func FindOneSession(ctx *gin.Context) {
	var params findOneParameters
	if err := ctx.ShouldBindJSON(&params); err != nil {
		core.HandleError(ctx, 400, "Could not bind parameters", err)
		return
	}

	query := db.Client.Session.Query()

	if params.Filters != nil {
		err := query.ApplyFilterOptions(params.Filters)
		if err != nil {
			core.HandleError(ctx, 400, "Could not apply filters", err)
			return
		}
	}

	if params.Sorting != nil {
		query.ApplySortingOptions(params.Sorting)
	}

	queryResult, err := query.First(ctx)
	if err != nil {
		core.HandleError(ctx, 500, "could not query artists", errors.Wrapf(err, "FindOneArtists ERR: could not query artist"))
		return
	}

	ctx.JSON(200, queryResult.ToData())
}

func FindOneStoreApp(ctx *gin.Context) {
	var params findOneParameters
	if err := ctx.ShouldBindJSON(&params); err != nil {
		core.HandleError(ctx, 400, "Could not bind parameters", err)
		return
	}

	query := db.Client.StoreApp.Query()

	if params.Filters != nil {
		err := query.ApplyFilterOptions(params.Filters)
		if err != nil {
			core.HandleError(ctx, 400, "Could not apply filters", err)
			return
		}
	}

	if params.Sorting != nil {
		query.ApplySortingOptions(params.Sorting)
	}

	queryResult, err := query.First(ctx)
	if err != nil {
		core.HandleError(ctx, 500, "could not query artists", errors.Wrapf(err, "FindOneArtists ERR: could not query artist"))
		return
	}

	ctx.JSON(200, queryResult.ToData())
}

func FindOneStoreGenre(ctx *gin.Context) {
	var params findOneParameters
	if err := ctx.ShouldBindJSON(&params); err != nil {
		core.HandleError(ctx, 400, "Could not bind parameters", err)
		return
	}

	query := db.Client.StoreGenre.Query()

	if params.Filters != nil {
		err := query.ApplyFilterOptions(params.Filters)
		if err != nil {
			core.HandleError(ctx, 400, "Could not apply filters", err)
			return
		}
	}

	if params.Sorting != nil {
		query.ApplySortingOptions(params.Sorting)
	}

	queryResult, err := query.First(ctx)
	if err != nil {
		core.HandleError(ctx, 500, "could not query artists", errors.Wrapf(err, "FindOneArtists ERR: could not query artist"))
		return
	}

	ctx.JSON(200, queryResult.ToData())
}

func FindOneUser(ctx *gin.Context) {
	var params findOneParameters
	if err := ctx.ShouldBindJSON(&params); err != nil {
		core.HandleError(ctx, 400, "Could not bind parameters", err)
		return
	}

	query := db.Client.User.Query()

	if params.Filters != nil {
		err := query.ApplyFilterOptions(params.Filters)
		if err != nil {
			core.HandleError(ctx, 400, "Could not apply filters", err)
			return
		}
	}

	if params.Sorting != nil {
		query.ApplySortingOptions(params.Sorting)
	}

	queryResult, err := query.First(ctx)
	if err != nil {
		core.HandleError(ctx, 500, "could not query artists", errors.Wrapf(err, "FindOneArtists ERR: could not query artist"))
		return
	}

	ctx.JSON(200, queryResult.ToData())
}

func FindOneWywwGenre(ctx *gin.Context) {
	var params findOneParameters
	if err := ctx.ShouldBindJSON(&params); err != nil {
		core.HandleError(ctx, 400, "Could not bind parameters", err)
		return
	}

	query := db.Client.WywwGenre.Query()

	if params.Filters != nil {
		err := query.ApplyFilterOptions(params.Filters)
		if err != nil {
			core.HandleError(ctx, 400, "Could not apply filters", err)
			return
		}
	}

	if params.Sorting != nil {
		query.ApplySortingOptions(params.Sorting)
	}

	queryResult, err := query.First(ctx)
	if err != nil {
		core.HandleError(ctx, 500, "could not query artists", errors.Wrapf(err, "FindOneArtists ERR: could not query artist"))
		return
	}

	ctx.JSON(200, queryResult.ToData())
}

func FindOneWywwMovie(ctx *gin.Context) {
	var params findOneParameters
	if err := ctx.ShouldBindJSON(&params); err != nil {
		core.HandleError(ctx, 400, "Could not bind parameters", err)
		return
	}

	query := db.Client.WywwMovie.Query()

	if params.Filters != nil {
		err := query.ApplyFilterOptions(params.Filters)
		if err != nil {
			core.HandleError(ctx, 400, "Could not apply filters", err)
			return
		}
	}

	if params.Sorting != nil {
		query.ApplySortingOptions(params.Sorting)
	}

	queryResult, err := query.First(ctx)
	if err != nil {
		core.HandleError(ctx, 500, "could not query artists", errors.Wrapf(err, "FindOneArtists ERR: could not query artist"))
		return
	}

	ctx.JSON(200, queryResult.ToData())
}

