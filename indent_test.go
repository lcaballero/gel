package gel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndent_Level(t *testing.T) {
	cases := []struct {
		indent func() Indent
		name   string
		want   int
	}{
		{
			name: "Decrementing an indent should step back the level by it's Inc",
			indent: func() Indent {
				indent := NewIndent()
				indent.Increment = 3
				return indent.Inc().Inc().Dec()
			},
			want: 3,
		},
		{
			name: "",
			indent: func() Indent {
				indent := NewIndent()
				indent.Increment = 3
				return indent.Inc().Inc()
			},
			want: 6,
		},
		{
			name: "NewIndent should start with the defaults",
			indent: func() Indent {
				indent := NewIndent()
				return indent
				//  So(indent.Tab, ShouldEqual, DefaultTab)
				//  So(indent.Level, ShouldEqual, DefaultLevel)
				//  So(indent.Inc, ShouldEqual, DefaultIncrement)
				// })
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			indent := tc.indent()
			assert.Equal(t, tc.want, indent.Level,
				"expected: %v, got: %v", tc.want, indent.Level)
		})
	}
}

func TestIndent_Prefix(t *testing.T) {
	cases := []struct {
		indent   func() Indent
		name     string
		want     string
		hasPanic bool
	}{
		{
			name: "Indent level of 3, and 2 spaces produce '      ' 6 spaces for indent ",
			indent: func() Indent {
				indent := Indent{
					Level:     3,
					Tab:       "  ",
					Increment: 1,
				}
				return indent
			},
			want: "      ",
		},
		{
			name: "Decrementing below 0 should panic.",
			indent: func() Indent {
				return NewIndent().Dec()
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				pain := recover()
				err, ok := pain.(error)
				if tc.hasPanic {
					assert.True(t, ok)
					assert.NotNil(t, err)
				}
			}()
			indent := tc.indent()
			assert.Equal(t, tc.want, indent.String(),
				"expected: %v, got: %v", tc.want, indent.String())
		})
	}
}
