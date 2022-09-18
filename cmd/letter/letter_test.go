package letter

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLetters(t *testing.T) {
	assert.Equal(t, "A", fmt.Sprint(Letter(1)))
}

func TestLetters_FromString(t *testing.T) {
	assert.Equal(t, A, FromString("A"))
}
