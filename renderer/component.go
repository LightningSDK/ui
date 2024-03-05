package renderer

import (
	"golang.org/x/net/html"
)

type Component interface {
	Node(f Frame) (*html.Node, error)
}
