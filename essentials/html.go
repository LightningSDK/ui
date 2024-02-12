package essentials

type Include struct {
	// these will be the default elements that will be in each of the renderable elements
	SASS          string
	SASSContainer string
	SASSFile      string
	JSFile        string
	Class         string `yaml:"class"`
	ChildClass    string `yaml:"child_class"`
}
