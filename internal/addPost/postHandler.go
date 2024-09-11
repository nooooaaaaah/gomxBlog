package addpost

import (
	"Blog/internal/base"
	"Blog/pkg/db"
	"html/template"
	"io"
	"net/http"
	"time"
)

var uploadTemplate = template.Must(template.ParseFiles("ui/html/pages/upload.html"))

type PostHandler struct {
	BaseHandler base.BaseHandlerInterface
	Service     *PostService
}

func NewPostHandler(service *PostService, baseHandler base.BaseHandlerInterface) *PostHandler {
	return &PostHandler{
		BaseHandler: baseHandler,
		Service:     service,
	}
}
func (h *PostHandler) ShowUploadPage(w http.ResponseWriter, r *http.Request) {

	h.BaseHandler.RenderPage(w, r, "upload", uploadTemplate, nil)
}

func (h *PostHandler) AddPost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	publishedOnStr := r.FormValue("date")
	publishedOn, err := parsePublishedDate(publishedOnStr)
	if err != nil {
		http.Error(w, "Invalid date format. Expected YYYY-MM-DD", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("markdown")
	if err != nil {
		http.Error(w, "Failed to get markdown file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Failed to read markdown file", http.StatusInternalServerError)
		return
	}

	post := db.Post{
		Title:       r.FormValue("title"),
		Content:     string(content),
		Description: r.FormValue("description"),
		PublishedOn: publishedOn,
	}

	if err := h.Service.MakePost(post); err != nil {
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Blog post uploaded successfully"))
}

func parsePublishedDate(dateStr string) (time.Time, error) {
	return time.Parse("2006-01-02", dateStr)
}
