package server

import (
	"b3_ux_backend/internal/server/auth"
	"b3_ux_backend/internal/server/content"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	router.GET("/images/:id", ServeFile)

	// --- AUTH --- //
	{
		authGroup := router.Group("/auth")

		authGroup.POST("/signin", auth.SignIn)
		authGroup.POST("/login", auth.Login)
		authGroup.POST("/logout", auth.Logout)
		authGroup.POST("/me", auth.ValidateClientAccess(), auth.Me)
	}

	// --- Store --- //
	{
		storeGroup := router.Group("/store")
		storeGroup.Use(auth.ValidateClientAccess())

		storeGroup.POST("/app/findOne", content.FindOneStoreApp)
		storeGroup.POST("/app/findMany", content.FindManyStoreApps)

		storeGroup.POST("/genre/findOne", content.FindOneStoreGenre)
		storeGroup.POST("/genre/findMany", content.FindManyStoreGenres)
	}

	// --- WYWW --- //
	{
		wywwGroup := router.Group("/wyww")
		wywwGroup.Use(auth.ValidateClientAccess())

		wywwGroup.GET("/movie/newRecommendation", content.NewRecommendation)
		wywwGroup.GET("/movie/swapRecommendation", content.SwapRecommendation)
		wywwGroup.GET("/movie/watchlist", content.MovieInWatchlist)
		wywwGroup.PATCH("/movie/toggleWatched", content.ToggleWatchedMovie)
		wywwGroup.PATCH("/movie/toggleRecommended", content.ToggleRecommendedMovie)

		wywwGroup.POST("/movie/findOne", content.FindOneWywwMovie)
		wywwGroup.POST("/movie/findMany", content.FindManyWywwMovies)

		wywwGroup.POST("/genre/findOne", content.FindOneWywwGenre)
		wywwGroup.POST("/genre/findMany", content.FindManyWywwGenres)
	}

	// --- ADMIN --- //
	{
		apiGroup := router.Group("/admin")
		apiGroup.Use(auth.ValidateAdminAccess())
	}
}
