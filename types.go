package editorjs

import "encoding/json"

// RawBlock represents a raw block before type decoding.
type RawBlock struct {
	ID    string                 `json:"id"`
	Type  string                 `json:"type"`
	Data  json.RawMessage        `json:"data"`
	Tunes map[string]interface{} `json:"tunes,omitempty"`
}

// RawDocument represents a full Editor.js document.
type RawDocument struct {
	Time    int64      `json:"time"`
	Version string     `json:"version"`
	Blocks  []RawBlock `json:"blocks"`
}

// Block is a typed Editor.js block.
type Block[T any] struct {
	ID    string
	Type  string
	Data  T
	Tunes map[string]interface{}
}

// Context provides tools to renderer functions.
type Context struct {
	RenderBlocks func([]RawBlock) (string, error)
	RenderBlock  func(RawBlock) (string, error)
	Converter    *Converter
}

// Renderer defines a renderer function for a typed block.
type Renderer[T any] func(Block[T], *Context) (string, error)

// decoderFn is an internal decoder function type.
type decoderFn func(RawBlock) (any, error)

// rendererFn is an internal renderer function type.
type rendererFn func(any, *Context) (string, error)
