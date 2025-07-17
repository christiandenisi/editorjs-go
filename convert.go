package editorjs

import (
	"encoding/json"
	"strings"
)

// Convert parses Editor.js JSON and returns rendered HTML.
func (c *Converter) Convert(jsonData []byte) (string, error) {
	var doc RawDocument
	if err := json.Unmarshal(jsonData, &doc); err != nil {
		return "", err
	}

	ctx := &Context{
		EscapeHTML: EscapeHTML,
		Converter:  c,
	}

	ctx.Render = func(blocks []RawBlock) (string, error) {
		var b strings.Builder
		for _, rb := range blocks {
			html, err := c.renderBlock(rb, ctx)
			if err != nil {
				return "", err
			}
			b.WriteString(html)
		}
		return b.String(), nil
	}

	var out strings.Builder
	for _, rb := range doc.Blocks {
		html, err := c.renderBlock(rb, ctx)
		if err != nil {
			return "", err
		}
		out.WriteString(html)
	}
	return out.String(), nil
}
