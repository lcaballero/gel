package gel

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGel(t *testing.T) {

	Convey(`A Text node should produce only it's given text.`, t, func() {
		s := Text("text")
		So(s.String(), ShouldEqual, `text`)
		So(s.CData, ShouldEqual, `text`)
		So(s.Tag, ShouldEqual, "")
		So(s.Children, ShouldBeNil)
		So(s.Atts, ShouldBeNil)
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
