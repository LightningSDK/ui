package ui

import "github.com/lightningsdk/core"

type Module struct {
	core.DefaultModule
	Config
}

type Config struct {
	Framework string `yaml:"framework" default:"bootstrap"`
}

func NewModule(app *core.App) core.Module {
	return &Module{}
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
	return Config{}
}
func (m *Module) SetConfig(cfg any) {
	m.Config = cfg.(Config)
}
