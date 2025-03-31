package serializer_test

import (
	"niksmo/online-store/pkg/serializer"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInt(t *testing.T) {
	t.Run("regular integer", func(t *testing.T) {
		expected := []byte{0, 0, 0, 0, 0, 0, 0, 1}
		n := 1
		actual := serializer.Int(n)
		assert.Equal(t, expected, actual)
	})
}
