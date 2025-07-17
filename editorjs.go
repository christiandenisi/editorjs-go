package editorjs

// decoderFn is an internal decoder function type.
type decoderFn func(RawBlock) (any, error)

// rendererFn is an internal renderer function type.
type rendererFn func(any, *Context) (string, error)

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
