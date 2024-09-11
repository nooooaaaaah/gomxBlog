package db

import (
	"Blog/pkg/utils"
	"html/template"
)

func (p *Post) PublishedDate() string {
	t := p.PublishedOn
	return t.Format("2006-01-02 15:04:05")
}

func (p *Post) HtmlContent() template.HTML {
	return utils.MdToHTML(p.Content)
}
