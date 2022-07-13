package gel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNode_WriteToIndented(t *testing.T) {
	// name: `empty tags shouldn't include whitespace when render with indent`,
	// view: func() *Node {
	// }
	//      buf := bytes.NewBuffer([]byte{})
	//      Div().ToNode().WriteToIndented(NewIndent().Inc(), buf)
	//      So(buf.String(), ShouldEqual, "  <div></div>\n")
	//  })

	//  Convey(`Add should silently ignore Atts for non-element Nodes`, t, func() {
	//      txt := Text("Hello, World!").ToNode()
	//      txt.Add(Att("id", "my-id"))
	//      So(txt.Attributes, ShouldBeNil)
	//      frag := Frag(Att("href", "/")).ToNode()
	//      So(frag.Attributes, ShouldBeNil)
	//      atts := Att("src", "favicon.ico").ToNode().Add(Att("id", "id-1")).ToNode()
	//      So(atts.Attributes, ShouldBeNil)
	//  })

	//  Convey(`void tags should not have a closing tag`, t, func() {
	//      buf := bytes.NewBuffer([]byte{})
	//      Meta().ToNode().WriteToIndented(NewIndent(), buf)
	//      So(buf.String(), ShouldEqual, "<meta/>")
	//  })

	//  Convey(`div using Add(...) many atts should render as <div class="container" id="id-1">text</div>`, t, func() {
	//      s := Div(
	//          Att("class", "container"),
	//          Att("id", "id-1"),
	//          Text("text"),
	//          Div(
	//              Att("class", "row"),
	//              Text("Hello, World!"),
	//          ),
	//      )
	//      expected := `<div class="container" id="id-1">
	//   text
	//   <div class="row">
	//     Hello, World!
	//   </div>
	// </div>`
	//      buf := bytes.NewBuffer([]byte{})
	//      s.ToNode().WriteToIndented(NewIndent(), buf)
	//      actual := buf.String()
	//      So(actual, ShouldEqual, expected)
	//  })
}

func TestNode(t *testing.T) {
	cases := []struct {
		view           func() *Node
		name           string
		wantTag        string
		wantString     string
		wantKidCount   int
		wantNonNilKids bool
		wantCData      string
		wantAttsCount  int
		wantKey        string
		wantValue      string
		wantType       Type
	}{
		{
			name: `An attribute list can only have children but no other node fields`,
			view: func() *Node {
				// So(list.Type, ShouldEqual, AttributeList)
				return Atts().ToNode()
			},
			wantType: AttributeList,
		},
		{
			name: `A None node should be an empty Text *Node`,
			view: func() *Node {
				return None().ToNode()
			},
			wantType: Textual,
		},
		{
			name: `A Fragment node should have propery fields set`,
			view: func() *Node {
				return Frag().ToNode()
			},
			wantType: NodeList,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			node := tc.view()
			assert.Equal(t, tc.wantTag, node.Tag)
			assert.Equal(t, tc.wantString, node.String())
			assert.Equal(t, tc.wantKidCount, len(node.Children))
			if tc.wantNonNilKids {
				assert.NotNil(t, node.Children != nil)
			}
			assert.Equal(t, tc.wantCData, node.CData)
			assert.Equal(t, tc.wantAttsCount, len(node.Attributes))
			assert.Equal(t, tc.wantKey, node.Key)
			assert.Equal(t, tc.wantValue, node.Value)
			assert.Equal(t, tc.wantType, node.Type)
		})
	}
}

func TestFragment(t *testing.T) {
	cases := []struct {
		view    func() *Fragment
		name    string
		wantLen int
	}{
		{
			name:    `Adding to a Views list should increase it's length`,
			wantLen: 0,
			view: func() *Fragment {
				return NewFragment()
			},
		},
		{
			view: func() *Fragment {
				v := NewFragment()
				v.Add(Div.Text("a"))
				return v
			},
			wantLen: 1,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			frag := tc.view()
			assert.Equal(t, tc.wantLen, frag.Len())
		})
	}
}

