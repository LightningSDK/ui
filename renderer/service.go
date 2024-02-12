package renderer

import (
	"context"
	"errors"
)

type ServiceImpl struct {
	Templates map[string]any
}
type Service interface {
	RenderWithTemplate(ctx context.Context, template string, wrapper int) (string, error)
}

func New() Service {
	return &ServiceImpl{}
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

func (s *ServiceImpl) RenderWithTemplate(ctx context.Context, template string, wrapper int) (string, error) {
	// load the template
	t, ok := s.Templates[template]
	if !ok {
		return "", errors.New("template not found")
	}

	return render(ctx, t, wrapper)
}
