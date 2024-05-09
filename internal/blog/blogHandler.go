package blog

import (
	"Blog/internal/base"
	"Blog/pkg/logger"
	"html/template"
	"net/http"
	"strconv"
)

var (
	allBlogsTemplate = template.Must(template.ParseFiles("ui/html/pages/blogs.html"))
	blogTemplate     = template.Must(template.ParseFiles("ui/html/pages/blog.html"))
)

// blogService defines the interface for our blog service
type blogService interface {
	getAllPosts() ([]Post, error)
	getPostByID(id int) (Post, error)
}

// BlogHandler holds dependencies for blog routes
type BlogHandler struct {
	BaseHandler base.BaseHandlerInterface
	Service     blogService
}

// NewBlogHandler creates a new blog handler
func NewBlogHandler(service blogService, baseHandler base.BaseHandlerInterface) *BlogHandler {
	return &BlogHandler{
		BaseHandler: baseHandler,
		Service:     service,
	}
}

func (h *BlogHandler) AllBlogs(w http.ResponseWriter, r *http.Request) {
	h.BaseHandler.RenderPage(w, r, "Blogs", allBlogsTemplate, nil)
}

func (h *BlogHandler) BlogByID(w http.ResponseWriter, r *http.Request) {
	idstr := r.PathValue("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		logger.LogError.Println("Invalid Blog Id: ", idstr)
		http.Error(w, "Invalid blog ID", http.StatusBadRequest)
		return
	}

	post, err := h.Service.getPostByID(id)
	if err != nil {
		// Assuming getPostByID returns an error when the post is not found
		logger.LogError.Println("Blog post not found")
		http.Error(w, "Blog post not found", http.StatusNotFound)
		return
	}

	h.BaseHandler.RenderPage(w, r, "Blog", blogTemplate, post)
}
