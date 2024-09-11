package routes

import (
	addpost "Blog/internal/addPost"
	"Blog/internal/blog"
	"Blog/internal/home"
	"Blog/pkg/logger"
	"net/http"
)

type AppHandlers struct {
	HomeHandler *home.HomeHandler
	BlogHandler *blog.BlogHandler
	PostHandler *addpost.PostHandler
}

func SetupRoutes(handlers AppHandlers) *http.ServeMux {
	mux := http.NewServeMux()
	//HomeHandler routes
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Serving the HomePage")
		handlers.HomeHandler.HomePage(w, r)
	})
	mux.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Serving the HomePage")
		handlers.HomeHandler.HomePage(w, r)
	})
	mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Serving the AboutPage")
		handlers.HomeHandler.AboutPage(w, r)
	})
	mux.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Serving the ContactPage")
		handlers.HomeHandler.ContactPage(w, r)
	})

	mux.HandleFunc("/addapost", func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Serving the Upload page")
		handlers.PostHandler.ShowUploadPage(w, r)
	})

	// BlogHandler routes
	mux.HandleFunc("/blogs", func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Serving the Blogs page")
		handlers.BlogHandler.AllBlogs(w, r)
	})
	mux.HandleFunc("GET /blogs/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Retrieve the 'id' from the path
		id := r.PathValue("id")
		logger.Info("Displaying blog post with ID: %s\n", id)
		handlers.BlogHandler.BlogByID(w, r)
	})

	// robot and sitemap
	mux.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "robots.txt")
	})
	mux.HandleFunc("/sitemap.xml", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "sitemap.xml")
	})

	// Serve static files
	fileServer := http.FileServer(http.Dir("ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
	logger.Info("static files served")

	return mux
}
