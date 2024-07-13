package copytext

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCopy(t *testing.T) {
	srcFile, err := os.CreateTemp("", "src")
	assert.NoError(t, err)
	defer func() {
		err := os.Remove(srcFile.Name())
		assert.NoError(t, err)
	}()

	_, err = srcFile.WriteString("hello, world!")
	assert.NoError(t, err)
	err = srcFile.Close()
	assert.NoError(t, err)

	destFile, err := os.CreateTemp("", "dest")
	assert.NoError(t, err)
	defer func() {
		err := os.Remove(destFile.Name())
		assert.NoError(t, err)
	}()

	err = Copy(srcFile.Name(), destFile.Name(), 0, 0)
	assert.NoError(t, err)

	destData, err := os.ReadFile(destFile.Name())
	assert.NoError(t, err)
	assert.Equal(t, "hello, world!", string(destData))

	destFile2, err := os.CreateTemp("", "dest2")
	assert.NoError(t, err)
	defer func() {
		err := os.Remove(destFile2.Name())
		if err != nil {

		}
	}()

	err = Copy(destFile.Name(), destFile2.Name(), 7, 0)
	assert.NoError(t, err)

	destData2, err := os.ReadFile(destFile2.Name())
	assert.NoError(t, err)
	assert.Equal(t, "hello, world!", string(destData2))
}
