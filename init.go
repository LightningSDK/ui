package ui

import (
	"github.com/lightningsdk/ui/essentials/layout"
	"github.com/lightningsdk/ui/standard"
	"reflect"
)

var Components = map[string]reflect.Type{
	"html":     reflect.TypeOf(standard.HTML{}),
	"sections": reflect.TypeOf(layout.Sections{}),
}
