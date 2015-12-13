package xmldoc

import (
	"testing"
)

func TestNewXDCDocument(t *testing.T) {
	n := NewXDCDocument()

	if n.GetType() != DocumentType {
		t.Error("Incorrect type for document")
	}

	if n.Children != nil {
		t.Error("Children should be uninitialised.")
	}

}
