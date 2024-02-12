package model

type Metadata struct {
	Keywords    []string `sql:"keywords"`
	Slug        string   `sql:"slug"`
	Title       string   `sql:"title"`
	Description string   `sql:"description"`
}
