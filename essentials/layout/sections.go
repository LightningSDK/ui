package layout

import (
	"github.com/lightningsdk/ui/parser"
	"github.com/lightningsdk/ui/renderer"
	"golang.org/x/net/html"
	"gopkg.in/yaml.v3"
)

type Sections struct {
	Contents   []renderer.Component `xml:"section"`
	Class      string               `yaml:"class"`
	ChildClass string               `yaml:"childClass"`
}

func (s *Sections) UnmarshalYAML(n *yaml.Node) error {
	var err error
	s.Contents, err = parser.ParseRendererList("contents", n)
	return err
}

func (s *Sections) Node() (*html.Node, error) {
	attr := []html.Attribute{}
	if s.Class != "" {
		attr = append(attr, html.Attribute{
			Key: "class",
			Val: s.Class,
		})
	}
	n := &html.Node{
		Type: html.ElementNode,
		Data: "sections",
		Attr: attr,
	}

	cattr := []html.Attribute{}
	if s.ChildClass != "" {
		attr = append(cattr, html.Attribute{
			Key: "class",
			Val: s.ChildClass,
		})
	}

	for _, c := range s.Contents {
		sw := &html.Node{
			Type: html.ElementNode,
			Data: "section",
			Attr: cattr,
		}
		sn, err := c.Node()
		if err != nil {
			return nil, err
		}
		sw.AppendChild(sn)
		n.AppendChild(sw)
	}

	return n, nil
}
