package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"news/api"
	"news/api/gen"
	"news/auth"
	"news/config"
	"news/db/postgres"
	"os"
	"os/signal"
	// "news/db/mongo"
	"news/category"
	"news/device"
	"news/object"
	"news/user"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start API server",
	Long:  `Start API server`,
	Run:   RunCommand,
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

func RunCommand(cmd *cobra.Command, args []string) {

	config := config.Init()
	postgresDB := postgres.Init(config)
	// mongoDB := mongo.Init(config)

	e := echo.New()
	g := e.Group("/api/v1")

	jwtKey := config.JWTKey
	if len(jwtKey) == 0 {
		jwtKey = "MEZGMEFDRTYtRTEyRS00RTdGLUJCNjYtMTVGODc4RUE4NjIyCg=="
	}

	authenticator := api.NewAppAuthenticator(jwtKey)

	authMiddleware, err := api.CreateAuthMiddleware(authenticator)

	if err != nil {
		log.Fatal(err)
	}

	e.Use(middleware.CORS())
	g.Use(middleware.CORS())
	g.Use(authMiddleware...)
	g.Use(middleware.Logger())

	userRepo := user.NewRepo(postgresDB)

	userService := user.NewService(&userRepo)
	authService := auth.NewService(authenticator, userService)
	categoryRepo := category.NewRepo(postgresDB)
	deviceRepo := device.NewRepo(postgresDB)
	objectRepo := object.NewRepo(postgresDB)
	deviceService := device.NewService(&deviceRepo)
	objectService := object.NewService(&objectRepo)
	categoryService := category.NewService(&categoryRepo)

	serverImpl := api.NewServerImpl(
		authService,
		categoryService,
		deviceService,
		objectService,
	)

	gen.RegisterHandlersWithBaseURL(e, &serverImpl, "api/v1")

	go func() {
		if err := e.Start(fmt.Sprintf("0.0.0.0:%d", config.Port)); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		log.Println("Cannot shutdown")
		e.Logger.Fatal(err)
	}

	log.Println("Shutting down")
}
