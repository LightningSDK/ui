package layout

import (
	"github.com/lightningsdk/ui/parser"
	"github.com/lightningsdk/ui/renderer"
	"github.com/lightningsdk/ui/standard"
	"github.com/stretchr/testify/assert"
	"github.com/tdewolff/parse/buffer"
	"golang.org/x/net/html"
	"gopkg.in/yaml.v3"
	"reflect"
	"testing"
)

func Test_Sections(t *testing.T) {
	err := parser.AddType("html", reflect.TypeOf(standard.HTML{}))
	assert.NoError(t, err)
	err = parser.AddType("sections", reflect.TypeOf(Node{}))
	assert.NoError(t, err)
	err = parser.AddType("section", reflect.TypeOf(Node{}))
	assert.NoError(t, err)

	ey := `
type: sections
class: "outer"
childClass: "inner"
name: "template"
contents:
  - type: section
    name: header
    class: "section one"
    id: section_one
    contents:
      type: html
      contents: <div>contents</div>
  - type: html
    name: body
    class: "section two"
    id: section_two
    contents: "{{ contents2 }}"
`
	s := &Node{}
	err = yaml.Unmarshal([]byte(ey), s)
	assert.NoError(t, err)
	assert.Equal(t, &Node{
		Element: "sections",
		Name:    "template",
		Class:   "outer",
		Contents: []renderer.Component{
			&Node{
				Element: "section",
				Name:    "header",
				Class:   "section one",
				ID:      "section_one",
				Contents: []renderer.Component{
					&HTML{
						Contents: "<div>contents</div>",
					},
				},
			},
			&Node{
				Element: "section",
				Name:    "body",
				Class:   "section two",
				ID:      "section_two",
				Contents: []renderer.Component{
					&HTML{
						Contents: "{{ contents2 }}",
					},
				},
			},
		},
	}, s)

	ex := `<section>header</section><section>body</section>`
	b := []byte{}
	w := buffer.NewWriter(b)
	//err = html.Render(w, s.Node())
	f := &renderer.FrameRender{}
	//err = s.Render(context.Background(), w, f)
	n, err := s.Node(f)
	assert.NoError(t, err)
	err = html.Render(w, n)
	assert.NoError(t, err)
	assert.Equal(t, ex, string(w.Bytes()))

	ay, err := yaml.Marshal(s)
	assert.NoError(t, err)
	assert.Equal(t, ey, string(ay))
}
