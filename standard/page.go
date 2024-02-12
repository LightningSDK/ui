package standard

import (
	"github.com/lightningsdk/ui/parser"
	"github.com/lightningsdk/ui/renderer"
	"gopkg.in/yaml.v3"
)

// todo: removoe this, page is a wrapper not a render component
type Page struct {
	PageFields
	Contents renderer.Component // this does not have type so it's not converted automatically
}

type PageFields struct {
	Name string `yaml:"name"`
	// include js queue, css queue, etc
}

func (p *Page) UnmarshalYAML(n *yaml.Node) error {
	err := n.Decode(&p.PageFields)
	if err != nil {
		return err
	}

	p.Contents, err = parser.ParseRenderer("contents", n)

	return err
}

//func (p *Page) Render() (string, error) {
//	return p.Contents.Render(p)
//}

func (p *Page) AddJSFile(path string) {

}
func (p *Page) AddCSSFile(path string) {

}
func (p *Page) AddSASSBlock(path string) {

}
