package requestid

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandString(t *testing.T) {
	const size = 32
	for i := 0; i < 100; i++ {
		s1 := randString(size)
		s2 := randString(size)
		assert.Equal(t, len(s1), size)
		assert.Equal(t, len(s2), size)
		assert.NotEqual(t, s1, s2)
	}
	fmt.Println(randString(size))
}
