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
	expectedObjects := &Node{
		Element: "sections",
		Attrs: map[string]string{
			"name":  "template",
			"class": "outer",
		},
		Contents: []renderer.Component{
			&Node{
				Element: "section",
				Attrs: map[string]string{
					"name":  "header",
					"class": "section one",
					"id":    "section_one",
				},
				Contents: []renderer.Component{
					&HTML{
						Contents: "<div>contents</div>",
					},
				},
			},
			&Node{
				Element: "section",
				Attrs: map[string]string{
					"name":  "body",
					"class": "section two",
					"id":    "section_two",
				},
				Contents: []renderer.Component{
					&HTML{
						Contents: "{{ contents2 }}",
					},
				},
			},
		},
	}

	expectedMarshaledYAML := `type: sections
class: "outer"
name: "template"
contents:
  - type: section
    name: header
    class: "section one"
    id: section_one
    contents:
      - type: html
        contents: <div>contents</div>
  - type: section
    name: body
    class: "section two"
    id: section_two
    contents:
    - type: html
      contents: "{{ contents2 }}"
`

	s := &Node{}
	err = yaml.Unmarshal([]byte(ey), s)
	assert.NoError(t, err)
	assert.Equal(t, expectedObjects, s)

	my, err := yaml.Marshal(expectedObjects)
	assert.NoError(t, err)
	assert.YAMLEq(t, expectedMarshaledYAML, string(my))

	ex := `<sections name="template" class="outer"><section class="section one" id="section_one" name="header"><div>contents</div></section><section name="body" class="section two" id="section_two">{{ contents2 }}</section></sections>`
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
