package auth

import (
	"b3_ux_backend/internal/server/core"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/olekukonko/errors"
)

func Me(ctx *gin.Context) {
	value, exists := ctx.Get(UserAuthContextKey)
	if !exists {
		core.HandleWarning(ctx, http.StatusInternalServerError, "no auth cookie", errors.New("no auth cookie"))
		return
	}

	result, ok := value.(*UserAuthData)
	if !ok || result == nil {
		core.HandleError(ctx, http.StatusInternalServerError, "invalid auth cookie data", errors.New("could not cast auth cookie data"))
		return
	}

	ctx.JSON(http.StatusOK, result)
}
