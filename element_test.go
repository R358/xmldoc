package xmldoc

import "testing"

func TestNewXDElement(t *testing.T) {

	var name = XDName{LocalName: "fish", NameSpace: "fingers"}

	n := NewXDElement(name)

	if n.GetType() != ElementType {
		t.Error("Incorrect type for element")
	}

	if n.Attributes == nil {
		t.Error("Attributes should be initialized.")
	}

}
