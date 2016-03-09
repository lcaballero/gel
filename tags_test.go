package gel

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"fmt"
)

func TestTags(t *testing.T) {

	Convey("voidSet should be initialized", t, func() {
		So(voidSet, ShouldNotBeNil)
	})

	Convey("IsVoidTag should return false for NormalTags", t, func() {
		for _,v := range NormalTags {
			So(v.IsSelfClosing(), ShouldBeFalse)
		}
	})

	Convey("IsVoidTag should return true for VoidTags", t, func() {
		for _,v := range VoidTags {
			if !v.IsSelfClosing() {
				fmt.Println(v.String())
			}
			So(v.IsSelfClosing(), ShouldBeTrue)
		}
	})

	Convey("Number of void + normal tags should equal ALL tags", t, func() {
		count := len(VoidTags) + len(NormalTags)
		So(count, ShouldEqual, len(AllTags))
	})

	Convey("void tags shouldn't be found in non-void tags", t, func() {
		for _, v := range VoidTags {
			for _, n := range NormalTags {
				So(v, ShouldNotEqual, n)
			}
		}
	})
}
