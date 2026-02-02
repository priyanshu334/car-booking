package server

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
)

func Start() {
	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(
			fiber.Map{
				"status": "ok",
			},
		)

	})
	go func() {
		if err := app.Listen(":8000"); err != nil {
			panic(err)
		}

	}()
	quit :=
		make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	_ = app.Shutdown()
}
