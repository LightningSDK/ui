package standard

import (
	"context"
	"github.com/lightningsdk/ui/renderer"
	"golang.org/x/net/html"
	"io"
)

type HTML struct {
	Contents string `yaml:"contents" xml:",innerxml"`
	JS       string `yaml:"js"`
}

func (h *HTML) Render(ctx context.Context, w io.Writer, f renderer.Frame) error {
	_, err := w.Write([]byte(h.Contents))
	return err
}

func (h *HTML) Node(f renderer.Frame) (*html.Node, error) {
	return &html.Node{
		Parent:      nil,
		FirstChild:  nil,
		LastChild:   nil,
		PrevSibling: nil,
		NextSibling: nil,
		Type:        html.RawNode,
		DataAtom:    0,
		Data:        h.Contents,
		Namespace:   "",
		Attr:        nil,
	}, nil
}
