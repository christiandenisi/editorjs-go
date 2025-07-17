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

// TypedRenderer defines a renderer function for a typed block.
type TypedRenderer[T any] func(Block[T], *Context) (string, error)
