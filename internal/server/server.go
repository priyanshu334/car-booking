package server

import (
	"os"
	"os/signal"

	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/priyanshu334/go_car_book/internal/config"
	"github.com/priyanshu334/go_car_book/internal/db"
	"github.com/priyanshu334/go_car_book/internal/logger"
	"github.com/priyanshu334/go_car_book/internal/middleware"
	"github.com/priyanshu334/go_car_book/internal/modules/booking"
	"github.com/priyanshu334/go_car_book/internal/modules/car"
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
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowCredentials: true,
		AllowHeaders:     "Origin, Content-Type, Accept",
	}))

	userRepo := user.NewRepository(db.DB)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)
	api := app.Group("/api")
	api.Post("/users/register", userHandler.Register)
	api.Post("/users/login", userHandler.Login)
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(
			fiber.Map{
				"status": "ok",
			},
		)

	})

	protected := app.Group("/protected")
	protected.Get("/me", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"user_id": c.Locals("user_id"),
			"role":    c.Locals("role"),
		})
	})
	db.DB.AutoMigrate(&car.Car{})
	carRepo := car.NewRepository(db.DB)
	carService := car.NewService(carRepo)
	carHandler := car.NewHandler(carService)
	carRoutes := app.Group("/cars")
	carRoutes.Get("/", carHandler.ListCars)
	carRoutes.Post("/", middleware.RequireAuth(), carHandler.AddCar)

	db.DB.AutoMigrate(&booking.Booking{})
	bookinRepo := booking.NewRepository(db.DB)
	bookingService := booking.NewService(bookinRepo, db.DB)

	bookingHandler := booking.NewHandler(bookingService)
	bookingRoutes := app.Group("/booking", middleware.RequireAuth())
	bookingRoutes.Get("/", bookingHandler.CreateBooking)
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
