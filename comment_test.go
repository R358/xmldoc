package xmldoc

import (
	"bytes"
	"testing"
)

func TestNewXDComment(t *testing.T) {
	tv := []byte("Cats")
	n := NewXDComment(tv)

	if n.GetType() != CommentType {
		t.Error("Incorrect type for comment")
	}

	if !bytes.Equal(n.Comment, tv) {
		t.Error("Incorrect data.")
	}
}
