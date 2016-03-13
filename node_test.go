package gel

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"bytes"
)

func TestGel(t *testing.T) {

	Convey(`A None node should be an empty Text *Node`, t, func() {
		none := None().ToNode()
		So(none.Atts, ShouldBeNil)
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
		So(txt.Atts, ShouldBeNil)

		frag := Frag(Att("href", "/")).ToNode()
		So(frag.Atts, ShouldBeNil)

		atts := Att("src", "favicon.ico").ToNode().Add(Att("id", "id-1")).ToNode()
		So(atts.Atts, ShouldBeNil)
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
		So(f.Type, ShouldEqual, Fragment)
		So(f.Children, ShouldNotBeNil)
		So(f.Atts, ShouldBeNil)
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
		So(d.Atts, ShouldBeNil)
		So(len(d.Children), ShouldEqual, 0)
		So(len(d.Atts), ShouldEqual, 0)
		So(d.Key, ShouldEqual, "class")
		So(d.Value, ShouldEqual, "container")
		So(d.CData, ShouldEqual, "")
	})

	Convey(`An Element node should only produce html tags`, t, func() {
		d := Div().ToNode()
		So(d.String(), ShouldEqual, "<div></div>")
		So(d.Type, ShouldEqual, Element)
		So(d.Children, ShouldNotBeNil)
		So(d.Atts, ShouldNotBeNil)
		So(len(d.Children), ShouldEqual, 0)
		So(len(d.Atts), ShouldEqual, 0)
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
		So(s.Atts, ShouldBeNil)
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
