package form

import (
	"github.com/lightningsdk/ui/essentials"
	"github.com/lightningsdk/ui/renderer"
	"golang.org/x/net/html"
)

type Div struct {
	essentials.Include
	Location string               `yaml:"location"`
	Contents []renderer.Component `yaml:"contents"`
}

func (d *Div) Node(f renderer.Frame) (*html.Node, error) {
	f.AddJSFile("")
	n := &html.Node{
		Type: html.ElementNode,
		Data: "div",
		Attr: []html.Attribute{
			{
				Key: "class",
				Val: d.Class,
			},
		},
	}

	for _, c := range d.Contents {
		cn, err := c.Node(f)
		if err != nil {
			return nil, err
		}
		n.AppendChild(cn)
	}

	return n, nil
}
