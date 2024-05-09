package main

import (
	"Blog/internal/base"
	"Blog/internal/home"
	"Blog/pkg/logger"
	"net/http"
)

type AppHandlers struct {
	BaseHandler *base.BaseHandler
	HomeHandler *home.HomeHandler
}

func setupRoutes(handlers AppHandlers) *http.ServeMux {
	mux := http.NewServeMux()
	//Home
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logger.LogInfo.Println("Serving the HomePage")
		handlers.HomeHandler.HomePage(w, r)
	})
	mux.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		logger.LogInfo.Println("Serving the HomePage")
		handlers.HomeHandler.HomePage(w, r)
	})
	mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		logger.LogInfo.Println("Serving the AboutPage")
		handlers.HomeHandler.AboutPage(w, r)
	})
	mux.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		logger.LogInfo.Println("Serving the ContactPage")
		handlers.HomeHandler.ContactPage(w, r)
	})

	// // Blog routes
	// mux.HandleFunc("/blog", handlers.BlogHandler.ShowBlogPosts)
	// mux.HandleFunc("/blog/post/", handlers.BlogHandler.ShowBlogPost)

	// Serve static files
	fileServer := http.FileServer(http.Dir("ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
	logger.LogInfo.Println("static files served")
	return mux
}
