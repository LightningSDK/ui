package bootstrap

import (
	"github.com/lightningsdk/ui/essentials"
	"github.com/lightningsdk/ui/parser"
	"github.com/lightningsdk/ui/renderer"
	"golang.org/x/net/html"
	"gopkg.in/yaml.v3"
)

type Columns struct {
	essentials.Include
	Contents []renderer.Component
}

func (c *Columns) UnmarshalYAML(n *yaml.Node) error {
	var err error
	c.Contents, err = parser.ParseRendererList("contents", n)
	return err
}

func (c *Columns) Node() (*html.Node, error) {
	wn := &html.Node{
		Type: html.ElementNode,
		Data: "div",
		Attr: nil,
	}
	class := "sm-5"
	for _, col := range c.Contents {
		cn := &html.Node{
			Type: html.ElementNode,
			Data: "div",
			Attr: []html.Attribute{{
				Key: "class",
				Val: "col " + class,
			}},
		}
		ch, err := col.Node()
		if err != nil {
			return nil, err
		}
		cn.AppendChild(ch)
		wn.AppendChild(cn)
	}
	return wn, nil
}
