package blog

import (
	"Blog/pkg/logger"
	"fmt"
	"net/http"
	"strconv"
)

// blogService defines the interface for our blog service
type blogService interface {
	getAllPosts() ([]Post, error)
	getPostByID(id int) (Post, error)
}

// Handler holds dependencies for blog routes
type Handler struct {
	Service blogService
}

// NewHandler creates a new blog handler
func NewHandler(service blogService) *Handler {
	return &Handler{
		Service: service,
	}
}

// ShowBlogPosts displays a list of blog posts
func (h *Handler) ShowBlogPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := h.Service.getAllPosts()
	if err != nil {
		logger.LogError.Println("Failed to fetch posts", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// Assume posts are rendered to HTML
	fmt.Fprintf(w, "Posts: %+v", posts)
}

// ShowBlogPost displays a specific blog post by ID
func (h *Handler) ShowBlogPost(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/blog/post/"):])
	if err != nil {
		logger.LogError.Println("Invalid post ID", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	post, err := h.Service.getPostByID(id)
	if err != nil {
		logger.LogError.Println("Failed to fetch post", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// Assume post is rendered to HTML
	fmt.Fprintf(w, "Post: %+v", post)
}
