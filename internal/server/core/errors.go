package core

import (
	"b3_ux_backend/internal/logx"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func HandleWarning(ctx *gin.Context, code int, msg string, err error) {
	if err != nil {
		logx.Logger.Warn(msg, zap.Error(err))
		ctx.JSON(code, gin.H{"client_error": msg})
	} else {
		logx.Logger.Warn(msg)
		ctx.JSON(code, gin.H{"client_error": msg})
	}
}

func HandleError(ctx *gin.Context, code int, msg string, err error) {
	if err != nil {
		logx.Logger.Error(msg, zap.Error(err))
		ctx.JSON(code, gin.H{"client_error": msg})
	} else {
		logx.Logger.Error(msg)
		ctx.JSON(code, gin.H{"client_error": msg})
	}
}
