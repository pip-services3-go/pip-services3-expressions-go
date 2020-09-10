package test_io

import (
	"testing"

	"github.com/pip-services3-go/pip-services3-expressions-go/io"
)

func TestOperations(t *testing.T) {
	content := "Test String"
	reader := io.NewStringPushbackReader(content)
	fixture := NewPushbackReaderFixture(reader, content)

	fixture.TestOperations(t)
}

func TestPushback(t *testing.T) {
	content := "Test String"
	reader := io.NewStringPushbackReader(content)
	fixture := NewPushbackReaderFixture(reader, content)

	fixture.TestPushback(t)
}
