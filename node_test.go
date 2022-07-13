package gel

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNode_WriteToIndented(t *testing.T) {
	cases := []struct {
		view       func() *Node
		name       string
		wantString string
	}{
		{
			name: `Text node using WriteToIndented`,
			view: func() *Node {
				return Text("here").ToNode()
			},
			wantString: "here\n",
		},
		{
			name: `Div node using WriteToIndented`,
			view: func() *Node {
				return Div.Text("text").ToNode()
			},
			wantString: "<div>\n  text\n</div>",
		},
		{
			name: `Div(Div) node using WriteToIndented`,
			view: func() *Node {
				return Div.Add(Div.Text("text")).ToNode()
			},
			wantString: "<div>\n  <div>\n    text\n  </div>\n</div>",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			node := tc.view()
			buf := bytes.NewBufferString("")
			node.WriteWithIndention(NewIndent(), buf)
			assert.Equal(t, tc.wantString, buf.String())
		})
	}
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
				return Fragment{}.ToNode()
			},
			wantType: NodeList,
		},
		{
			name: `A Text node should produce only it's given text.`,
			view: func() *Node {
				return Text("text").ToNode()
			},
			wantString: "text",
			wantCData:  "text",
			wantType:   Textual,
		},
		{
			name: `An Element node should only produce html tags`,
			view: func() *Node {
				return Div().ToNode()
			},
			wantTag:    "div",
			wantString: `<div></div>`,
			wantType:   Element,
		},
		{
			name: `An Attribute node should produce (html) attribute pairs`,
			view: func() *Node {
				return Att("class", "container").ToNode()
			},
			wantType:   Attribute,
			wantString: ` class="container"`,
			wantKey:    "class",
			wantValue:  "container",
		},
		{
			name: `Adding an Attribute to AttributeList should add children to the Attribute pairs`,
			view: func() *Node {
				return Atts("class", "row", "id", "id-1").ToNode().
					Add(Att("type", "text")).ToNode()
			},
			wantType:     AttributeList,
			wantString:   ` class="row" id="id-1" type="text"`,
			wantKidCount: 3,
		},
		{
			name: `Output self-closing element`,
			view: func() *Node {
				return Img.Atts("href", "/img.png").ToNode()
			},
			wantType:      Element,
			wantString:    `<img href="/img.png"/>`,
			wantKidCount:  0,
			wantAttsCount: 1,
			wantTag:       "img",
		},
		{
			name: `Accept multiple strings for Text node`,
			view: func() *Node {
				return Div().ToNode().Text("a", "b", " ", "x", "z").ToNode()
			},
			wantType:     Element,
			wantString:   `<div>ab xz</div>`,
			wantKidCount: 5,
			wantTag:      "div",
		},
		{
			name: `Custom Tag creation`,
			view: func() *Node {
				MyDiv := E("MyDiv")
				return MyDiv.Text("hello").ToNode()
			},
			wantType:     Element,
			wantString:   `<MyDiv>hello</MyDiv>`,
			wantKidCount: 1,
			wantTag:      "MyDiv",
		},
		{
			name: `Custom Tag with Add(View...)`,
			view: func() *Node {
				MyDiv := E("MyDiv")
				return MyDiv.Add(Div.Text("my-div")).ToNode()
			},
			wantType:     Element,
			wantString:   `<MyDiv><div>my-div</div></MyDiv>`,
			wantKidCount: 1,
			wantTag:      "MyDiv",
		},
		{
			name: `Add an AttributeList to El.Add(Atts(...)...)`,
			view: func() *Node {
				return Atts("id", "my-id").ToNode().Add(
					Atts("class", "container", "onclick", "func()")).ToNode()
			},
			wantType:      AttributeList,
			wantKidCount:  3,
			wantTag:       "",
			wantString:    ` id="my-id" class="container" onclick="func()"`,
			wantAttsCount: 0,
		},
		{
			name: `NodeList.Add(...) adds the children`,
			view: func() *Node {
				f := Fragment{Div(), Div()}
				return f.ToNode().Add(
					Fragment{Div(), Div()},
					Fragment{Div(), Div()},
				).ToNode()
			},
			wantType:      NodeList,
			wantKidCount:  6,
			wantTag:       "",
			wantString:    `<div></div><div></div><div></div><div></div><div></div><div></div>`,
			wantAttsCount: 0,
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
		view    func() Fragment
		name    string
		wantLen int
	}{
		{
			name:    `Adding to a Views list should increase it's length`,
			wantLen: 0,
			view: func() Fragment {
				return Fragment{}
			},
		},
		{
			name: "fragment should have 1 child",
			view: func() Fragment {
				v := Fragment{}
				v = v.Add(Div.Text("a"))
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
				v := Fragment{}
				v = v.Add(
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
				fr := Fragment{
					Div(
						Div(
							Text("2nd level"),
						),
						Div(
							Text(""),
						),
					),
				}
				return fr.ToNode()
			},
			wantHtml: `<div><div>2nd level</div><div></div></div>`,
		},
		{
			name: `A new fragment with text and divs should properly render`,
			view: func() *Node {
				fr := Fragment{
					Div(
						Text("Here"),
						Div.Text("See"),
					),
					Div(
						Text("Freedom"),
					),
					Text("All"),
				}
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
}
