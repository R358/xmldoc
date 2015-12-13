package xmldoc

import (
	"bytes"
	"testing"
)

func TestNewXDirective(t *testing.T) {
	tv := []byte("Cats")
	n := NewXDirective(tv)

	if n.GetType() != DirectiveType {
		t.Error("Incorrect type for comment")
	}

	if !bytes.Equal(n.Directive, tv) {
		t.Error("Incorrect data.")
	}
}
