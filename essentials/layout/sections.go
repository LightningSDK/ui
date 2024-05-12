package layout

import (
	"fmt"
	"github.com/lightningsdk/ui/renderer"
	"golang.org/x/net/html"
	"gopkg.in/yaml.v3"
)

type Node struct {
	Element    string
	Contents   []renderer.Component
	Name       string // should be deleted and use attrs
	Class      string // should be deleted and use attrs
	ChildClass string // should be deleted and use attrs
	ID         string // should be deleted and use attrs
	Attrs      map[string]string
}

type HTML struct {
	Contents string
}

// considerations: ul / li might not need to have an entire node created for it, but it would be necessary if the elements were to have specific IDs
var parentWrappers = map[string]string{
	"sections": "section",
	"divs":     "div",
	"columns":  "column",
	"rows":     "row",
}

// this is the top level node, it hijacks the YAML unmarshaler from here forward
// example node:
// n
//
//	Tag: !!map
//	Content: []*yaml.Node // converts to sections Node
//	-  Tag: "!!str" #this is the title node
//	   Value: "type"
//	-  Tag: "!!str" #this is the value node
//	   Value: "sections"
//	-  Tag: !!str
//	   Value: contents
//	-  Tag: !!seq
//	   Content: []*yaml.Node // converts to a section Node // yaml node contents && lightning node contents
//	   - Tag: !!map
//	     Content: []*yaml.Node // converts to html element // yaml node contents
//	     - Tag: !!str
//	       Value: "type"
//	     - Tag: !!str
//	       Value: "html"
//	     - Tag: !!str
//	       Value: "contents"
//	     - Tag: !!str // lightning node contents
//	       Value: "<html>"

// this should always be type map creating a single lightning node which might have children
func (no *Node) UnmarshalYAML(n *yaml.Node) error {
	contents, err := getElementAndContents(n)
	if err != nil {
		return err
	}
	var children *yaml.Node
	for k, v := range contents {
		if _, ok := v.(*yaml.Node); !ok {
			// TODO: improve this
			return fmt.Errorf("not a valid node")
		}
		vn := v.(*yaml.Node)
		switch k {
		case "type":
			no.Element = vn.Value
		case "name":
			no.Name = vn.Value
		case "class":
			no.Class = vn.Value
		case "id":
			no.ID = vn.Value
		case "contents":
			children = vn
		default:
			if no.Attrs == nil {
				no.Attrs = map[string]string{}
			}
			no.Attrs[k] = vn.Value
		}
	}
	expect := ""
	if ex, ok := parentWrappers[no.Element]; ok {
		expect = ex
	}
	no.Contents, err = parseContents(children, expect)
	if err != nil {
		return fmt.Errorf("failed to parse contents: %s", err)
	}

	return nil
}

// example input from above
//   - Tag: !!seq
//     Content: []*yaml.Node
func parseContents(n *yaml.Node, expected string) ([]renderer.Component, error) {
	switch n.Tag {
	case "!!seq":
		// this should be a !!seq of !!map - if not, it's malformed
		lightningNodes := []renderer.Component{}
		// parse as a list of contents
		for _, sub := range n.Content {
			cn, err := getNode(sub, expected)
			if err != nil {
				return nil, err
			}
			lightningNodes = append(lightningNodes, cn)
		}
		return lightningNodes, nil
	case "!!map":
		// parse yaml node
		cn, err := getNode(n, expected)
		if err != nil {
			return nil, err
		}
		return []renderer.Component{cn}, nil
	// parse as single element html
	case "!!str":
		return []renderer.Component{
			&HTML{
				Contents: n.Value,
			},
		}, nil
	default:
		return nil, fmt.Errorf("bad contents type: %s", n.Tag)
	}
}

// this will create a lightning node based on the list of values in a map
func getNode(n *yaml.Node, expected string) (renderer.Component, error) {
	// TODO: this line is duplicated inside Unmarshal - is there a better way?
	contents, err := getElementAndContents(n)
	if err != nil {
		return nil, err
	}

	var cn yaml.Unmarshaler
	element := ""
	if e, ok := contents["type"].(*yaml.Node); ok {
		element = e.Value
	}
	if contents["type"].(*yaml.Node).Value == "html" {
		cn = &HTML{}
	} else {
		cn = &Node{}
	}
	err = cn.UnmarshalYAML(n)
	if err != nil {
		return nil, fmt.Errorf("failed to get content type: %s", err)
	}

	if expected != "" && element != expected {
		// fix assumed children of another type with a wrapper
		wn := &Node{}
		err = wn.UnmarshalYAML(n)
		if err != nil {
			return nil, fmt.Errorf("failed to get content type: %s", err)
		}
		wn.Element = expected
		wn.Contents = []renderer.Component{
			cn.(renderer.Component),
		}
		cn = wn
	}
	return cn.(renderer.Component), nil
}

func getElementAndContents(n *yaml.Node) (map[string]any, error) {
	contents := map[string]any{}
	for i := 0; i < len(n.Content); i += 2 {
		contents[n.Content[i].Value] = n.Content[i+1]
	}

	return contents, nil
}

func (no *Node) Node(f renderer.Frame) (*html.Node, error) {
	attr := []html.Attribute{}
	if no.Class != "" {
		attr = append(attr, html.Attribute{
			Key: "class",
			Val: no.Class,
		})
	}
	n := &html.Node{
		Type: html.ElementNode,
		Data: no.Element,
		Attr: attr,
	}

	for _, c := range no.Contents {
		sn, err := c.Node(f)
		if err != nil {
			return nil, err
		}
		n.AppendChild(sn)
	}

	return n, nil
}

func (h *HTML) UnmarshalYAML(n *yaml.Node) error {
	// todo: add error handling
	contents, err := getElementAndContents(n)
	if err != nil {
		return err
	}
	h.Contents = contents["contents"].(*yaml.Node).Value

	return nil
}
func (h *HTML) Node(f renderer.Frame) (*html.Node, error) {
	//attr := []html.Attribute{}
	//if no.Class != "" {
	//	attr = append(attr, html.Attribute{
	//		Key: "class",
	//		Val: no.Class,
	//	})
	//}
	//n := &html.Node{
	//	Type: html.ElementNode,
	//	Data: no.Element,
	//	Attr: attr,
	//}
	//
	//for _, c := range no.Contents {
	//	sn, err := c.Node(f)
	//	if err != nil {
	//		return nil, err
	//	}
	//	n.AppendChild(sn)
	//}

	return nil, nil
}
