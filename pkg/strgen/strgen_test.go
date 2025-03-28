package strgen_test

import (
	"fmt"
	"niksmo/online-store/pkg/strgen"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStrGen(t *testing.T) {
	t.Run("Positive len", func(t *testing.T) {
		expectedLen := 10
		str := strgen.Len(expectedLen)
		require.Len(t, str, expectedLen)

		pattern := fmt.Sprintf("^[A-Z]{%d}$", expectedLen)
		assert.True(t, regexp.MustCompile(pattern).Match([]byte(str)))
	})

	t.Run("Negative cases", func(t *testing.T) {
		lengthSlice := []int{0, -1}
		for _, l := range lengthSlice {
			t.Run(fmt.Sprintf("Len=(%d)", l), func(t *testing.T) {
				str := strgen.Len(l)
				expectedLen := 1
				assert.Len(t, str, expectedLen)
			})
		}
	})
}
