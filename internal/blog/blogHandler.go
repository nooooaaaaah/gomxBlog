package blog

import (
	"Blog/internal/base"
	"Blog/pkg/db"
	"Blog/pkg/logger"
	"html/template"
	"net/http"

	"github.com/edgedb/edgedb-go"
)

type PageData struct {
	Posts []db.Post
}

var (
	allBlogsTemplate = template.Must(template.ParseFiles("ui/html/pages/blogs.html"))
	blogTemplate     = template.Must(template.ParseFiles("ui/html/pages/blog.html"))
)

// BlogHandler holds dependencies for blog routes
type BlogHandler struct {
	BaseHandler base.BaseHandlerInterface
	Service     *BlogService
}

// NewBlogHandler creates a new blog handler
func NewBlogHandler(service *BlogService, baseHandler base.BaseHandlerInterface) *BlogHandler {
	return &BlogHandler{
		BaseHandler: baseHandler,
		Service:     service,
	}
}

func (h *BlogHandler) AllBlogs(w http.ResponseWriter, r *http.Request) {
	posts, err := h.Service.getAllPosts() // This should return []Post and error
	if err != nil {
		logger.LogError.Println("Error fetching blog posts: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	pageData := PageData{
		Posts: posts,
	}

	h.BaseHandler.RenderPage(w, r, "Blogs", allBlogsTemplate, pageData)
}

func (h *BlogHandler) BlogByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := edgedb.ParseUUID(idStr)
	if err != nil {
		logger.LogError.Println("Invalid Blog Id: ", idStr)
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
