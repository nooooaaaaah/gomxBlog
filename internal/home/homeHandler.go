package home

import (
	"Blog/internal/base" // Make sure the import path is correct
	"html/template"
	"net/http"
)

var (
	homeTemplate    = template.Must(template.ParseFiles("ui/html/pages/home.html"))
	aboutTemplate   = template.Must(template.ParseFiles("ui/html/pages/about.html"))
	contactTemplate = template.Must(template.ParseFiles("ui/html/pages/contact.html"))
)

type HomeHandler struct {
	*base.BaseHandler
}

func NewHomeHandler(baseHandler *base.BaseHandler) *HomeHandler {
	return &HomeHandler{baseHandler}
}

// Home page handler
func (h *HomeHandler) HomePage(w http.ResponseWriter, r *http.Request) {
	// Remove redirect to prevent loops, let BaseHandler decide rendering method
	h.BaseHandler.RenderPage(w, r, "Home", homeTemplate, nil)
}

// About page handler
func (h *HomeHandler) AboutPage(w http.ResponseWriter, r *http.Request) {
	h.BaseHandler.RenderPage(w, r, "About Me", aboutTemplate, nil)
}

// Contact page handler
func (h *HomeHandler) ContactPage(w http.ResponseWriter, r *http.Request) {
	h.BaseHandler.RenderPage(w, r, "Contact Me", contactTemplate, nil)
}
