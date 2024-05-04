package parser

import (
	"fmt"
	"github.com/lightningsdk/ui/renderer"
	"gopkg.in/yaml.v3"
	"reflect"
)

type Typer struct {
	Type string `yaml:"type"`
}

type Parser struct {
	Types map[string]reflect.Type
}

func GetContentsInstance(n *yaml.Node) (renderer.Component, error) {
	t := &Typer{}
	err := n.Decode(t)
	if err != nil {
		return nil, err
	}
	// TODO: This should be handled more gracefully
	if t.Type == "" {
		t.Type = "template"
	}
	if p, ok := parsers[t.Type]; ok {
		return reflect.New(p).Interface().(renderer.Component), nil
	}
	return nil, fmt.Errorf("parser type %s not defined", t.Type)
}

//func GetType(unmarshal func(interface{}) error) (string, error) {
//	t := &Typer{}
//	err := unmarshal(t)
//	if err != nil {
//		return "", err
//	}
//	return t.Type, nil
//}
//
//func GetInstance(n string) (renderer.Component, error) {
//	if p, ok := parsers[n]; ok {
//		return reflect.New(p).Interface().(renderer.Component), nil
//	}
//	return nil, fmt.Errorf("parser type %s not defined", n)
//}

var parsers = map[string]reflect.Type{}

func AddType(n string, t reflect.Type) error {
	parsers[n] = t
	return nil
}

type RenderParser struct {
	field    string
	Renderer renderer.Component
}

func (p *RenderParser) UnmarshalYAML(n *yaml.Node) error {
	var err error
	p.Renderer, err = GetContentsInstance(n)
	if err != nil {
		return err
	}
	return nil
}

//type RenderListParser struct {
//	field    string
//	Renderer []renderer.Component
//}

func ParseRenderer(s string, n *yaml.Node) (renderer.Component, error) {
	// find the label node
	var rn *yaml.Node
	for i, nn := range n.Content {
		if nn.Value == s {
			rn = n.Content[i+1]
			break
		}
	}
	if rn == nil {
		return nil, nil
	}
	if rn.Tag != "!!map" {
		return nil, fmt.Errorf("can not parse a render list of type %s in field %s", rn.Tag, s)
	}

	// find the type
	r, err := GetContentsInstance(rn)
	if err != nil {
		return nil, err
	}

	err = rn.Decode(r)
	return r, err
}

func ParseRendererList(s string, n *yaml.Node) ([]renderer.Component, error) {
	// find the label node
	var rn *yaml.Node
	for i, nn := range n.Content {
		if nn.Value == s {
			rn = n.Content[i+1]
			break
		}
	}
	if rn == nil {
		return nil, nil
	}
	if rn.Tag != "!!seq" {
		return nil, fmt.Errorf("can not parse a render list of type %s in field %s", rn.Tag, s)
	}

	// create a list
	rl := []renderer.Component{}

	// find the type
	for _, rni := range rn.Content {
		r, err := GetContentsInstance(rni)
		if err != nil {
			return nil, err
		}

		err = rni.Decode(r)
		if err != nil {
			return nil, err
		}
		rl = append(rl, r)
	}
	return rl, nil
}
