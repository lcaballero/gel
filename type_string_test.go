package gel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Type_String(t *testing.T) {
	cases := []struct {
		kind Type
		want string
	}{
		{kind: Textual, want: "Textual"},
		{kind: Element, want: "Element"},
		{kind: Attribute, want: "Attribute"},
		{kind: NodeList, want: "NodeList"},
		{kind: AttributeList, want: "AttributeList"},
		{kind: Type(42), want: "Type(42)"},
	}
	for _, tc := range cases {
		t.Run(tc.want, func(t *testing.T) {
			assert.Equal(t, tc.want, tc.kind.String())
		})
	}
}
