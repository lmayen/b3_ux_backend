package main

import (
	"b3_ux_backend/internal/config"
	"b3_ux_backend/internal/db"
	"b3_ux_backend/internal/logx"
	"b3_ux_backend/internal/server"
	"context"
	"log"
	"net/http"

	"go.uber.org/zap"
)

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
func main() {
	//db.CreateAppEntities()

	err := logx.InitLogger()
	if err != nil {
		panic(err)
	}

	ctx, appCancel := context.WithCancel(context.Background())
	defer appCancel()

	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	client, err := db.InitDb(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// --- Server --- //
	shutdownCh := make(chan struct{})
	//_, httpServer, err := server.InitServer(cfg, probe, shutdownCh)
	_, httpServer, err := server.InitServer(cfg, shutdownCh)
	if err != nil {
		logx.Logger.Fatal("server initialization failed", zap.Error(err))
		panic(err)
	}
	logx.Logger.Info("init http server success")

	// --- Start --- //
	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("http server failed", err)
		}
	}()
}
