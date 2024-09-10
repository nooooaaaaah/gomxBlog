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
		logger.Error("Error getting GitHub info:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	posts, err := h.Service.getBlogs()
	if err != nil {
		logger.Error("Error getting blog posts:", err)
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

type aboutMeData struct {
	Title                   string
	Greeting                string
	Introduction            string
	WorkDescription         string
	ContactInfo             string
	Email                   string
	CallToAction            string
	NextRoleQuestion        string
	NextRoleAnswer          string
	ProudProjectTitle       string
	ProudProjectDescription string
}

// About page handler
func (h *HomeHandler) AboutPage(w http.ResponseWriter, r *http.Request) {
	// TODO Fetch data from the database here
	data := aboutMeData{
		Title:                   "About Me!",
		Greeting:                "Hi! 🌟",
		Introduction:            "I'm Noah, a software engineer who loves tinkering with technology. I specialize in Go and C# development, and I have a knack for turning complex challenges into streamlined solutions, whether it's for websites, apps, or backend systems.",
		WorkDescription:         "I'm a freelance software engineer who enjoys solving challenging technical problems and creating impactful solutions. I'm always ready to customize my work to meet your specific needs.",
		ContactInfo:             "If you need help with technology – be it building a website, developing an app, or setting up robust backend systems – feel free to",
		Email:                   "nspielman96@gmail.com",
		CallToAction:            "I'm always up for a challenge and ready to help bring your ideas to life.",
		NextRoleQuestion:        "What am I looking for in my next role?",
		NextRoleAnswer:          "I'm looking forward to working with a team where I can contribute my skills and also learn from others' feedback. While I'm mainly experienced in backend technologies, I'm just as comfortable handling full-stack projects. Lately, I've been working with Go, Next.js, HTMX, and sqlite.",
		ProudProjectTitle:       "Project that I am proud of:",
		ProudProjectDescription: "I created a reservation system for a client who needed real-time updates and automated workflows to snag hard-to-get reservations on platforms like Resy. Using a containerized Go application, I achieved an impressive 80% success rate. I handled everything from gathering requirements to deployment, making sure all the client's needs were met and delivering clear, impactful results.",
	}
	h.BaseHandler.RenderPage(w, r, "About Me", aboutTemplate, data)
}

// Contact page handler
func (h *HomeHandler) ContactPage(w http.ResponseWriter, r *http.Request) {
	h.BaseHandler.RenderPage(w, r, "Contact Me", contactTemplate, nil)
}
