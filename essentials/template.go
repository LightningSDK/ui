package essentials

import (
	"github.com/lightningsdk/ui/parser"
	"github.com/lightningsdk/ui/renderer"
	"golang.org/x/net/html"
	"gopkg.in/yaml.v3"
)

type Template struct {
	Contents []renderer.Component `xml:"contents"`
	Name     string               `yaml:"name"`
}

func (s *Template) UnmarshalYAML(n *yaml.Node) error {
	var err error
	s.Contents, err = parser.ParseRendererList("contents", n)
	return err
}

func (s *Template) Node(f renderer.Frame) (*html.Node, error) {
	n := &html.Node{
		Type: html.ElementNode,
	}

	for _, c := range s.Contents {
		cn, err := c.Node(f)
		if err != nil {
			return nil, err
		}
		n.AppendChild(cn)
	}

	return n, nil
}
