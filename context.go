package editorjs

// Context provides tools to renderer functions.
type Context struct {
	EscapeHTML func(string) string
	Render     func([]RawBlock) (string, error)
	Converter  *Converter
}
