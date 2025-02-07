package encryption

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const key = "mHxRcbpaeBdHMCmk2T7VfZ3+6Y(IF!82"

func TestEncryptAESGCM(t *testing.T) {
	got, err := EncryptAESGCM([]byte(key), []byte("hello world!"))
	assert.Nil(t, err)
	fmt.Println(got)
}

func TestDecryptAESGCM(t *testing.T) {
	got, err := DecryptAESGCM([]byte(key), "I6puN8ZjMXjWVI/jNdWjwm53cyMYZujoOKJVt1Gu/wf34nMeu9xdWw==")
	assert.Nil(t, err)
	assert.Equal(t, "hello world!", got)
	fmt.Println(got)
}
