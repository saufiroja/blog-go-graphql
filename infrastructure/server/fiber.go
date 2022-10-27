package server

import (
	"graphql/blog-go-graphql/database"
	"graphql/blog-go-graphql/infrastructure/routers"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

func InitServer() *fiber.App {
	app := fiber.New()
	conf := database.Config{}
	database.InitDatabase(&conf)

	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	handler := routers.Routes(app, conf)

	app.Get("/graphql", func(c *fiber.Ctx) error {
		fasthttpadaptor.NewFastHTTPHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			handler.ServeHTTP(w, r)
		})(c.Context())
		return nil
	})

	app.Post("/graphql", func(c *fiber.Ctx) error {
		fasthttpadaptor.NewFastHTTPHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			handler.ServeHTTP(w, r)
		})(c.Context())
		return nil
	})

	return app
}
