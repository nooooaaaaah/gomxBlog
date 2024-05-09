package main

import (
	"Blog/internal/blog"
	"Blog/internal/home"
	"Blog/pkg/logger"
	"net/http"
)

type AppHandlers struct {
	HomeHandler *home.HomeHandler
	BlogHandler *blog.BlogHandler
}

func setupRoutes(handlers AppHandlers) *http.ServeMux {
	mux := http.NewServeMux()
	//HomeHandler routes
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

	// BlogHandler routes
	mux.HandleFunc("/blogs", func(w http.ResponseWriter, r *http.Request) {
		logger.LogInfo.Println("Serving the Blogs page")
		handlers.BlogHandler.AllBlogs(w, r)
	})
	mux.HandleFunc("GET /blogs/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Retrieve the 'id' from the path
		id := r.PathValue("id")
		logger.LogInfo.Printf("Displaying blog post with ID: %s\n", id)
		handlers.BlogHandler.BlogByID(w, r)
	})

	// Serve static files
	fileServer := http.FileServer(http.Dir("ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
	logger.LogInfo.Println("static files served")
	return mux
}
