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
	db := database.InitDatabase(&conf)

	// repositories
	userRepository := repository.NewUserRepository(db)
	articleRepository := repository.NewArticleRepository(db)

	// services
	userService := service.NewUserService(userRepository)
	articleService := service.NewArticleService(articleRepository, conf)

	// resolvers
	userResolvers := resolvers.NewUserResolvers(userService)
	articleResolvers := resolvers.NewArticleResolvers(articleService)

	// schema
	userSchema := schema.NewUserSchema(userResolvers)
	articleSchema := schema.NewArticleSchema(articleResolvers)

	rootSchema := schema.NewRootSchema(userSchema, articleSchema)

	// graphql
	gh := handler.New(&handler.Config{
		Schema:     rootSchema.Root(),
		Pretty:     true,
		GraphiQL:   true,
		Playground: true,
	})

	return gh
}
