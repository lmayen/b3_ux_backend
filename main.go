package main

import (
	"b3_ux_backend/internal/assets"
	"b3_ux_backend/internal/config"
	"b3_ux_backend/internal/db"
	"b3_ux_backend/internal/logx"
	"b3_ux_backend/internal/server"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

func main() {
	ctx, appCancel := context.WithCancel(context.Background())

	// --------------- //
	// --- Logger --- //
	// ------------- //
	err := logx.InitLogger()
	if err != nil {
		panic(err)
	}
	logx.Logger.Info("initialization of the logger is successful")

	// --------------- //
	// --- Config --- //
	// ------------- //
	cfg, err := config.InitConfig()
	if err != nil {
		logx.Logger.Fatal("app configuration initialization failed", zap.Error(err))
		log.Fatal(err)
	}
	logx.Logger.Info("initialization of the app configuration is successful")

	// ----------------- //
	// --- Database --- //
	// --------------- //
	_, err = db.InitDb(ctx, cfg)
	if err != nil {
		logx.Logger.Fatal("database initialization failed", zap.Error(err))
		panic(err)
	}
	logx.Logger.Info("initialization of the database is successful")

	// ----------------//
	// --- Server --- //
	// ------------- //
	shutdownCh := make(chan struct{})
	//_, httpServer, err := server.InitServer(cfg, probe, shutdownCh)
	_, httpServer, err := server.InitServer(cfg, shutdownCh)
	if err != nil {
		logx.Logger.Fatal("server initialization failed", zap.Error(err))
		panic(err)
	}
	logx.Logger.Info("initialization of http server is successful")

	// ------------------------ //
	// --- Embedded Assets --- //
	// ---------------------- //
	err = assets.InitAssets()
	if err != nil {
		logx.Logger.Fatal("asset initialization failed", zap.Error(err))
		panic(err)
	}
	logx.Logger.Info("initialization of embedded assets is successful")

	// -------------- //
	// --- Start --- //
	// ------------ //
	logx.Logger.Info("app is running...")
	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logx.Logger.Fatal("http server initialization failed", zap.Error(err))
			log.Fatal("http server failed", err)
		}
	}()

	// ----------------- //
	// --- Shutdown --- //
	// --------------- //
	waitForShutdown(httpServer, appCancel, shutdownCh)
}

func waitForShutdown(srv *http.Server, appCancel context.CancelFunc, shutdownCh <-chan struct{}) {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(sigCh)

	select {
	case sig := <-sigCh:
		logx.Logger.Info("shutdown requested (os signal)", zap.String("signal", sig.String()))
	case <-shutdownCh:
		logx.Logger.Info("shutdown requested (api)")
	}

	// Global shutdown timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Stop accepting new HTTP requests
	if err := srv.Shutdown(ctx); err != nil {
		logx.Logger.Error("shutdown error", zap.Error(err))
	}

	// Cancel app context (propagates to workers & tasks)
	appCancel()

	// Close DB after server shutdown
	if err := db.Client.Close(); err != nil {
		logx.Logger.Error("db close error", zap.Error(err))
	}

	logx.Logger.Info("app shutdown complete...")
}
