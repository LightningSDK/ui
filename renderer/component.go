package renderer

import (
	"golang.org/x/net/html"
)

type Component interface {
	// the frame is required so that any node can add js and css to the container
	Node(f Frame) (*html.Node, error)
}
