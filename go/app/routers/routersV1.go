package routers

import (
	"../handlers"
	"../middleware"
	"github.com/gofiber/fiber/v2"
)

func Version1(app *fiber.App, tokenKey string) {
	app.Get("/api/v1", handlers.Welcome)
	setupCompaniesRoutes(app, tokenKey)
	setupUsersRoutes(app, tokenKey)
	setupArticlesRoutes(app, tokenKey)
}

func setupUsersRoutes(app *fiber.App, tokenKey string) {
	users := app.Group("/api/v1/users")
	users.Post("/login", handlers.LoginUser)
	users.Use(middleware.JwtMiddleware(tokenKey))
	users.Post("/", handlers.SaveUser)
	users.Get("/:id", handlers.GetUser)
	users.Put("/:id", handlers.UpdateUser)
	users.Delete("/:id", handlers.DeleteUser)
	users.Get("/", handlers.GetUsers)
}

func setupCompaniesRoutes(app *fiber.App, tokenKey string) {
	companies := app.Group("/api/v1/companies")
	companies.Use(middleware.JwtMiddleware(tokenKey))
	companies.Get("/", handlers.GetCompanies)
	companies.Get("/:rnc", handlers.GetCompany)
	//companies.Get("/:rnc/articles", handlers.GetArticlesByCompany)
	//companies.Get("/:rnc/:branch_offices/articles", handlers.GetArticlesByBranchOffices)
}

func setupArticlesRoutes(app *fiber.App, tokenKey string) {
	articles := app.Group("/api/v1/:rnc/:branch_office/articles")
	articles.Use(middleware.JwtMiddleware(tokenKey))
	articles.Get("/", handlers.GetArticles)
	articles.Post("/", handlers.SaveArticle)
	articles.Get("/:id", handlers.GetArticle)
	articles.Delete("/:id", handlers.DeleteArticle)
	articles.Put("/:id", handlers.UpdateArticle)
}
