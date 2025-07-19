package editorjs

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Converter manages decoders and renderers for Editor.js blocks.
type Converter struct {
	decoders  map[string]decoderFn
	renderers map[string]rendererFn
}

// New creates a new Converter instance.
func New() *Converter {
	return &Converter{
		decoders:  make(map[string]decoderFn),
		renderers: make(map[string]rendererFn),
	}
}

// Convert parses Editor.js JSON and returns rendered HTML.
func (c *Converter) Convert(jsonData []byte) (string, error) {
	var doc RawDocument
	if err := json.Unmarshal(jsonData, &doc); err != nil {
		return "", err
	}

	ctx := &Context{
		Converter: c,
	}

	ctx.RenderBlock = func(block RawBlock) (string, error) {
		return c.renderBlock(block, ctx)
	}

	ctx.RenderBlocks = func(blocks []RawBlock) (string, error) {
		return c.renderBlocks(blocks, ctx)
	}

	return c.renderBlocks(doc.Blocks, ctx)

}

// renderBlocks renders a slice of RawBlock using the provided context.
func (c *Converter) renderBlocks(blocks []RawBlock, ctx *Context) (string, error) {
	var out strings.Builder
	for _, rb := range blocks {
		html, err := c.renderBlock(rb, ctx)
		if err != nil {
			return "", err
		}
		out.WriteString(html)
	}
	return out.String(), nil
}

// renderBlock decodes and renders a single block.
func (c *Converter) renderBlock(rb RawBlock, ctx *Context) (string, error) {
	decoder, ok := c.decoders[rb.Type]
	if !ok {
		return "", fmt.Errorf("no decoder for block type %q", rb.Type)
	}
	renderer, ok := c.renderers[rb.Type]
	if !ok {
		return "", fmt.Errorf("no renderer for block type %q", rb.Type)
	}
	block, err := decoder(rb)
	if err != nil {
		return "", err
	}
	return renderer(block, ctx)
}
