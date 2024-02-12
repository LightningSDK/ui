package renderer

import (
	"golang.org/x/net/html"
)

type Component interface {
	Node() (*html.Node, error)
}
