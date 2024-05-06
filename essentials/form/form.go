package form

import (
	"github.com/lightningsdk/ui/renderer"
	"golang.org/x/net/html"
)

type Form struct {
	Location string               `yaml:"location"`
	Hidden   map[string]string    `yaml:"hidden"`
	Method   string               `yaml:"method" default:"POST"`
	Contents []renderer.Component `xml:"columns"`
	Class    string               `yaml:"class"`
}

func (f *Form) Node(fr renderer.Frame) (*html.Node, error) {
	n := &html.Node{
		Type: html.ElementNode,
		Data: "form",
		Attr: []html.Attribute{
			{
				Key: "class",
				Val: f.Class,
			},
		},
	}

	for k, v := range f.Hidden {
		hn := &html.Node{
			Type: html.ElementNode,
			Data: "input",
			Attr: []html.Attribute{{
				Key: "type",
				Val: "hidden",
			}, {
				Key: "name",
				Val: k,
			}, {
				Key: "value",
				Val: v,
			}},
		}
		n.AppendChild(hn)
	}

	for _, c := range f.Contents {
		cn, err := c.Node(fr)
		if err != nil {
			return nil, err
		}
		n.AppendChild(cn)
	}

	return n, nil
}
