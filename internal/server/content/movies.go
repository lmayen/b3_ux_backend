package content

import (
	"b3_ux_backend/internal/db"
	"b3_ux_backend/internal/entities"
	"b3_ux_backend/internal/entities/user"
	"b3_ux_backend/internal/entities/wywwmovie"
	"b3_ux_backend/internal/server/auth"
	"b3_ux_backend/internal/server/core"
	"math/rand"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/olekukonko/errors"
)

const recommendationPoolSize = 15

type scoredMovie struct {
	movie *entities.WywwMovie
	score int
}

func pickMovieCandidates(userEntity entities.User, movieEntities []*entities.WywwMovie, maxPoolSize int, genreId *uuid.UUID) ([]scoredMovie, error) {
	var candidates []scoredMovie

	recommendedMovies := userEntity.Edges.RecommendedMovies
	watchedMovies := userEntity.Edges.WatchedMovies

	// Build genre affinity
	genreWeights := make(map[uuid.UUID]int)
	for _, movie := range watchedMovies {
		for _, genre := range movie.Edges.Genres {
			genreWeights[genre.ID]++
		}
	}

	// Build exclusion sets
	excluded := make(map[uuid.UUID]struct{})
	for _, movie := range watchedMovies {
		excluded[movie.ID] = struct{}{}
	}

	for _, movie := range recommendedMovies {
		excluded[movie.ID] = struct{}{}
	}

	if genreId != nil {
		for _, movie := range movieEntities {
			matchesGenre := false
			for _, genre := range movie.Edges.Genres {
				if genre.ID == *genreId {
					matchesGenre = true
				}
			}
			if !matchesGenre {
				excluded[movie.ID] = struct{}{}
			}
		}
	}

	// --- Load and score candidates --- //
	for _, movie := range movieEntities {
		if _, exists := excluded[movie.ID]; exists {
			continue
		}

		score := 0
		for _, genre := range movie.Edges.Genres {
			score += genreWeights[genre.ID]
		}

		candidates = append(candidates, scoredMovie{
			movie: movie,
			score: score,
		})
	}

	if len(candidates) == 0 {
		return nil, errors.New("no recommendation available")
	}

	// --- Pick randomly among the best candidates --- //
	rand.Shuffle(len(candidates), func(i, j int) {
		candidates[i], candidates[j] = candidates[j], candidates[i]
	})

	sort.SliceStable(candidates, func(i, j int) bool {
		return candidates[i].score > candidates[j].score
	})

	poolSize := min(maxPoolSize, len(candidates))
	pool := candidates[:poolSize]

	return pool, nil
}

