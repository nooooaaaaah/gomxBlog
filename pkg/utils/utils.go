package utils

import (
	"bytes"
	"html/template"
	"regexp"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
	"golang.org/x/net/html"
)

// markdown input, and returns HTML directly
func MdToHTML(markdownInput string) template.HTML {

	unsafeHTML := blackfriday.Run([]byte(markdownInput))

	p := bluemonday.UGCPolicy()
	p.AllowAttrs("class").Matching(regexp.MustCompile("^language-[a-zA-Z0-9]+$")).OnElements("code")
	safeHTML := p.SanitizeBytes(unsafeHTML)

	doc, err := html.Parse(bytes.NewReader(safeHTML))
	if err != nil {
		panic("Failed to parse HTML")
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode {
			switch n.Data {
			case "h1":
				n.Attr = append(n.Attr, html.Attribute{Key: "class", Val: "text-3xl font-bold my-6"})
			case "h2":
				n.Attr = append(n.Attr, html.Attribute{Key: "class", Val: "text-3xl font-semibold my-4 text-purp-400"})
			case "h3":
				n.Attr = append(n.Attr, html.Attribute{Key: "class", Val: "text-2xl font-medium my-3 text-syan-200"})
			case "p":
				n.Attr = append(n.Attr, html.Attribute{Key: "class", Val: "my-2 "})
			case "a":
				n.Attr = append(n.Attr, html.Attribute{Key: "class", Val: "text-syan-400 underline"})
			case "ul":
				n.Attr = append(n.Attr, html.Attribute{Key: "class", Val: "list-disc ml-6 my-2"})
			case "ol":
				n.Attr = append(n.Attr, html.Attribute{Key: "class", Val: "list-decimal ml-6 my-2"})
			case "li":
				n.Attr = append(n.Attr, html.Attribute{Key: "class", Val: "my-1"})
			case "code":
				n.Attr = append(n.Attr, html.Attribute{Key: "class", Val: "rounded px-1 py-0.5"})
			case "pre":
				n.Attr = append(n.Attr, html.Attribute{Key: "class", Val: "p-4 rounded my-4"})
			}
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
