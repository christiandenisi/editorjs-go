
# editorjs-go

**editorjs-go** is a flexible, type-safe, and extensible Go library for parsing and rendering [Editor.js](https://editorjs.io/) documents into HTML, Markdown, or any other output format.

## ✨ Features

- ✅ Type-safe decoding with generics (`Block[T]`)
- 🔌 Plugin-based architecture (register custom decoders/renderers)
- 🔄 Support for nested/recursive blocks
- 💡 Context-passing for render logic and helpers
- ⚙️ Zero-opinion on escaping/output format – **you decide** (HTML, Markdown, plaintext…)

---

## 📦 Installation

```bash
go get github.com/christiandenisi/editorjs-go
```

---

## 🚀 Quick Start

```go
package main

import (
    "fmt"
    "github.com/yourusername/editorjs-go"
    "html"
)

type ParagraphData struct {
    Text string `json:"text"`
}

func RenderParagraph(b editorjs.Block[ParagraphData], ctx *editorjs.Context) (string, error) {
    return "<p>" + html.EscapeString(b.Data.Text) + "</p>", nil
}

func main() {
    jsonData := []byte(`{
        "time": 1234567890,
        "version": "2.27.0",
        "blocks": [
            {
                "id": "abc",
                "type": "paragraph",
                "data": { "text": "Hello <world>" }
            }
        ]
    }`)

    conv := editorjs.New()
    editorjs.Register(conv, "paragraph", RenderParagraph)

    out, err := conv.Convert(jsonData)
    if err != nil {
        panic(err)
    }

    fmt.Println(out)
}
```

---

## 🧠 Architecture

```go
type Block[T any] struct {
    ID    string
    Type  string
    Data  T
    Tunes map[string]interface{}
}
```

- The library decodes raw JSON `RawBlock`s into strongly typed `Block[T]` using registered decoder functions.
- Each block type (e.g. `paragraph`, `quote`, `list`) is mapped to a decoder + renderer.

You register a block type using:

```go
func Register[T any](conv *Converter, blockType string, renderer Renderer[T])
```

---

## 🧱 Context Object

Render functions receive a `*Context`, giving access to:

```go
type Context struct {
    RenderBlock  func(RawBlock) (string, error)
    RenderBlocks func([]RawBlock) (string, error)
    Converter    *Converter
}
```

Use this for recursive rendering or accessing other converters.

---

## 🔄 Example: Recursive Block (Quote)

```go
type QuoteData struct {
    Items []editorjs.RawBlock `json:"items"`
}

func RenderQuote(b editorjs.Block[QuoteData], ctx *editorjs.Context) (string, error) {
    inner, err := ctx.RenderBlocks(b.Data.Items)
    if err != nil {
        return "", err
    }
    return "<blockquote>" + inner + "</blockquote>", nil
}
```

---

## ✍️ Implementing Custom Block Types

Each block needs:

- A data struct matching the expected JSON
- A renderer function with signature:

```go
func(Block[T], *Context) (string, error)
```

- Registration using `editorjs.Register(...)`

---

## 🔒 Escaping & Output Format

**Escaping is not handled automatically.**  
Each renderer decides whether to use escaping, and what kind (e.g. HTML, Markdown, plaintext):

```go
html.EscapeString(...)
```

This design allows editorjs-go to work for:
- HTML output
- Markdown generation
- Text-only rendering
- Custom formats (LaTeX, XML, etc.)

---

## ✅ Best Practices

- Always escape content in HTML renderers unless intentionally inserting raw HTML
- Use `RawBlock.Type` to dynamically dispatch rendering
- Prefer composition and reuse of renderer logic (e.g. block-with-children)

---

## 🧪 Testing

You can test block rendering by feeding JSON directly and comparing output.

```go
html, err := converter.Convert([]byte(myJSON))
if html != expectedHTML {
    t.Errorf("render mismatch")
}
```

---

## 📁 Example: Mixed Nested Blocks

Input:

```json
{
  "time": 0,
  "version": "2.27.0",
  "blocks": [
    {
      "type": "quote",
      "data": {
        "items": [
          {
            "type": "paragraph",
            "data": { "text": "Hello inside quote" }
          }
        ]
      }
    }
  ]
}
```

Rendered:

```html
<blockquote><p>Hello inside quote</p></blockquote>
```

---

## 📌 Limitations

- You must register all block types used in the document
- No default escaping or sanitization
- Rendering logic is fully manual (by design)

---

## 📦 Roadmap Ideas (Post-v1)

- Plugin-based renderer loader
- Optional escaping middleware

---

## 👤 License

MIT – free to use, modify, and distribute.
