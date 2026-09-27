
package api

import (
	"b3_ux_backend/internal/db"
	"b3_ux_backend/internal/entities"
	"b3_ux_backend/internal/server/core"

	"github.com/gin-gonic/gin"
	"github.com/olekukonko/errors"
)


type findManyParameters struct {
	Filters    []entities.FilterItem       `json:"filters"`
	Sorting    *entities.SortOptions       `json:"sorting,omitempty"`
	Pagination *entities.PaginationOptions `json:"pagination,omitempty"`
}


func FindManyColor(ctx *gin.Context) {
	var params findManyParameters
	if err := ctx.ShouldBindJSON(&params); err != nil {
		core.HandleError(ctx, 400, "could not bind parameters", errors.Wrapf(err, "FindManyColor ERR: could not bind parameters"))
		return
	}

	query := db.Client.Color.Query()

	if params.Filters != nil {
		err := query.ApplyFilterOptions(params.Filters)
		if err != nil {
		    core.HandleError(ctx, 400, "could not apply filters to the database query.", errors.Wrapf(err, "FindManyColor ERR: could not apply filters to the database query"))
			return
		}
	}

	if params.Pagination != nil {
		query.ApplyPaginationOptions(params.Pagination)
	}

	if params.Sorting != nil {
		query.ApplySortingOptions(params.Sorting)
	}

	queryResult, err := query.All(ctx)
	if err != nil {
		core.HandleError(ctx, 500, "could not query the Color table in the database", errors.Wrapf(err, "FindManyColor ERR: could not query the Color table in the databas"))
		return
	}

	var dataResult []*entities.ColorData
	for _, record := range queryResult {
		dataResult = append(dataResult, record.ToData())
	}

	ctx.JSON(200, dataResult)
}

func FindManyImage(ctx *gin.Context) {
	var params findManyParameters
	if err := ctx.ShouldBindJSON(&params); err != nil {
		core.HandleError(ctx, 400, "could not bind parameters", errors.Wrapf(err, "FindManyImage ERR: could not bind parameters"))
		return
	}

	query := db.Client.Image.Query()

	if params.Filters != nil {
		err := query.ApplyFilterOptions(params.Filters)
		if err != nil {
		    core.HandleError(ctx, 400, "could not apply filters to the database query.", errors.Wrapf(err, "FindManyImage ERR: could not apply filters to the database query"))
			return
		}
	}

	if params.Pagination != nil {
		query.ApplyPaginationOptions(params.Pagination)
	}

	if params.Sorting != nil {
		query.ApplySortingOptions(params.Sorting)
	}

	queryResult, err := query.All(ctx)
	if err != nil {
		core.HandleError(ctx, 500, "could not query the Image table in the database", errors.Wrapf(err, "FindManyImage ERR: could not query the Image table in the databas"))
		return
	}

	var dataResult []*entities.ImageData
	for _, record := range queryResult {
		dataResult = append(dataResult, record.ToData())
	}

	ctx.JSON(200, dataResult)
}

func FindManySession(ctx *gin.Context) {
	var params findManyParameters
	if err := ctx.ShouldBindJSON(&params); err != nil {
		core.HandleError(ctx, 400, "could not bind parameters", errors.Wrapf(err, "FindManySession ERR: could not bind parameters"))
		return
	}

	query := db.Client.Session.Query()

	if params.Filters != nil {
		err := query.ApplyFilterOptions(params.Filters)
		if err != nil {
		    core.HandleError(ctx, 400, "could not apply filters to the database query.", errors.Wrapf(err, "FindManySession ERR: could not apply filters to the database query"))
			return
		}
	}

	if params.Pagination != nil {
		query.ApplyPaginationOptions(params.Pagination)
	}

	if params.Sorting != nil {
		query.ApplySortingOptions(params.Sorting)
	}

	queryResult, err := query.All(ctx)
	if err != nil {
		core.HandleError(ctx, 500, "could not query the Session table in the database", errors.Wrapf(err, "FindManySession ERR: could not query the Session table in the databas"))
		return
	}

	var dataResult []*entities.SessionData
	for _, record := range queryResult {
		dataResult = append(dataResult, record.ToData())
	}

	ctx.JSON(200, dataResult)
}

func FindManyStoreApp(ctx *gin.Context) {
	var params findManyParameters
	if err := ctx.ShouldBindJSON(&params); err != nil {
		core.HandleError(ctx, 400, "could not bind parameters", errors.Wrapf(err, "FindManyStoreApp ERR: could not bind parameters"))
		return
	}

	query := db.Client.StoreApp.Query()

	if params.Filters != nil {
		err := query.ApplyFilterOptions(params.Filters)
		if err != nil {
		    core.HandleError(ctx, 400, "could not apply filters to the database query.", errors.Wrapf(err, "FindManyStoreApp ERR: could not apply filters to the database query"))
			return
		}
	}

	if params.Pagination != nil {
		query.ApplyPaginationOptions(params.Pagination)
	}

	if params.Sorting != nil {
		query.ApplySortingOptions(params.Sorting)
	}

	queryResult, err := query.All(ctx)
	if err != nil {
		core.HandleError(ctx, 500, "could not query the StoreApp table in the database", errors.Wrapf(err, "FindManyStoreApp ERR: could not query the StoreApp table in the databas"))
		return
	}

	var dataResult []*entities.StoreAppData
	for _, record := range queryResult {
		dataResult = append(dataResult, record.ToData())
	}

	ctx.JSON(200, dataResult)
}

