package auth

import (
	"b3_ux_backend/internal/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/olekukonko/errors"
)

func GetSessionTokenCookie(ctx *gin.Context) (string, error) {
	token, err := ctx.Cookie(SessionCookieName)
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			return "", errors.Wrapf(err, "Session cookie not found")
		}
		return "", errors.Wrapf(err, "Could not read session")
	}

	return token, nil
}

func SetSessionTokenCookie(ctx *gin.Context, token string) {
	ctx.SetSameSite(http.SameSiteLaxMode)

	ctx.SetCookie(
		SessionCookieName,
		token,
		sessionCookieMaxAge,
		"/",
		"",
		config.CONFIG.HTTP.SecureCookies,
		true,
	)
}

func ClearSessionTokenCookie(ctx *gin.Context) {
	ctx.SetSameSite(http.SameSiteLaxMode)

	ctx.SetCookie(
		SessionCookieName,
		"",
		-1,
		"/",
		"",
		config.CONFIG.HTTP.SecureCookies,
		true,
	)
}
