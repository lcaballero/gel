package gel

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"bytes"
)

func TestGel(t *testing.T) {

	Convey(`empty tags shouldn't include whitespace when render with indent`, t, func() {
		buf := bytes.NewBuffer([]byte{})
		Div.New().WriteToIndented(NewIndent().Incr(), buf)
		So(buf.String(), ShouldEqual, "  <div></div>\n")
	})

	Convey(`void tags should not have a closing tag`, t, func() {
		buf := bytes.NewBuffer([]byte{})
		Meta.New().WriteToIndented(NewIndent(), buf)
		So(buf.String(), ShouldEqual, "<meta/>")
	})

	Convey(`div using Add(...) many atts should render as <div class="container" id="id-1">text</div>`, t, func() {
		s := Div.Add(
			Att("class", "container"),
			Att("id", "id-1"),
			Text("text"),
		).Add(
			Div.New(
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
		s.WriteToIndented(NewIndent(), buf)
		actual := buf.String()
		So(actual, ShouldEqual, expected)
	})

	Convey(`An AAttribute node should html attribute pairs`, t, func() {
		d := Att("class", "container")
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
		d := Div.New()
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
		s := Text("text")
		So(s.String(), ShouldEqual, `text`)
		So(s.CData, ShouldEqual, `text`)
		So(s.Tag, ShouldEqual, 0)
		So(s.Children, ShouldBeNil)
		So(s.Atts, ShouldBeNil)
		So(s.Type, ShouldEqual, Textual)
		So(s.Key, ShouldEqual, "")
		So(s.Value, ShouldEqual, "")
	})

	Convey(`div using Add(...) many atts should render as <div class="container" id="id-1">text</div>`, t, func() {
		s := Div.Add(
			Att("class", "container"),
			Att("id", "id-1"),
			Text("text"),
		)
		So(s.String(), ShouldEqual, `<div class="container" id="id-1">text</div>`)
	})

	Convey(`div using Add(...) should render as <div>text</div>`, t, func() {
		s := Div.Add(Text("text")).String()
		So(s, ShouldEqual, `<div>text</div>`)
	})

	Convey(`div using New(...) should render as <div>text</div>`, t, func() {
		s := Div.New().Add(Text("text")).String()
		So(s, ShouldEqual, `<div>text</div>`)
	})

	Convey(`div using New(...) should render as <div class="container">text</div>`, t, func() {
		s := Div.New(Att("class", "container")).Add(Text("text")).String()
		So(s, ShouldEqual, `<div class="container">text</div>`)
	})

	Convey(`div using New(...) should render as <div class="container"></div>`, t, func() {
		s := Div.New(Att("class", "container")).String()
		So(s, ShouldEqual, `<div class="container"></div>`)
	})

	Convey("div using New(...) should render as <div><div></div></div>", t, func() {
		s := Div.New(Div.New()).String()
		So(s, ShouldEqual, "<div><div></div></div>")
	})

	Convey("div renders as <div><div></div></div>", t, func() {
		s := Div.New().Add(Div.New()).String()
		So(s, ShouldEqual, "<div><div></div></div>")
	})

	Convey("div renders as <div></div>", t, func() {
		s := Div.New().Add(Div.New()).String()
		So(s, ShouldEqual, "<div><div></div></div>")
	})

	Convey("div renders as <div></div>", t, func() {
		s := Div.New().String()
		So(s, ShouldEqual, "<div></div>")
	})
}