func NewRecommendation(ctx *gin.Context) {
	value, exists := ctx.Get(auth.UserAuthContextKey)
	if !exists {
		core.HandleError(ctx, http.StatusUnauthorized, "not authenticated", errors.New("not authenticated"))
		return
	}

	authData, ok := value.(*auth.UserAuthData)
	if !ok {
		core.HandleError(ctx, http.StatusInternalServerError, "invalid user in context", errors.New("invalid user in context"))
		return
	}

	// --------------------------------- //
	// --- Load main recommendation --- //
	// ------------------------------- //

	// --- Load user --- //
	tx, err := db.Client.Tx(ctx)
	if err != nil {
		core.HandleError(ctx, http.StatusInternalServerError, "failed to begin transaction", errors.New("failed to begin transaction"))
		return
	}
	defer tx.Rollback()

	userEntity, err := tx.User.Query().
		Where(user.IDEQ(authData.ID)).
		WithRecommendedMovies(func(query *entities.WywwMovieQuery) {
			query.WithGenres()
		}).
		WithWatchedMovies(func(query *entities.WywwMovieQuery) {
			query.WithGenres()
		}).
		Only(ctx)
	if err != nil {
		core.HandleError(ctx, http.StatusInternalServerError, "failed to query user", errors.New("failed to query user"))
		return
	}

	// --- Load movies --- //
	movies, err := tx.WywwMovie.Query().
		WithGenres().
		WithImages(func(query *entities.ImageQuery) {
			query.WithColors()
		}).
		All(ctx)
	if err != nil {
		core.HandleError(ctx, http.StatusInternalServerError, "failed to load movies", errors.New("failed to load movies"))
		return
	}

	mainCandidates, err := pickMovieCandidates(*userEntity, movies, recommendationPoolSize, nil)
	if err != nil {
		core.HandleError(ctx, http.StatusInternalServerError, "failed to load candidates", errors.Wrapf(err, "failed to load candidates"))
		return
	}
	mainSelected := mainCandidates[rand.Intn(len(mainCandidates))].movie

	// ---------------------------------- //
	// --- Load candidates per genre --- //
	// -------------------------------- //

	selectedGenres := make(map[string][]*entities.WywwMovieData)
	for _, genre := range mainSelected.Edges.Genres {
		candidates, err := pickMovieCandidates(*userEntity, movies, 5, &genre.ID)
		if err != nil {
			core.HandleError(ctx, http.StatusInternalServerError, "failed to load candidates", errors.Wrapf(err, "failed to load candidates"))
			return
		}
		var selection []*entities.WywwMovieData
		for _, candidate := range candidates {
			selection = append(selection, candidate.movie.ToData())
		}
		if selection != nil && len(selection) > 0 {
			selectedGenres[genre.Name] = selection
		}
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"main":   mainSelected.ToData(),
			"genres": selectedGenres,
		},
	)
}

func SwapRecommendation(ctx *gin.Context) {
	// Auth
	value, exists := ctx.Get(auth.UserAuthContextKey)
	if !exists {
		core.HandleError(ctx, http.StatusUnauthorized, "not authenticated", errors.New("not authenticated"))
		return
	}

	authData, ok := value.(*auth.UserAuthData)
	if !ok {
		core.HandleError(ctx, http.StatusInternalServerError, "invalid user in context", errors.New("invalid user in context"))
		return
	}

	// parameters
	type queryParameters struct {
		RawId string    `form:"id" binding:"required" json:"id"`
		Id    uuid.UUID `json:"-"`
	}

	var params queryParameters
	if err := ctx.ShouldBindQuery(&params); err != nil {
		core.HandleError(ctx, http.StatusBadRequest, "Could not parse parameters", err)
		return
	}

	id, err := uuid.Parse(params.RawId)
	if err != nil {
		core.HandleError(ctx, http.StatusBadRequest, "Could not parse id", err)
		return
	}
	params.Id = id

	// --------------------------------- //
	// --- Load main recommendation --- //
	// ------------------------------- //

	// --- Load user --- //
	tx, err := db.Client.Tx(ctx)
	if err != nil {
		core.HandleError(ctx, http.StatusInternalServerError, "failed to begin transaction", errors.New("failed to begin transaction"))
		return
	}
	defer tx.Rollback()

	userEntity, err := tx.User.Query().
		Where(user.IDEQ(authData.ID)).
		WithRecommendedMovies(func(query *entities.WywwMovieQuery) {
			query.WithGenres()
		}).
		WithWatchedMovies(func(query *entities.WywwMovieQuery) {
			query.WithGenres()
		}).
		Only(ctx)
	if err != nil {
		core.HandleError(ctx, http.StatusInternalServerError, "failed to query user", errors.New("failed to query user"))
		return
	}

	// --- Load movies --- //
	movies, err := tx.WywwMovie.Query().
		WithGenres().
		WithImages(func(query *entities.ImageQuery) {
			query.WithColors()
		}).
		All(ctx)
	if err != nil {
		core.HandleError(ctx, http.StatusInternalServerError, "failed to load movies", errors.New("failed to load movies"))
		return
	}

	mainSelected, err := tx.WywwMovie.Query().
		Where(wywwmovie.IDEQ(params.Id)).
		WithImages(func(query *entities.ImageQuery) {
			query.WithColors()
		}).
		WithGenres().Only(ctx)
	if err != nil {
		core.HandleError(ctx, http.StatusInternalServerError, "failed to load movies", errors.New("failed to load movies"))
		return
	}

	// ---------------------------------- //
	// --- Load candidates per genre --- //
	// -------------------------------- //

	selectedGenres := make(map[string][]*entities.WywwMovieData)
	for _, genre := range mainSelected.Edges.Genres {
		candidates, err := pickMovieCandidates(*userEntity, movies, 5, &genre.ID)
		if err != nil {
			core.HandleError(ctx, http.StatusInternalServerError, "failed to load candidates", errors.Wrapf(err, "failed to load candidates"))
			return
		}
		var selection []*entities.WywwMovieData
		for _, candidate := range candidates {
			selection = append(selection, candidate.movie.ToData())
		}
		if selection != nil && len(selection) > 3 {
			selectedGenres[genre.Name] = selection
		}
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"main":   mainSelected.ToData(),
			"genres": selectedGenres,
		},
	)
}

