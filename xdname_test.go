package xmldoc

import "testing"

func TestXDName_Namespace(t *testing.T) {
	xd := XDName{LocalName: "fish", NameSpace: "fingers"}
	if xd.String() != "fingers:fish" {
		t.Error("Name incorrect: ", xd.String())
	}
}

func TestXDName_NoNameSpace(t *testing.T) {
	xd := XDName{LocalName: "fish"}
	if xd.String() != "fish" {
		t.Error("Name incorrect.")
	}
}