func TestGel(t *testing.T) {
	cases := []struct {
		view          func() *Node
		name          string
		wantHtml      string
		wantAttsCount int
	}{
		{
			name: `Class(.) should add the class attribute`,
			view: func() *Node {
				return Div.Class("row").Text("Hello, World!").ToNode()
			},
			wantHtml:      `<div class="row">Hello, World!</div>`,
			wantAttsCount: 1,
		},
		{
			name: `Adding an Attribute to AttributeList should add children to the Attribute`,
			view: func() *Node {
				return Div.Atts("class", "row", "id", "id-1")().ToNode()
			},
			wantHtml:      `<div class="row" id="id-1"></div>`,
			wantAttsCount: 2,
		},
		{
			name: `A Views list should properly render to html like any fragment`,
			view: func() *Node {
				v := NewFragment()
				v.Add(
					Div.Text("a"),
					Div.Text("b"),
					Div.Text("c"),
				)
				return v.ToNode()
			},
			wantHtml: `<div>a</div><div>b</div><div>c</div>`,
		},
		{
			name: `Adding an AttributeList to an Element should add the lists children to Elements atts`,
			view: func() *Node {
				at := Atts("class", "row")
				d := Div(at).ToNode()
				return d
			},
			wantAttsCount: 1,
			wantHtml:      `<div class="row"></div>`,
		},
		{
			name: `div using Add(...) many atts should render as <div class="container" id="id-1">text</div>`,
			view: func() *Node {
				return Div(
					Att("class", "container"),
					Att("id", "id-1"),
					Text("text"),
				).ToNode()
			},
			wantHtml:      `<div class="container" id="id-1">text</div>`,
			wantAttsCount: 2,
		},
		{
			name: `div using Fmt(...) should render as <div>hello, world!</div>`,
			view: func() *Node {
				return Div.Fmt("hello, %s", "world!").ToNode()
			},
			wantHtml:      `<div>hello, world!</div>`,
			wantAttsCount: 0,
		},
		{
			name: `div using Add(...) should render as <div>text</div>`,
			view: func() *Node {
				return Div.Text("text").ToNode()
			},
			wantHtml: `<div>text</div>`,
		},
		{
			name: `div using New(...) should render as <div>text</div>`,
			view: func() *Node {
				return Div(Text("text")).ToNode()
			},
			wantHtml: `<div>text</div>`,
		},
		{
			name: `div using New(...) should render as <div class="container">text</div>`,
			view: func() *Node {
				return Div(
					Att("class", "container"),
					Text("text"),
				).ToNode()
			},
			wantHtml:      `<div class="container">text</div>`,
			wantAttsCount: 1,
		},
		{
			name: `div using New(...) should render as <div class="container"></div>`,
			view: func() *Node {
				return Div(Att("class", "container")).ToNode()
			},
			wantHtml:      `<div class="container"></div>`,
			wantAttsCount: 1,
		},
		{
			name: "div renders as <div><div></div></div>",
			view: func() *Node {
				return Div(Div()).ToNode()
			},
			wantHtml: "<div><div></div></div>",
		},
		{
			name: "div renders as <div></div>",
			view: func() *Node {
				return Div().ToNode()
			},
			wantHtml: "<div></div>",
		},
		{
			name: `empty text should NOT be added to a div which is already (especially with indented)`,
			view: func() *Node {
				fr := Frag(
					Div(
						Div(
							Text("2nd level"),
						),
						Div(
							Text(""),
						),
					),
				)
				return fr.ToNode()
			},
			wantHtml: `<div><div>2nd level</div><div></div></div>`,
		},
		{
			name: `A new fragment with text and divs should properly render`,
			view: func() *Node {
				fr := Frag(
					Div(
						Text("Here"),
						Div.Text("See"),
					),
					Div(
						Text("Freedom"),
					),
					Text("All"),
				)
				return fr.ToNode()
			},
			wantHtml: `<div>Here<div>See</div></div><div>Freedom</div>All`,
		},
		{
			name: `When using Tag.Text and adding strings to a Node they will be appended as text nodes`,
			view: func() *Node {
				d := Div.Text("Hello,", " World")
				return d.ToNode()
			},
			wantHtml: "<div>Hello, World</div>",
		},
		{
			name: `Adding strings to a Node will append text nodes`,
			view: func() *Node {
				d := Div.Text("Hello,", " World")
				return d.ToNode()
			},
			wantHtml: "<div>Hello, World</div>",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			node := tc.view()
			got := node.String()
			assert.Equal(t, tc.wantHtml, got,
				"expected: %v, got: %v", tc.wantHtml, got)
			assert.Equal(t, tc.wantAttsCount, len(node.Attributes))
		})
	}

	//  Convey(`An Attribute node should html attribute pairs`, t, func() {
	//      d := Att("class", "container").ToNode()
	//      So(d.String(), ShouldEqual, ` class="container"`)
	//      So(d.Type, ShouldEqual, Attribute)
	//      So(d.Children, ShouldBeNil)
	//      So(d.Attributes, ShouldBeNil)
	//      So(len(d.Children), ShouldEqual, 0)
	//      So(len(d.Attributes), ShouldEqual, 0)
	//      So(d.Key, ShouldEqual, "class")
	//      So(d.Value, ShouldEqual, "container")
	//      So(d.CData, ShouldEqual, "")
	//  })

	//  Convey(`An Element node should only produce html tags`, t, func() {
	//      d := Div().ToNode()
	//      So(d.String(), ShouldEqual, "<div></div>")
	//      So(d.Type, ShouldEqual, Element)
	//      So(d.Children, ShouldNotBeNil)
	//      So(d.Attributes, ShouldNotBeNil)
	//      So(len(d.Children), ShouldEqual, 0)
	//      So(len(d.Attributes), ShouldEqual, 0)
	//      So(d.Key, ShouldEqual, "")
	//      So(d.Value, ShouldEqual, "")
	//      So(d.CData, ShouldEqual, "")
	//  })

	//  Convey(`A Text node should produce only it's given text.`, t, func() {
	//      s := Text("text").ToNode()
	//      So(s.String(), ShouldEqual, `text`)
	//      So(s.CData, ShouldEqual, `text`)
	//      So(s.Tag, ShouldEqual, "")
	//      So(s.Children, ShouldBeNil)
	//      So(s.Attributes, ShouldBeNil)
	//      So(s.Type, ShouldEqual, Textual)
	//      So(s.Key, ShouldEqual, "")
	//      So(s.Value, ShouldEqual, "")
	//  })
}

func TestAtts(t *testing.T) {
	//  Convey(`Adding an Attribute to AttributeList should add children to the Attribute`, t, func() {
	//      at := Atts("class", "row", "id", "id-1").ToNode()
	//      So(at.Children, ShouldHaveLength, 2)
	//      at.Add(Att("type", "text"))
	//      So(at.Children, ShouldHaveLength, 3)
	//      So(at.String(), ShouldEqual, ` class="row" id="id-1" type="text"`)
	//  })

	//  Convey(`Adding an AttributeList to an Element should add the lists children to Elements atts`, t, func() {
	//      at := Atts("class", "row", "id", "id-1").ToNode()
	//      So(at.Children, ShouldHaveLength, 2)
	//      So(at.String(), ShouldEqual, ` class="row" id="id-1"`)
	//  })
}
