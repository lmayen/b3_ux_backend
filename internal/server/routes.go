package server

import (
	"b3_ux_backend/internal/server/content"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	// --- Store --- //
	{
		storeGroup := router.Group("/store")

		storeGroup.GET("/apps", content.GetManyStoreApp)
		//storeGroup.POST("/logout", auth.Logout)
	}

	// --- WYWW --- //
	{
		wywwGroup := router.Group("/wyww")

		wywwGroup.GET("/genre", content.GetOneWywwGenre)
		wywwGroup.GET("/genres", content.GetManyWywwGenre)
	}
}
