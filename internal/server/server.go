package server

import (
	"os"
	"os/signal"

	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/priyanshu334/go_car_book/internal/config"
	"github.com/priyanshu334/go_car_book/internal/db"
	"github.com/priyanshu334/go_car_book/internal/logger"
	"github.com/priyanshu334/go_car_book/internal/modules/user"
	"go.uber.org/zap"
)

func Start() {
	config.Load()
	logger.Init(config.Cfg.AppEnv)

	if err := db.Connect(); err != nil {
		logger.Log.Fatal("db connection failed")
	}

	if err := db.DB.AutoMigrate(&user.User{}); err != nil {
		logger.Log.Fatal("auto migration failed", zap.Error(err))
	}
	app := fiber.New()

	userRepo := user.NewRepository(db.DB)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)
	api := app.Group("/api")
	api.Post("/users/register", userHandler.Register)
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(
			fiber.Map{
				"status": "ok",
			},
		)

	})
	go func() {
		if err := app.Listen(":8080"); err != nil {
			panic(err)
		}

	}()
	quit :=
		make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Log.Info("graceful Shutdown")
	_ = app.Shutdown()
}
