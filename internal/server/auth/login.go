package auth

import (
	"b3_ux_backend/internal/db"
	"b3_ux_backend/internal/entities"
	"b3_ux_backend/internal/entities/user"
	"b3_ux_backend/internal/logx"
	"b3_ux_backend/internal/server/core"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/olekukonko/errors"
	"go.uber.org/zap"
)

func SignIn(ctx *gin.Context) {
	type queryParameters struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
		IsAdmin  bool   `json:"is_admin"`
	}

	var params queryParameters
	if err := ctx.ShouldBind(&params); err != nil {
		core.HandleError(ctx, http.StatusBadRequest, "Could not parse credentials", err)
		return
	}

	exists, err := db.Client.User.Query().
		Where(user.Or(user.UsernameEQ(params.Username), user.Email(params.Email))).
		Exist(ctx)
	if err != nil {
		core.HandleError(ctx, http.StatusInternalServerError, "Could check if user exists", err)
		return
	}
	if exists {
		core.HandleError(ctx, http.StatusBadRequest, "User already exists", err)
		return
	}

	userEntity, err := db.Client.User.Create().
		SetUsername(params.Username).
		SetEmail(params.Email).
		SetPassword(params.Password).
		SetIsAdmin(params.IsAdmin).
		Save(ctx)
	if err != nil {
		core.HandleError(ctx, http.StatusInternalServerError, "Could not create user", err)
		return
	}

	sessionToken, err := NewSession(ctx, userEntity.ID)
	if err != nil {
		core.HandleError(ctx, http.StatusInternalServerError, "Could not create session", err)
		return
	}

	SetSessionTokenCookie(ctx, sessionToken)
	logx.Logger.Debug("/login",
		zap.String(SessionCookieName, ctx.GetHeader(SessionCookieName)),
		zap.String(UserAuthContextKey, ctx.GetHeader(UserAuthContextKey)),
	)

	ctx.JSON(http.StatusOK, userEntity.ToData())
}

func Login(ctx *gin.Context) {
	type queryParameters struct {
		Username string `form:"username" binding:"required" json:"username"`
		Password string `form:"password" binding:"required" json:"password"`
	}

	var params queryParameters
	if err := ctx.ShouldBind(&params); err != nil {
		core.HandleError(ctx, http.StatusBadRequest, "Could not parse credentials", err)
		return
	}

	entityArtist, err := db.Client.User.Query().
		Where(user.UsernameEQ(params.Username), user.PasswordEQ(params.Password)).
		Only(ctx)
	if err != nil {
		if entities.IsNotFound(err) {
			core.HandleError(ctx, http.StatusUnauthorized, "Bad credentials", err)
			return
		}

		core.HandleError(ctx, http.StatusInternalServerError, "Could not authenticate", err)
		return
	}

	sessionToken, err := NewSession(ctx, entityArtist.ID)
	if err != nil {
		core.HandleError(ctx, http.StatusInternalServerError, "Could not create session", err)
		return
	}

	SetSessionTokenCookie(ctx, sessionToken)
	logx.Logger.Debug("/login",
		zap.String(SessionCookieName, ctx.GetHeader(SessionCookieName)),
		zap.String(UserAuthContextKey, ctx.GetHeader(UserAuthContextKey)),
	)

	ctx.JSON(http.StatusOK, entityArtist.ToData())
}

func Logout(ctx *gin.Context) {
	token, err := ctx.Cookie(SessionCookieName)
	if err != nil && !errors.Is(err, http.ErrNoCookie) {
		core.HandleError(ctx, http.StatusInternalServerError, "Could not read session", err)
		return
	}

	if token != "" {
		if err := DeleteOneSession(ctx, token); err != nil {
			core.HandleError(ctx, http.StatusInternalServerError, "Could not close session", err)
			return
		}
	}

	ClearSessionTokenCookie(ctx)

	ctx.Status(http.StatusNoContent)
}
