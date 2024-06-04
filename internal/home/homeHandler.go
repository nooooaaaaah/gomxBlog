package home

import (
	"Blog/internal/base" // Make sure the import path is correct
	"Blog/pkg/db"
	"Blog/pkg/github"
	"Blog/pkg/logger"
	"html/template"
	"net/http"
)

var (
	homeTemplate = template.Must(template.ParseFiles(
		"ui/html/pages/home.html",
		"ui/html/partials/pinnedRepos.html"))
	aboutTemplate   = template.Must(template.ParseFiles("ui/html/pages/about.html"))
	contactTemplate = template.Must(template.ParseFiles("ui/html/pages/contact.html"))
)

type HomeHandler struct {
	BaseHandler base.BaseHandlerInterface
	Service     *HomeService
}

type hompageData struct {
	GhPro       github.GitHubProfile
	PinnedRepos []github.Repo
	Posts       []db.Post
}

func NewHomeHandler(service *HomeService, baseHandler base.BaseHandlerInterface) *HomeHandler {
	return &HomeHandler{
		BaseHandler: baseHandler,
		Service:     service,
	}
}

// Home page handler

func (h *HomeHandler) HomePage(w http.ResponseWriter, r *http.Request) {
	ghInfo, err := h.Service.GetCachedGhInfo()
	if err != nil {
		logger.LogError.Println("Error getting GitHub info:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	posts, err := h.Service.getBlogs()
	if err != nil {
		logger.LogError.Println("Error getting blog posts:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := hompageData{
		GhPro:       *ghInfo.GhPro,
		PinnedRepos: ghInfo.PinnedRepos,
		Posts:       posts,
	}
	h.BaseHandler.RenderPage(w, r, "Home", homeTemplate, data)
}

// About page handler
func (h *HomeHandler) AboutPage(w http.ResponseWriter, r *http.Request) {
	h.BaseHandler.RenderPage(w, r, "About Me", aboutTemplate, nil)
}

// Contact page handler
func (h *HomeHandler) ContactPage(w http.ResponseWriter, r *http.Request) {
	h.BaseHandler.RenderPage(w, r, "Contact Me", contactTemplate, nil)
}
