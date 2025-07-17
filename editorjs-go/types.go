package editorjsgo

// Block represents a single Editor.js content block.
type Block struct {
	ID    string                 `json:"id"`
	Type  string                 `json:"type"`
	Data  map[string]interface{} `json:"data"`
	Tunes map[string]interface{} `json:"tunes,omitempty"`
}

// EditorJSDocument represents the full Editor.js document.
type EditorJSDocument struct {
	Time    int64   `json:"time"`
	Version string  `json:"version"`
	Blocks  []Block `json:"blocks"`
}
