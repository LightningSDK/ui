package renderer

import (
	"bytes"
	"context"
	"errors"
)

type ServiceImpl struct {
	Templates map[string]Component
}
type Service interface {
	RenderTemplate(ctx context.Context, template string) (*bytes.Buffer, error)
}

func New(t map[string]Component) Service {
	return &ServiceImpl{
		Templates: t,
	}
}

func (s *ServiceImpl) LoadAll() error {
	// this should be called when the server starts
	// it should first start by looking through all the modules and loading any default templates
	// then it should look in the application config and replace them with any overrides
	return nil
}

func (s *ServiceImpl) Compile() error {
	// this should build and minify all the sass and JS files
	// this should be run at build time
	return nil
}

func (s *ServiceImpl) RenderTemplate(ctx context.Context, template string) (*bytes.Buffer, error) {
	// load the template
	t, ok := s.Templates[template]
	if !ok {
		return nil, errors.New("template not found")
	}

	return render(ctx, t)
}
