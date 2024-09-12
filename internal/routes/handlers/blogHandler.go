package handlers

import (
	"Blog/internal/db"
	"Blog/internal/services"
	"Blog/pkg/logger"
	"html/template"
	"net/http"

	"github.com/google/uuid"
)

type pageData struct {
	Posts []db.Post
}

var (
	allBlogsTemplate = template.Must(template.ParseFiles("cmd/web/html/pages/blogs.html"))
	blogTemplate     = template.Must(template.ParseFiles("cmd/web/html/pages/blog.html"))
)

// BlogHandler holds dependencies for blog routes
type BlogHandler struct {
	BaseHandler BaseHandlerInterface
	Service     *services.BlogService
}

// NewBlogHandler creates a new blog handler
func NewBlogHandler(service *services.BlogService, baseHandler BaseHandlerInterface) *BlogHandler {
	return &BlogHandler{
		BaseHandler: baseHandler,
		Service:     service,
	}
}

func (h *BlogHandler) AllBlogs(w http.ResponseWriter, r *http.Request) {
	posts, err := h.Service.GetAllPosts() // This should return []Post and error
	if err != nil {
		logger.Error("Error fetching blog posts: %e", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	pageData := pageData{
		Posts: posts,
	}

	h.BaseHandler.RenderPage(w, r, "Blogs", allBlogsTemplate, pageData)
}

func (h *BlogHandler) BlogByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		logger.Error("Invalid Blog Id: %s", idStr)
		http.Error(w, "Invalid blog ID", http.StatusBadRequest)
		return
	}

	post, err := h.Service.GetPostByID(id)
	if err != nil {
		// Assuming getPostByID returns an error when the post is not found
		logger.Error("Blog post not found")
		http.Error(w, "Blog post not found", http.StatusNotFound)
		return
	}

	h.BaseHandler.RenderPage(w, r, "Blog", blogTemplate, post)
}
