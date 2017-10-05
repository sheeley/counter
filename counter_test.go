package counter

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCounter(t *testing.T) {
	c := &String{}
	assert.Len(t, c.Sorted(), 0)
	c.Add("1")
	assert.Len(t, c.Sorted(), 1)
	c.Add("1")
	assert.Len(t, c.Sorted(), 1)
	c.Add("2")
	s := c.Sorted()
	assert.Len(t, s, 2)
	assert.Equal(t, s[0], &StringCount{"1", 2})
	assert.Equal(t, s[1], &StringCount{"2", 1})

	var buf []byte
	buffer := bytes.NewBuffer(buf)
	c.Print(buffer)
	assert.Equal(t, buffer.String(), "2\t1\n1\t2\n")
}
