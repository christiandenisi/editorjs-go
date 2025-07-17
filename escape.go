package editorjs

import "html"

// EscapeHTML escapes user input for safe HTML output.
func EscapeHTML(s string) string {
	return html.EscapeString(s)
}
