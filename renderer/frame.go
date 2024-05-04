package renderer

import (
	"bytes"
	"context"
	"golang.org/x/net/html"
)

const (
	WrapperTypePage = iota
	WrapperTypeWidget
	WrapperTypeEmail
)

type Frame interface {
	AddJSFile(path string)
	AddCSSFile(path string)
	AddJSParam(param string, value any)
}

type PageFrame struct {
	BrowserFrame
}
type WidgetFrame struct {
	BrowserFrame
}
type EmailFrame struct {
}
type BrowserFrame struct {
	JSFiles  map[string]string
	CSSFiles map[string]string
	JSParams map[string]any
}

func (b *BrowserFrame) AddJSFile(path string) {
	b.JSFiles[path] = path
}
func (b *BrowserFrame) AddCSSFile(path string) {
	b.CSSFiles[path] = path
}
func (b *BrowserFrame) AddJSParam(name string, value any) {
	b.JSParams[name] = value
}

type FrameRender struct {
	JSFiles  map[string]string
	CSSFiles map[string]string
	JSParams map[string]any
}

func render(ctx context.Context, template Component) (*bytes.Buffer, error) {
	// the template should have a type and each type should have its own renderer.
	// starting at the root level
	// there can be a page renderer added at the top level, that will contain the metadata
	// if there are sub elements, those should each have their own renderer and that renderer should be called
	// they can optionally be called with a cache wrapper
	// if a function does not implement cache, or has a shorter cache, it should relay that information upstream

	b := []byte{}
	w := bytes.NewBuffer(b)
	f := &PageFrame{}

	// forget the wrapper for now, lets just render the template
	n, err := template.Node(f)
	if err != nil {
		return nil, err
	}
	err = html.Render(w, n)
	if err != nil {
		return nil, err
	}

	return w, nil
}

func (f *FrameRender) AddJSFile(path string) {
	f.JSFiles[path] = path
}
func (f *FrameRender) AddCSSFile(path string) {
	f.CSSFiles[path] = path
}
func (f *FrameRender) AddJSParam(name string, value any) {
	f.JSParams[name] = value
}
