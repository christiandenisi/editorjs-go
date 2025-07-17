package editorjs

import (
	"testing"
)

// Temporary paragraph renderer for testing purposes.
type ParagraphData struct {
	Text string `json:"text"`
}

func RenderParagraph(b Block[ParagraphData], ctx *Context) (string, error) {
	return "<p>" + ctx.EscapeHTML(b.Data.Text) + "</p>", nil
}

func TestParagraphRender(t *testing.T) {
	json := []byte(`{
		"time": 1752781597903,
		"blocks": [
			{
				"id": "abc123",
				"type": "paragraph",
				"data": {
					"text": "Hello <b>world</b>"
				}
			}
		],
		"version": "2.27.0"
	}`)

	converter := New()

	RegisterTyped(converter, "paragraph", RenderParagraph)

	output, err := converter.Convert(json)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "<p>Hello &lt;b&gt;world&lt;/b&gt;</p>"
	if output != expected {
		t.Errorf("expected %q, got %q", expected, output)
	}
}
