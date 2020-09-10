package test_io

import (
	"testing"

	"github.com/pip-services3-go/pip-services3-expressions-go/io"
	"github.com/pip-services3-go/pip-services3-expressions-go/tokenizers/utilities"
	"github.com/stretchr/testify/assert"
)

type PushbackReaderFixture struct {
	reader  io.IPushbackReader
	content []rune
}

func NewPushbackReaderFixture(reader io.IPushbackReader, content string) *PushbackReaderFixture {
	return &PushbackReaderFixture{
		reader:  reader,
		content: []rune(content),
	}
}

func (c *PushbackReaderFixture) TestOperations(t *testing.T) {
	chr, err := c.reader.Peek()
	assert.Nil(t, err)
	assert.Equal(t, c.content[0], chr)

	chr, err = c.reader.Read()
	assert.Nil(t, err)
	assert.Equal(t, c.content[0], chr)

	chr, err = c.reader.Read()
	assert.Nil(t, err)
	assert.Equal(t, c.content[1], chr)

	c.reader.Pushback('#')
	chr, err = c.reader.Read()
	assert.Nil(t, err)
	assert.Equal(t, '#', chr)

	c.reader.PushbackString("@$")
	chr, err = c.reader.Read()
	assert.Nil(t, err)
	assert.Equal(t, '@', chr)
	chr, err = c.reader.Read()
	assert.Nil(t, err)
	assert.Equal(t, '$', chr)

	for i := 2; i < len(c.content); i++ {
		chr, err = c.reader.Read()
		assert.Nil(t, err)
		assert.Equal(t, c.content[i], chr)
	}

	chr, err = c.reader.Read()
	assert.Nil(t, err)
	assert.Equal(t, rune(-1), chr)

	chr, err = c.reader.Read()
	assert.Nil(t, err)
	assert.Equal(t, rune(-1), chr)
}

func (c *PushbackReaderFixture) TestPushback(t *testing.T) {
	var chr rune
	var lastChr rune
	var err error
	for chr, err = c.reader.Read(); err == nil && !utilities.CharValidator.IsEof(chr); chr, err = c.reader.Read() {
		assert.Nil(t, err)
		lastChr = chr
	}

	c.reader.Pushback(lastChr)
	c.reader.Pushback(chr)

	chr1, err1 := c.reader.Peek()
	assert.Nil(t, err1)
	assert.Equal(t, lastChr, chr1)

	chr1, err1 = c.reader.Read()
	assert.Nil(t, err1)
	assert.Equal(t, lastChr, chr1)
}
