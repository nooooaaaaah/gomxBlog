package base

import (
	"Blog/pkg/logger"
	"bytes"
	"html/template"
	"net/http"
	"time"
)

type BaseHandler struct {
	BaseTemplate *template.Template
}

type PageData struct {
	Title       string
	Year        int
	SidebarOpen bool
	Links       []Link
	Content     template.HTML
}

type Link struct {
	Href string
	Text string
}

// Constructor for the BaseHandler
func NewBaseHandler(baseTemplatePath string, partials ...string) *BaseHandler {
	allPaths := append([]string{baseTemplatePath}, partials...)
	tmpl, err := template.ParseFiles(allPaths...)
	if err != nil {
		logger.LogError.Printf("Error parsing template files: %v", err)
		return nil
	}
	return &BaseHandler{
		BaseTemplate: tmpl,
	}
}

// Method to handle rendering logic for pages
func (bh *BaseHandler) RenderPage(w http.ResponseWriter, r *http.Request, pageTitle string, contentTemplate *template.Template, partialData interface{}) {
	logger.LogInfo.Printf("Rendering page: %s", pageTitle)
	content, err := renderToString(contentTemplate, partialData)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		logger.LogError.Printf("Error rendering content: %v", err)
		return
	}

	// Decide whether to render just the content or the whole page
	if isHTMXRequest(r) {
		w.Write([]byte(content))
	} else {
		bh.RenderFullPage(w, r, pageTitle, content)
	}
}

// Renders the full page using the partial content
func (bh *BaseHandler) RenderFullPage(w http.ResponseWriter, r *http.Request, pageTitle, content string) {
	data := &PageData{
		Title:       pageTitle,
		Year:        time.Now().Year(),
		SidebarOpen: r.URL.Query().Get("sidebar") == "open",
		Links:       getDefaultLinks(),
		Content:     template.HTML(content),
	}

	if err := bh.BaseTemplate.ExecuteTemplate(w, "base.html", data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		logger.LogError.Printf("Error executing base template: %v", err)
	}
}

// Utilities
func getDefaultLinks() []Link {
	return []Link{
		{Href: "/home", Text: "Home"},
		{Href: "/blogs", Text: "Blogs"},
		{Href: "/about", Text: "About"},
		{Href: "/contact", Text: "Contact"},
	}
}

func renderToString(t *template.Template, data interface{}) (string, error) {
	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		logger.LogError.Printf("Error executing template to string: %v", err)
		return "", err
	}
	return buf.String(), nil
}

func isHTMXRequest(r *http.Request) bool {
	return r.Header.Get("HX-Request") != ""
}
