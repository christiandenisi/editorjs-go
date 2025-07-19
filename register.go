package editorjs

import (
	"encoding/json"
	"fmt"
)

// RegisterTyped adds a decoder and typed renderer for a block type.
func Register[T any](c *Converter, blockType string, renderer Renderer[T]) {
	c.decoders[blockType] = func(rb RawBlock) (any, error) {
		return unmarshalBlock[T](blockType, rb)
	}

	c.renderers[blockType] = func(block any, ctx *Context) (string, error) {
		typed, ok := block.(*Block[T])
		if !ok {
			return "", fmt.Errorf("internal error: wrong type for block %q", blockType)
		}
		return renderer(*typed, ctx)
	}
}

func unmarshalBlock[Data any](blockType string, rb RawBlock) (*Block[Data], error) {
	var data Data
	if err := json.Unmarshal(rb.Data, &data); err != nil {
		return nil, fmt.Errorf("invalid data for block %q: %w", blockType, err)
	}
	return &Block[Data]{
		ID:    rb.ID,
		Type:  rb.Type,
		Data:  data,
		Tunes: rb.Tunes,
	}, nil
}