func MovieInWatchlist(ctx *gin.Context) {
	// Auth
	value, exists := ctx.Get(auth.UserAuthContextKey)
	if !exists {
		core.HandleError(ctx, http.StatusUnauthorized, "not authenticated", errors.New("not authenticated"))
		return
	}

	authData, ok := value.(*auth.UserAuthData)
	if !ok {
		core.HandleError(ctx, http.StatusInternalServerError, "invalid user in context", errors.New("invalid user in context"))
		return
	}

	// parameters
	type queryParameters struct {
		RawId string    `form:"id" binding:"required" json:"id"`
		Id    uuid.UUID `json:"-"`
	}

	var params queryParameters
	if err := ctx.ShouldBindQuery(&params); err != nil {
		core.HandleError(ctx, http.StatusBadRequest, "Could not parse parameters", err)
		return
	}

	id, err := uuid.Parse(params.RawId)
	if err != nil {
		core.HandleError(ctx, http.StatusBadRequest, "Could not parse id", err)
		return
	}
	params.Id = id

	// ----------------------- //
	// --- Query database --- //
	// --------------------- //
	tx, err := db.Client.Tx(ctx)
	if err != nil {
		core.HandleError(ctx, http.StatusInternalServerError, "failed to begin transaction", errors.New("failed to begin transaction"))
		return
	}
	defer tx.Rollback()

	watched, err := tx.User.Query().Where(
		user.IDEQ(authData.ID),
		user.HasWatchedMoviesWith(wywwmovie.IDEQ(params.Id)),
	).Exist(ctx)
	if err != nil {
		core.HandleError(ctx, http.StatusInternalServerError, "failed to query user watched movies", errors.New("failed to query user watched movies"))
		return
	}

	recommended, err := tx.User.Query().Where(
		user.IDEQ(authData.ID),
		user.HasRecommendedMoviesWith(wywwmovie.IDEQ(params.Id)),
	).Exist(ctx)
	if err != nil {
		core.HandleError(ctx, http.StatusInternalServerError, "failed to query user watched movies", errors.New("failed to query user watched movies"))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"watched": watched, "recommended": recommended})
}

