package routers

import (
	"graphql/blog-go-graphql/database"
	"graphql/blog-go-graphql/graphql/resolvers"
	"graphql/blog-go-graphql/graphql/schema"
	"graphql/blog-go-graphql/repository"
	"graphql/blog-go-graphql/service"

	"github.com/gofiber/fiber/v2"
	"github.com/graphql-go/handler"
)

func Routes(fiber *fiber.App, conf database.Config) *handler.Handler {
	db := database.InitDatabase(conf)

	// repositories
	userRepository := repository.NewUserRepository(db)

	// services
	userService := service.NewUserService(userRepository, conf)

	// resolvers
	userResolvers := resolvers.NewUserResolvers(userService)

	// schema
	schema := schema.NewUserSchema(userResolvers)

	// graphql
	gh := handler.Config{
		Schema:   schema.Root(),
		Pretty:   true,
		GraphiQL: true,
	}

	return handler.New(&gh)
}
