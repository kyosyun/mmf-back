package cmd

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kyosyun/mmf-back/api"
	"github.com/kyosyun/mmf-back/internal/env"
	"github.com/kyosyun/mmf-back/internal/infra/postgres"
	"github.com/kyosyun/mmf-back/internal/interface/handler"
	"github.com/spf13/cobra"
)

const (
	readHeaderTimeout       = 10 * time.Second
	gracefulShutdownTimeout = 1 * time.Minute
	healthCheckPath         = "/health"
)

var runCmd = &cobra.Command{
	Use:   "app",
	Short: "Run the application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start application server ...")
		serve()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func serveAddr() string {
	return ":8080"
}

func serve() {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	router := buildEngine(ctx)
	srv := &http.Server{
		Addr:              serveAddr(),
		Handler:           router,
		ReadHeaderTimeout: readHeaderTimeout,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("listen error:", "err", err)
			os.Exit(1)
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	slog.Info("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has gracefulShutdownTimeout to finish
	// the request it is currently handling
	// see:https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/APIReference/API_StopTask.html
	ctx, cancel := context.WithTimeout(context.Background(), gracefulShutdownTimeout)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("Server forced to shutdown: ", "err", err)
		os.Exit(1) //nolint:gocritic
	}
}

func buildEngine(ctx context.Context) *gin.Engine {

	environment := env.NewEnvironment()

	db, err := postgres.NewDatabase(environment.DBConfig)
	if err != nil {
		slog.Error("failed to connect to database", "err", err)
		os.Exit(1) //nolint:gocritic
	}
	repository := postgres.NewPostgresRepository(db)
	app := handler.NewApp(repository)
	router := gin.New()

	api.RegisterHandlersWithOptions(router, api.NewStrictHandler(app, nil), api.GinServerOptions{})

	return router
}
