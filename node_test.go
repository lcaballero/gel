package gel

import (
	"testing"

	"bytes"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGel(t *testing.T) {

	Convey(`Class(.) should add the class attribute`, t, func() {
		d := Div.Class("row").Text("Hello, World!")
		So(d.ToNode().String(), ShouldEqual, `<div class="row">Hello, World!</div>`)
	})

	Convey(`Adding an Attribute to AttributeList should add children to the Attribute`, t, func() {
		at := Div.Atts("class", "row", "id", "id-1")().ToNode()
		So(at.Attributes, ShouldHaveLength, 2)
		So(at.String(), ShouldEqual, `<div class="row" id="id-1"></div>`)
	})

	Convey(`Adding an Attribute to AttributeList should add children to the Attribute`, t, func() {
		at := Atts("class", "row", "id", "id-1").ToNode()
		So(at.Children, ShouldHaveLength, 2)
		at.Add(Att("type", "text"))
		So(at.Children, ShouldHaveLength, 3)
		So(at.String(), ShouldEqual, ` class="row" id="id-1" type="text"`)
	})

	Convey(`Adding an AttributeList to an Element should add the lists children to Elements atts`, t, func() {
		at := Atts("class", "row", "id", "id-1").ToNode()
		So(at.Children, ShouldHaveLength, 2)
		So(at.String(), ShouldEqual, ` class="row" id="id-1"`)
	})

	Convey(`Adding an AttributeList to an Element should add the lists children to Elements atts`, t, func() {
		at := Atts("class", "row")
		d := Div(at).ToNode()
		So(d.Attributes, ShouldHaveLength, 1)
		So(d.String(), ShouldEqual, `<div class="row"></div>`)
	})

	Convey(`An attribute list can only have children but no other node fields`, t, func() {
		list := Atts().ToNode()
		So(list.Children, ShouldNotBeNil)
		So(len(list.Children), ShouldEqual, 0)
		So(list.CData, ShouldBeEmpty)
		So(list.Attributes, ShouldBeEmpty)
		So(list.Key, ShouldBeEmpty)
		So(list.Value, ShouldBeEmpty)
		So(list.Tag, ShouldEqual, "")
		So(list.Type, ShouldEqual, AttributeList)
		So(list.String(), ShouldBeEmpty)
	})

	Convey(`A Views list should properly render to html like any fragment`, t, func() {
		v := NewFragment()
		v.Add(
			Div.Text("a"),
			Div.Text("b"),
			Div.Text("c"),
		)
		s := v.ToNode().String()
		So(s, ShouldEqual, `<div>a</div><div>b</div><div>c</div>`)
	})

	Convey(`Adding to a Views list should increase it's length`, t, func() {
		v := NewFragment()
		So(v.Len(), ShouldEqual, 0)
		v.Add(
			Div.Text("a"),
		)
		So(v.Len(), ShouldEqual, 1)
	})

	Convey(`A None node should be an empty Text *Node`, t, func() {
		none := None().ToNode()
		So(none.Attributes, ShouldBeNil)
		So(none.CData, ShouldBeEmpty)
		So(none.Children, ShouldBeEmpty)
		So(none.Key, ShouldBeEmpty)
		So(none.Value, ShouldBeEmpty)
		So(none.Tag, ShouldEqual, "")
		So(none.Type, ShouldEqual, Textual)
		So(none.String(), ShouldBeEmpty)
	})

	Convey(`empty text should NOT be added to a div which is already (especially with indented)`, t, func() {
		fr := Frag(
			Div(
				Div(
					Text("2nd level"),
				),
				Div(
					Text(""), // <= should not be preceded by indention
				),
			),
		)
		html := fr.ToNode().String()
		So(html, ShouldEqual, `<div><div>2nd level</div><div></div></div>`)
	})

	Convey(`A new fragment with text and divs should properly render`, t, func() {
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
		actual := fr.ToNode().String()
		So(actual, ShouldEqual, `<div>Here<div>See</div></div><div>Freedom</div>All`)
	})

	Convey(`Add should silently ignore Atts for non-element Nodes`, t, func() {
		txt := Text("Hello, World!").ToNode()
		txt.Add(Att("id", "my-id"))
		So(txt.Attributes, ShouldBeNil)

		frag := Frag(Att("href", "/")).ToNode()
		So(frag.Attributes, ShouldBeNil)

		atts := Att("src", "favicon.ico").ToNode().Add(Att("id", "id-1")).ToNode()
		So(atts.Attributes, ShouldBeNil)
	})

	Convey(`When using Tag.Text and adding strings to a Node they will be appended as text nodes`, t, func() {
		d := Div.Text("Hello,", " World")
		html := d.ToNode().String()
		So(html, ShouldEqual, "<div>Hello, World</div>")
	})

	Convey(`Adding strings to a Node will append text nodes`, t, func() {
		d := Div.Text("Hello,", " World")
		html := d.ToNode().String()
		So(html, ShouldEqual, "<div>Hello, World</div>")
	})

	Convey(`empty tags shouldn't include whitespace when render with indent`, t, func() {
		buf := bytes.NewBuffer([]byte{})
		Div().ToNode().WriteToIndented(NewIndent().Incr(), buf)
		So(buf.String(), ShouldEqual, "  <div></div>\n")
	})

	Convey(`void tags should not have a closing tag`, t, func() {
		buf := bytes.NewBuffer([]byte{})
		Meta().ToNode().WriteToIndented(NewIndent(), buf)
		So(buf.String(), ShouldEqual, "<meta/>")
	})

	Convey(`div using Add(...) many atts should render as <div class="container" id="id-1">text</div>`, t, func() {
		s := Div(
			Att("class", "container"),
			Att("id", "id-1"),
			Text("text"),
			Div(
				Att("class", "row"),
				Text("Hello, World!"),
			),
		)
		expected := `<div class="container" id="id-1">
  text
  <div class="row">
    Hello, World!
  </div>
</div>`
		buf := bytes.NewBuffer([]byte{})
		s.ToNode().WriteToIndented(NewIndent(), buf)
		actual := buf.String()
		So(actual, ShouldEqual, expected)
	})

	Convey(`A Fragment node should have propery fields set`, t, func() {
		f := Frag().ToNode()
		So(f.String(), ShouldEqual, ``)
		So(f.Type, ShouldEqual, NodeList)
		So(f.Children, ShouldNotBeNil)
		So(f.Attributes, ShouldBeNil)
		So(len(f.Children), ShouldEqual, 0)
		So(f.Key, ShouldEqual, "")
		So(f.Value, ShouldEqual, "")
		So(f.CData, ShouldEqual, "")
	})

	Convey(`An Attribute node should html attribute pairs`, t, func() {
		d := Att("class", "container").ToNode()
		So(d.String(), ShouldEqual, ` class="container"`)
		So(d.Type, ShouldEqual, Attribute)
		So(d.Children, ShouldBeNil)
		So(d.Attributes, ShouldBeNil)
		So(len(d.Children), ShouldEqual, 0)
		So(len(d.Attributes), ShouldEqual, 0)
		So(d.Key, ShouldEqual, "class")
		So(d.Value, ShouldEqual, "container")
		So(d.CData, ShouldEqual, "")
	})

	Convey(`An Element node should only produce html tags`, t, func() {
		d := Div().ToNode()
		So(d.String(), ShouldEqual, "<div></div>")
		So(d.Type, ShouldEqual, Element)
		So(d.Children, ShouldNotBeNil)
		So(d.Attributes, ShouldNotBeNil)
		So(len(d.Children), ShouldEqual, 0)
		So(len(d.Attributes), ShouldEqual, 0)
		So(d.Key, ShouldEqual, "")
		So(d.Value, ShouldEqual, "")
		So(d.CData, ShouldEqual, "")
	})

	Convey(`A Text node should produce only it's given text.`, t, func() {
		s := Text("text").ToNode()
		So(s.String(), ShouldEqual, `text`)
		So(s.CData, ShouldEqual, `text`)
		So(s.Tag, ShouldEqual, "")
		So(s.Children, ShouldBeNil)
		So(s.Attributes, ShouldBeNil)
		So(s.Type, ShouldEqual, Textual)
		So(s.Key, ShouldEqual, "")
		So(s.Value, ShouldEqual, "")
	})

	Convey(`div using Add(...) many atts should render as <div class="container" id="id-1">text</div>`, t, func() {
		s := Div(
			Att("class", "container"),
			Att("id", "id-1"),
			Text("text"),
		).ToNode()
		So(s.String(), ShouldEqual, `<div class="container" id="id-1">text</div>`)
	})

	Convey(`div using Fmt(...) should render as <div>hello, world!</div>`, t, func() {
		s := Div.Fmt("hello, %s", "world!").ToNode().String()
		So(s, ShouldEqual, `<div>hello, world!</div>`)
	})

	Convey(`div using Add(...) should render as <div>text</div>`, t, func() {
		s := Div.Text("text").ToNode().String()
		So(s, ShouldEqual, `<div>text</div>`)
	})

	Convey(`div using New(...) should render as <div>text</div>`, t, func() {
		s := Div(Text("text")).ToNode().String()
		So(s, ShouldEqual, `<div>text</div>`)
	})

	Convey(`div using New(...) should render as <div class="container">text</div>`, t, func() {
		s := Div(
			Att("class", "container"),
			Text("text"),
		).ToNode().String()
		So(s, ShouldEqual, `<div class="container">text</div>`)
	})

	Convey(`div using New(...) should render as <div class="container"></div>`, t, func() {
		s := Div(Att("class", "container")).ToNode().String()
		So(s, ShouldEqual, `<div class="container"></div>`)
	})

	Convey("div using New(...) should render as <div><div></div></div>", t, func() {
		s := Div(Div()).ToNode().String()
		So(s, ShouldEqual, "<div><div></div></div>")
	})

	Convey("div renders as <div><div></div></div>", t, func() {
		s := Div(Div()).ToNode().String()
		So(s, ShouldEqual, "<div><div></div></div>")
	})

	Convey("div renders as <div></div>", t, func() {
		s := Div(Div()).ToNode().String()
		So(s, ShouldEqual, "<div><div></div></div>")
	})

	Convey("div renders as <div></div>", t, func() {
		s := Div().ToNode().String()
		So(s, ShouldEqual, "<div></div>")
	})
}
