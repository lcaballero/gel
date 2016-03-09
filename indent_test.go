package gel

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIndent(t *testing.T) {

	Convey("Decrementing an indent should step back the level by it's Inc", t, func() {
		indent := NewIndent()
		indent.Inc = 3
		indent = indent.Incr().Incr().Decr()
		So(indent.Level, ShouldEqual, 3)
	})

	Convey("Decrementing an indent should step back the level by it's Inc", t, func() {
		indent := NewIndent()
		indent.Inc = 3
		indent = indent.Incr().Incr().Decr()
		So(indent.Level, ShouldEqual, 3)
	})

	Convey("Incrementing an indent should bump the level by it's Inc", t, func() {
		indent := NewIndent()
		indent.Inc = 3
		indent = indent.Incr().Incr()
		So(indent.Level, ShouldEqual, 6)
	})

	Convey("Indent level of 3, and 2 spaces produce '      ' 6 spaces for indent ", t, func() {
		indent := Indent{Level: 3, Tab: "  ", Inc: 1}
		So(indent.String(), ShouldEqual, "      ")
	})

	Convey("NewIndent should start with the defaults", t, func() {
		indent := NewIndent()
		So(indent.Tab, ShouldEqual, DefaultTab)
		So(indent.Level, ShouldEqual, DefaultLevel)
		So(indent.Inc, ShouldEqual, DefaultIncrement)
	})

	Convey("Decrementing below 0 should panic.", t, func() {
		defer func() {
			pain := recover()
			err, ok := pain.(error)
			So(ok, ShouldBeTrue)
			So(err, ShouldNotBeNil)
		}()
		NewIndent().Decr()
	})
}
