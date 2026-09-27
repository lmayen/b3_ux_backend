package server

import (
	"b3_ux_backend/internal/config"
	"b3_ux_backend/internal/db"
	"b3_ux_backend/internal/entities/image"
	"b3_ux_backend/internal/fsutils"
	"b3_ux_backend/internal/server/core"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/olekukonko/errors"
)

func ServeFile(ctx *gin.Context) {
	// Id
	idStr := ctx.Param("id")
	fileID, err := uuid.Parse(idStr)
	if err != nil {
		core.HandleError(ctx, 400, "Bad parameters! Invalid file id", err)
		return
	}

	if fileID == uuid.Nil {
		core.HandleError(ctx, 400, "File id is Nil", nil)
		return
	}

	fileEntity, err := db.Client.Image.Query().
		Where(image.IDEQ(fileID)).
		Only(ctx)
	if err != nil {
		core.HandleError(ctx, 500, "Could not find the file entity", err)
		return
	}

	// Path
	absPath := filepath.Join(config.CONFIG.App.RootDir, fileEntity.Path)
	if !fsutils.Exists(absPath) {
		core.HandleError(ctx, 404, "File not found", errors.New("file not found: "+fileEntity.Path))
		return
	}

	// Stream
	f, err := os.Open(absPath)
	if err != nil {
		core.HandleError(ctx, 500, "Could not open file", err)
		return
	}
	defer f.Close()

	st, err := f.Stat()
	if err != nil {
		core.HandleError(ctx, 500, "Could not stat file", err)
		return
	}

	http.ServeContent(ctx.Writer, ctx.Request, st.Name(), st.ModTime(), f)
}
