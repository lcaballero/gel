package gel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaybe(t *testing.T) {
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
			name: `Default(1, Text("2")) == "nil"`,
			view: func() *Node {
				return Default(1, Text("2")).ToNode()
			},
			wantType:   Textual,
			wantString: "2",
			wantCData:  "2",
		},
		{
			name: `Default(nil, Text("nil")) == "nil"`,
			view: func() *Node {
				return Default("not nil", Text("nil")).ToNode()
			},
			wantType:   Textual,
			wantString: "not nil",
			wantCData:  "not nil",
		},
		{
			name: `Default("nil") == "nil"`,
			view: func() *Node {
				return Default(nil, Text("nil")).ToNode()
			},
			wantType:   Textual,
			wantString: "nil",
			wantCData:  "nil",
		},
		{
			name: `Maybe(1) == ""`,
			view: func() *Node {
				return Maybe(1).ToNode()
			},
			wantType:   Textual,
			wantString: "",
			wantCData:  "",
		},
		{
			name: `Maybe("string") == ""`,
			view: func() *Node {
				return Maybe("what?").ToNode()
			},
			wantType:   Textual,
			wantString: "what?",
			wantCData:  "what?",
		},
		{
			name: `Maybe(nil) == ""`,
			view: func() *Node {
				return Maybe(nil).ToNode()
			},
			wantType:   Textual,
			wantString: "",
			wantCData:  "",
		},
		{
			name: `Maybe(Text("default")) == "default"`,
			view: func() *Node {
				return Maybe(Text("default")).ToNode()
			},
			wantType:   Textual,
			wantString: "default",
			wantCData:  "default",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			node := tc.view()
			assert.Equal(t, tc.wantTag, node.tag)
			assert.Equal(t, tc.wantString, node.String())
			assert.Equal(t, tc.wantKidCount, len(node.children))
			if tc.wantNonNilKids {
				assert.NotNil(t, node.children != nil)
			}
			assert.Equal(t, tc.wantCData, node.cdata)
			assert.Equal(t, tc.wantAttsCount, len(node.attributes))
			assert.Equal(t, tc.wantKey, node.key)
			assert.Equal(t, tc.wantValue, node.value)
			assert.Equal(t, tc.wantType, node.kind)
		})
	}
}
