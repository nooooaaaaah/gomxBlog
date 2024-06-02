package home

import (
	"Blog/internal/base" // Make sure the import path is correct
	"Blog/pkg/github"
	"Blog/pkg/logger"
	"html/template"
	"net/http"
)

var (
	homeTemplate    = template.Must(template.ParseFiles("ui/html/pages/home.html", "ui/html/partials/card.html"))
	aboutTemplate   = template.Must(template.ParseFiles("ui/html/pages/about.html"))
	contactTemplate = template.Must(template.ParseFiles("ui/html/pages/contact.html"))
)

type HomeHandler struct {
	BaseHandler base.BaseHandlerInterface
	Service     *HomeService
}

type hompageData struct {
	GhPro       github.GitHubProfile
	PinnedPosts []github.Repo
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
		logger.LogError.Println("Something bad happened ", err)
	}

	data := hompageData{
		GhPro:       *ghInfo.GhPro,
		PinnedPosts: ghInfo.PinnedRepos,
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
