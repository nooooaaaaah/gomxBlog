package db

import (
	"Blog/pkg/utils"
	"html/template"

	edgedb "github.com/edgedb/edgedb-go"
)

type Post struct {
	Id          edgedb.UUID             `edgedb:"id"`
	Title       string                  `edgedb:"title"`
	Content     string                  `edgedb:"content"`
	Description string                  `edgedb:"description"`
	Link        string                  `edgedb:"link"`
	PublishedOn edgedb.OptionalDateTime `edgedb:"published_on"`
}

type Category struct {
	Id    edgedb.UUID `edgedb:"id"`
	Name  string      `edgedb:"name"`
	Posts []Post      `edgedb:"posts"`
}

func (p *Post) PublishedDate() string {
	t, _ := p.PublishedOn.Get()
	return t.Format("2006-01-02 15:04:05")
}

func (p *Post) HtmlContent() template.HTML {
	return utils.MdToHTML(p.Content)
}
