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
	err = parser.AddType("sections", reflect.TypeOf(Sections{}))
	assert.NoError(t, err)

	ey := `type: sections
class: "outer"
childClass: "inner"
contents:
  - type: html
    name: header
    class: "scetion one"
	id: section_one
    contents: <div>contents</div>
  - type: html
    name: body
    contents: "{{ contents2 }}"
`
	s := &Sections{}
	err = yaml.Unmarshal([]byte(ey), s)
	assert.NoError(t, err)
	assert.Equal(t, &Sections{
		Contents: []renderer.Component{
			&standard.HTML{
				Contents: "<div>contents</div>",
				JS:       "",
			},
			&standard.HTML{
				Contents: "{{ contents2 }}",
				JS:       "",
			},
		},
	}, s)

	ex := `<section>header</section><section>body</section>`
	b := []byte{}
	w := buffer.NewWriter(b)
	//err = html.Render(w, s.Node())
	//f := &renderer.FrameRender{}
	//err = s.Render(context.Background(), w, f)
	n, err := s.Node()
	assert.NoError(t, err)
	err = html.Render(w, n)
	assert.NoError(t, err)
	assert.Equal(t, ex, string(w.Bytes()))

	ay, err := yaml.Marshal(s)
	assert.NoError(t, err)
	assert.Equal(t, ey, string(ay))
}
