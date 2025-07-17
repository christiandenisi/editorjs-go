package editorjs

import (
	"fmt"
)

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
