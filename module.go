package ui

import (
	"github.com/lightningsdk/core"
	"github.com/lightningsdk/ui/essentials"
	"github.com/lightningsdk/ui/essentials/layout"
	"github.com/lightningsdk/ui/parser"
	"github.com/lightningsdk/ui/renderer"
	"github.com/lightningsdk/ui/standard"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
)

const ModuleName = "github.com/lightningsdk/ui"

type Module struct {
	core.DefaultModule
	*Config
	renderer.Service
}

type Config struct {
	Framework string `yaml:"framework" default:"bootstrap"`
}

func NewModule(app *core.App) core.Module {
	initParsers()
	templates := getTemplates()
	return &Module{
		Service: renderer.New(templates),
	}
}

func initParsers() {
	// these will parse from html to specific components
	for k, v := range map[string]any{
		"html":     standard.HTML{},
		"sections": layout.Sections{},
		"template": essentials.Template{},
	} {
		// TODO: handle this error
		_ = parser.AddType(k, reflect.TypeOf(v))
	}
}

func getTemplates() map[string]renderer.Component {
	c := map[string]renderer.Component{}
	r, _ := regexp.Compile(".*ya?ml$")
	for _, p := range []string{"./content/module_defaults", "./content/templates"} {
		_ = filepath.WalkDir(p,
			func(path string, d os.DirEntry, err error) error {
				if err != nil {
					return err
				}
				if !r.Match([]byte(path)) {
					return nil
				}
				template := parser.RenderParser{}
				data, err := os.ReadFile(path)
				name, err := filepath.Rel(p, path)
				err = yaml.Unmarshal(data, &template)
				c[name] = template.Renderer
				return nil
			})
	}
	// TODO: return the error if it fails
	return c
}

func (m *Module) GetCommands() map[string]core.Command {
	return map[string]core.Command{
		"build": {
			Function: build,
			Help:     "builds all the tempaltes, css and js files",
		},
	}
}

func (m *Module) GetEmptyConfig() any {
	return &Config{}
}
func (m *Module) SetConfig(cfg any) {
	m.Config = cfg.(*Config)
}

func GetRenderer(app *core.App) renderer.Service {
	// TODO: if this is nil, it needs to report an error that the UI plugin needs to be added to the config
	return app.Modules[ModuleName].(*Module).Service
}
