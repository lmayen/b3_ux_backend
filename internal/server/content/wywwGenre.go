package content

import (
	"b3_ux_backend/internal/db"
	"b3_ux_backend/internal/entities"
	"b3_ux_backend/internal/server/core"

	"github.com/gin-gonic/gin"
)

func FindOneWywwGenre(ctx *gin.Context) {
	// Get params
	type queryParameters struct {
		Filters []entities.FilterItem `json:"filters"`
		Sorting *entities.SortOptions `json:"sorting"`
	}

	var params queryParameters
	if err := ctx.ShouldBindJSON(&params); err != nil {
		core.HandleError(ctx, 400, "Could not bind parameters", err)
		return
	}

	// Query
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

	entity, err := query.First(ctx)
	if err != nil {
		core.HandleError(ctx, 400, "Could not find actor", err)
		return
	}

	ctx.JSON(200, entity.ToData())
}

func FindManyWywwGenres(ctx *gin.Context) {
	// Get params
	type queryParameters struct {
		Filters    []entities.FilterItem       `json:"filters"`
		Sorting    *entities.SortOptions       `json:"sorting"`
		Pagination *entities.PaginationOptions `json:"pagination"`
	}

	var params queryParameters
	if err := ctx.ShouldBindJSON(&params); err != nil {
		core.HandleError(ctx, 400, "Could not bind parameters", err)
		return
	}

	// Query
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

	if params.Pagination != nil {
		query.ApplyPaginationOptions(params.Pagination)
	}

	queryResult, err := query.All(ctx)
	if err != nil {
		core.HandleError(ctx, 400, "Could not find actor", err)
		return
	}

	var result []*entities.WywwGenreData
	for _, entity := range queryResult {
		result = append(result, entity.ToData())
	}

	ctx.JSON(200, result)
}