func ToggleWatchedMovie(ctx *gin.Context) {
	// Auth
	value, exists := ctx.Get(auth.UserAuthContextKey)
	if !exists {
		core.HandleError(ctx, http.StatusUnauthorized, "not authenticated", errors.New("not authenticated"))
		return
	}

	authData, ok := value.(*auth.UserAuthData)
	if !ok {
		core.HandleError(ctx, http.StatusInternalServerError, "invalid user in context", errors.New("invalid user in context"))
		return
	}

	// parameters
	type queryParameters struct {
		Id uuid.UUID `json:"id"`
	}

	var params queryParameters
	if err := ctx.ShouldBindJSON(&params); err != nil {
		core.HandleError(ctx, http.StatusBadRequest, "Could not parse parameters", err)
		return
	}

	// ----------------------- //
	// --- Query database --- //
	// --------------------- //
	tx, err := db.Client.Tx(ctx)
	if err != nil {
		core.HandleError(ctx, http.StatusInternalServerError, "failed to begin transaction", errors.New("failed to begin transaction"))
		return
	}
	defer tx.Rollback()

	userEntity, err := tx.User.Query().
		WithWatchedMovies().
		Where(user.IDEQ(authData.ID)).
		Only(ctx)
	if err != nil {
		core.HandleError(ctx, http.StatusInternalServerError, "failed to query user watched movies", errors.New("failed to query user watched movies"))
		return
	}

	movieRemoved := false
	for _, movie := range userEntity.Edges.WatchedMovies {
		if movie.ID == params.Id {
			err = tx.User.UpdateOneID(authData.ID).RemoveWatchedMovieIDs(params.Id).Exec(ctx)
			if err != nil {
				core.HandleError(ctx, http.StatusInternalServerError, "failed to remove watched movie", errors.New("failed to remove watched movie"))
				return
			}
			movieRemoved = true
			break
		}
	}

	if !movieRemoved {
		err = tx.User.UpdateOneID(authData.ID).AddWatchedMovieIDs(params.Id).Exec(ctx)
		if err != nil {
			core.HandleError(ctx, http.StatusInternalServerError, "failed to add watched movie", errors.New("failed to add watched movie"))
		}
	}

	err = tx.Commit()
	if err != nil {
		core.HandleError(ctx, http.StatusInternalServerError, "failed to commit transaction", errors.New("failed to commit transaction"))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"watched": !movieRemoved})
}

func ToggleRecommendedMovie(ctx *gin.Context) {
	// Auth
	value, exists := ctx.Get(auth.UserAuthContextKey)
	if !exists {
		core.HandleError(ctx, http.StatusUnauthorized, "not authenticated", errors.New("not authenticated"))
		return
	}

	authData, ok := value.(*auth.UserAuthData)
	if !ok {
		core.HandleError(ctx, http.StatusInternalServerError, "invalid user in context", errors.New("invalid user in context"))
		return
	}

	// parameters
	type queryParameters struct {
		Id uuid.UUID `json:"id"`
	}

	var params queryParameters
	if err := ctx.ShouldBindJSON(&params); err != nil {
		core.HandleError(ctx, http.StatusBadRequest, "Could not parse parameters", err)
		return
	}

	// ----------------------- //
	// --- Query database --- //
	// --------------------- //
	tx, err := db.Client.Tx(ctx)
	if err != nil {
		core.HandleError(ctx, http.StatusInternalServerError, "failed to begin transaction", errors.New("failed to begin transaction"))
		return
	}
	defer tx.Rollback()

	userEntity, err := tx.User.Query().
		WithRecommendedMovies().
		Where(user.IDEQ(authData.ID)).
		Only(ctx)
	if err != nil {
		core.HandleError(ctx, http.StatusInternalServerError, "failed to query user recommended movies", errors.New("failed to query user recommended movies"))
		return
	}

	movieRemoved := false
	for _, movie := range userEntity.Edges.RecommendedMovies {
		if movie.ID == params.Id {
			err = tx.User.UpdateOneID(authData.ID).RemoveRecommendedMovieIDs(params.Id).Exec(ctx)
			if err != nil {
				core.HandleError(ctx, http.StatusInternalServerError, "failed to remove recommended movie", errors.New("failed to remove recommended movie"))
				return
			}
			movieRemoved = true
			break
		}
	}

	if !movieRemoved {
		err = tx.User.UpdateOneID(authData.ID).AddRecommendedMovieIDs(params.Id).Exec(ctx)
		if err != nil {
			core.HandleError(ctx, http.StatusInternalServerError, "failed to add recommended movie", errors.New("failed to add recommended movie"))
		}
	}

	err = tx.Commit()
	if err != nil {
		core.HandleError(ctx, http.StatusInternalServerError, "failed to commit transaction", errors.New("failed to commit transaction"))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"recommended": !movieRemoved})
}