func FindManyStoreGenre(ctx *gin.Context) {
	var params findManyParameters
	if err := ctx.ShouldBindJSON(&params); err != nil {
		core.HandleError(ctx, 400, "could not bind parameters", errors.Wrapf(err, "FindManyStoreGenre ERR: could not bind parameters"))
		return
	}

	query := db.Client.StoreGenre.Query()

	if params.Filters != nil {
		err := query.ApplyFilterOptions(params.Filters)
		if err != nil {
		    core.HandleError(ctx, 400, "could not apply filters to the database query.", errors.Wrapf(err, "FindManyStoreGenre ERR: could not apply filters to the database query"))
			return
		}
	}

	if params.Pagination != nil {
		query.ApplyPaginationOptions(params.Pagination)
	}

	if params.Sorting != nil {
		query.ApplySortingOptions(params.Sorting)
	}

	queryResult, err := query.All(ctx)
	if err != nil {
		core.HandleError(ctx, 500, "could not query the StoreGenre table in the database", errors.Wrapf(err, "FindManyStoreGenre ERR: could not query the StoreGenre table in the databas"))
		return
	}

	var dataResult []*entities.StoreGenreData
	for _, record := range queryResult {
		dataResult = append(dataResult, record.ToData())
	}

	ctx.JSON(200, dataResult)
}

func FindManyUser(ctx *gin.Context) {
	var params findManyParameters
	if err := ctx.ShouldBindJSON(&params); err != nil {
		core.HandleError(ctx, 400, "could not bind parameters", errors.Wrapf(err, "FindManyUser ERR: could not bind parameters"))
		return
	}

	query := db.Client.User.Query()

	if params.Filters != nil {
		err := query.ApplyFilterOptions(params.Filters)
		if err != nil {
		    core.HandleError(ctx, 400, "could not apply filters to the database query.", errors.Wrapf(err, "FindManyUser ERR: could not apply filters to the database query"))
			return
		}
	}

	if params.Pagination != nil {
		query.ApplyPaginationOptions(params.Pagination)
	}

	if params.Sorting != nil {
		query.ApplySortingOptions(params.Sorting)
	}

	queryResult, err := query.All(ctx)
	if err != nil {
		core.HandleError(ctx, 500, "could not query the User table in the database", errors.Wrapf(err, "FindManyUser ERR: could not query the User table in the databas"))
		return
	}

	var dataResult []*entities.UserData
	for _, record := range queryResult {
		dataResult = append(dataResult, record.ToData())
	}

	ctx.JSON(200, dataResult)
}

func FindManyWywwGenre(ctx *gin.Context) {
	var params findManyParameters
	if err := ctx.ShouldBindJSON(&params); err != nil {
		core.HandleError(ctx, 400, "could not bind parameters", errors.Wrapf(err, "FindManyWywwGenre ERR: could not bind parameters"))
		return
	}

	query := db.Client.WywwGenre.Query()

	if params.Filters != nil {
		err := query.ApplyFilterOptions(params.Filters)
		if err != nil {
		    core.HandleError(ctx, 400, "could not apply filters to the database query.", errors.Wrapf(err, "FindManyWywwGenre ERR: could not apply filters to the database query"))
			return
		}
	}

	if params.Pagination != nil {
		query.ApplyPaginationOptions(params.Pagination)
	}

	if params.Sorting != nil {
		query.ApplySortingOptions(params.Sorting)
	}

	queryResult, err := query.All(ctx)
	if err != nil {
		core.HandleError(ctx, 500, "could not query the WywwGenre table in the database", errors.Wrapf(err, "FindManyWywwGenre ERR: could not query the WywwGenre table in the databas"))
		return
	}

	var dataResult []*entities.WywwGenreData
	for _, record := range queryResult {
		dataResult = append(dataResult, record.ToData())
	}

	ctx.JSON(200, dataResult)
}

func FindManyWywwMovie(ctx *gin.Context) {
	var params findManyParameters
	if err := ctx.ShouldBindJSON(&params); err != nil {
		core.HandleError(ctx, 400, "could not bind parameters", errors.Wrapf(err, "FindManyWywwMovie ERR: could not bind parameters"))
		return
	}

	query := db.Client.WywwMovie.Query()

	if params.Filters != nil {
		err := query.ApplyFilterOptions(params.Filters)
		if err != nil {
		    core.HandleError(ctx, 400, "could not apply filters to the database query.", errors.Wrapf(err, "FindManyWywwMovie ERR: could not apply filters to the database query"))
			return
		}
	}

	if params.Pagination != nil {
		query.ApplyPaginationOptions(params.Pagination)
	}

	if params.Sorting != nil {
		query.ApplySortingOptions(params.Sorting)
	}

	queryResult, err := query.All(ctx)
	if err != nil {
		core.HandleError(ctx, 500, "could not query the WywwMovie table in the database", errors.Wrapf(err, "FindManyWywwMovie ERR: could not query the WywwMovie table in the databas"))
		return
	}

	var dataResult []*entities.WywwMovieData
	for _, record := range queryResult {
		dataResult = append(dataResult, record.ToData())
	}

	ctx.JSON(200, dataResult)
}

