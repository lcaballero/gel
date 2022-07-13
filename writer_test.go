package gel

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockErrorIOWriter int

func (m mockErrorIOWriter) Write(b []byte) (int, error) {
	return 0, fmt.Errorf("error on purpose")
}

func Test_Writer(t *testing.T) {
	var mock mockErrorIOWriter
	w := Writer{mock}
	defer func() {
		pain := recover()
		err, ok := pain.(error)
		assert.True(t, ok)
		assert.NotNil(t, err)
	}()
	w.Write("nope")
}
