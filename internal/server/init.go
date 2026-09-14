package server

import (
	"b3_ux_backend/internal/config"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitServer(cfg *config.Config, shutdownCh chan<- struct{}) (*gin.Engine, *http.Server, error) {
	engine := gin.New()
	engine.Use(gin.Recovery())

	// Logging middleware (Zap-backed)
	engine.Use(GinZapMiddleware())

	// CORS (only if enabled)
	if cfg.CORS.Enabled {
		engine.Use(CORSMiddleware(cfg))
	}

	// Infrastructure routes
	routes.InitRoutes(engine)

	httpServer := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.HTTP.Host, cfg.HTTP.Port),
		Handler: engine,
	}

	return engine, httpServer, nil
}
