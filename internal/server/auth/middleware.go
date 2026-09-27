package auth

import (
	"b3_ux_backend/internal/server/core"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/olekukonko/errors"
)

func sessionContext(ctx *gin.Context, checkAdmin bool) error {
	token, err := GetSessionTokenCookie(ctx)
	if err != nil {
		return err
	}

	// --- Resolve session --- //
	_, user, err := ResolveSession(ctx, token)
	if err != nil {
		return err
	}

	if checkAdmin {
		if !user.IsAdmin {
			return errors.New("not admin")
		}
	}

	// --- Create request actor --- //
	NewUserAuthDataFromDB(user).AddToContext(ctx)
	return nil
}

func ValidateClientAccess() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := sessionContext(ctx, false)
		if err != nil {
			ctx.Abort()
			core.HandleError(ctx, http.StatusUnauthorized, "Invalid access", err)
			return
		}

		ctx.Next()
	}
}

func ValidateAdminAccess() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := sessionContext(ctx, true)
		if err != nil {
			ctx.Abort()
			core.HandleError(ctx, http.StatusUnauthorized, "Invalid access", err)
			return
		}

		ctx.Next()
	}
}
