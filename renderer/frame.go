package renderer

import "context"

const (
	WrapperTypePage = iota
	WrapperTypeWidget
	WrapperTypeEmail
)

type Frame interface {
	//AddJSFile(path string)
	//AddCSSFile(path string)
	//AddJSParam(param string, value any)
}

type FrameRender struct {
}

func render(ctx context.Context, template any, wrapper int) (string, error) {
	// the template should have a type and each type should have its own renderer.
	// starting at the root level
	// there can be a page renderer added at the top level, that will contain the metadata
	// if there are sub elements, those should each have their own renderer and that renderer should be called
	// they can optionally be called with a cache wrapper
	// if a function does not implement cache, or has a shorter cache, it should relay that information upstream
	return "", nil
}
