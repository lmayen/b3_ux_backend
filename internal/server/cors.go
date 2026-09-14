package server

import (
	"b3_ux_backend/internal/config"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware(cfg *config.Config) gin.HandlerFunc {
	allowedOrigins := make(map[string]struct{}, len(cfg.CORS.AllowedOrigins))
	for _, origin := range cfg.CORS.AllowedOrigins {
		allowedOrigins[origin] = struct{}{}
	}

	allowedMethods := strings.Join(cfg.CORS.AllowedMethods, ", ")
	allowedHeaders := strings.Join(cfg.CORS.AllowedHeaders, ", ")

	return func(ctx *gin.Context) {
		origin := ctx.GetHeader("Origin")

		// Requests from scripts normally have no Origin header and are not
		// subject to browser CORS rules.
		if origin != "" {
			if _, allowed := allowedOrigins[origin]; !allowed {
				ctx.AbortWithStatus(http.StatusForbidden)
				return
			}

			ctx.Header("Access-Control-Allow-Origin", origin)
			ctx.Header("Vary", "Origin")

			if cfg.CORS.AllowCredentials {
				ctx.Header("Access-Control-Allow-Credentials", "true")
			}
		}

		ctx.Header("Access-Control-Allow-Methods", allowedMethods)
		ctx.Header("Access-Control-Allow-Headers", allowedHeaders)

		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(http.StatusNoContent)
			return
		}

		ctx.Next()
	}
}
