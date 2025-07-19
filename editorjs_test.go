package editorjs

import (
	"html"
	"strings"
	"testing"
)

// --- Plugin: Paragraph ---
type ParagraphData struct {
	Text string `json:"text"`
}

func RenderParagraph(b Block[ParagraphData], ctx *Context) (string, error) {
	return "<p>" + html.EscapeString(b.Data.Text) + "</p>", nil
}

// --- Plugin: Quote (con blocchi interni) ---
type QuoteData struct {
	Items []RawBlock `json:"items"`
}

func RenderQuote(b Block[QuoteData], ctx *Context) (string, error) {
	content, err := ctx.RenderBlocks(b.Data.Items)
	if err != nil {
		return "", err
	}
	return "<blockquote>" + content + "</blockquote>", nil
}

// --- Plugin: List ---
type ListData struct {
	Items []string `json:"items"`
}

func RenderList(b Block[ListData], ctx *Context) (string, error) {
	var bld strings.Builder
	bld.WriteString("<ul>")
	for _, item := range b.Data.Items {
		bld.WriteString("<li>" + html.EscapeString(item) + "</li>")
	}
	bld.WriteString("</ul>")
	return bld.String(), nil
}

func TestRecursivePlugins(t *testing.T) {
	jsonData := []byte(`{
		"time": 1752781597903,
		"blocks": [
			{
				"id": "q1",
				"type": "quote",
				"data": {
					"items": [
						{
							"id": "p1",
							"type": "paragraph",
							"data": {
								"text": "Nested paragraph inside quote"
							}
						},
						{
							"id": "l1",
							"type": "list",
							"data": {
								"items": ["One", "Two", "Three"]
							}
						}
					]
				}
			}
		],
		"version": "2.27.0"
	}`)

	converter := New()

	// Register multiple plugins
	Register(converter, "paragraph", RenderParagraph)
	Register(converter, "quote", RenderQuote)
	Register(converter, "list", RenderList)

	// Execute rendering
	html, err := converter.Convert(jsonData)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := `<blockquote><p>Nested paragraph inside quote</p><ul><li>One</li><li>Two</li><li>Three</li></ul></blockquote>`

	normalize := func(s string) string {
		return strings.ReplaceAll(strings.TrimSpace(s), "\n", "")
	}

	if normalize(html) != normalize(expected) {
		t.Errorf("expected:\n%s\n\ngot:\n%s", expected, html)
	}
}
