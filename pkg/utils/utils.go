package utils

import (
	"bytes"
	"html/template"
	"strings"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
	"golang.org/x/net/html"
)

// markdown input, and returns HTML directly
func MdToHTML(markdownInput string) template.HTML {

	unsafeHTML := blackfriday.Run([]byte(markdownInput))

	p := bluemonday.UGCPolicy()
	safeHTML := p.SanitizeBytes(unsafeHTML)

	doc, err := html.Parse(bytes.NewReader(safeHTML))
	if err != nil {
		panic("Failed to parse HTML")
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && strings.HasPrefix(n.Data, "h2") {
			classAttr := html.Attribute{Key: "class", Val: "text-xl font-semibold my-4"}
			n.Attr = append(n.Attr, classAttr)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	var buf bytes.Buffer
	html.Render(&buf, doc)
	return template.HTML(buf.String())
}
