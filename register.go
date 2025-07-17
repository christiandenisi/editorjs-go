package editorjs

import (
	"encoding/json"
	"fmt"
)

// RegisterTyped adds a decoder and typed renderer for a block type.
func RegisterTyped[T any](c *Converter, blockType string, renderer TypedRenderer[T]) {
	c.decoders[blockType] = func(rb RawBlock) (any, error) {
		var data T
		if err := json.Unmarshal(rb.Data, &data); err != nil {
			return nil, fmt.Errorf("invalid data for block %q: %w", blockType, err)
		}
		return Block[T]{
			ID:    rb.ID,
			Type:  rb.Type,
			Data:  data,
			Tunes: rb.Tunes,
		}, nil
	}

	c.renderers[blockType] = func(block any, ctx *Context) (string, error) {
		typed, ok := block.(Block[T])
		if !ok {
			return "", fmt.Errorf("internal error: wrong type for block %q", blockType)
		}
		return renderer(typed, ctx)
	}
}
