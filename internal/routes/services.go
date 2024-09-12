package routes

import (
	"Blog/internal/db"
	"Blog/internal/routes/handlers"
	"Blog/internal/services"
)

type Services struct {
	queries *db.Queries
}

func New(queries *db.Queries) *Services {
	return &Services{
		queries: queries,
	}
}

func (s *Services) CreateHandlers() AppHandlers {
	baseHandler := createBaseHandler()
	blogService := services.NewBlogService(s.queries)
	homeService := services.NewHomeService(blogService)
	uploadService := services.NewPostSerivce(*s.queries)

	return AppHandlers{
		HomeHandler: handlers.NewHomeHandler(homeService, baseHandler),
		BlogHandler: handlers.NewBlogHandler(blogService, baseHandler),
		PostHandler: handlers.NewPostHandler(uploadService, baseHandler),
	}
}

func createBaseHandler() handlers.BaseHandlerInterface {
	partials := []string{
		"cmd/web/html/layouts/footer.html",
		"cmd/web/html/layouts/header.html",
		"cmd/web/html/layouts/navbar.html",
	}
	return handlers.NewBaseHandler(
		"cmd/web/html/layouts/base.html",
		partials...,
	)
}
