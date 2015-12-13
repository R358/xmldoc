package xmldoc

import "testing"

func TestXDBaseNode_GetType(t *testing.T) {
	z := XDBaseNode{Type: DirectiveType}
	if z.GetType() != DirectiveType {
		t.Error("XD type incorrect")
	}
}

func TestXDBaseNode_Parent(t *testing.T) {
	p := &XDBaseNode{}

	z := XDBaseNode{Parent: p}
	if z.GetParent() != p {
		t.Error("XD parent incorrect")
	}
}

func TestXDBaseNode_AddChild(t *testing.T) {
	p := &XDBaseNode{}
	z := XDBaseNode{Children: make([]XDNode, 0)}
	z.AddChild(p)

	if len(z.Children) != 1 {
		t.Error("Expecting len = 1")
		if z.Children[0] != p {
			t.Error("Child incorrect.")
		}
	}
}

func TestXDBaseNode_GetName(t *testing.T) {
	tv := XDName{LocalName: "fish", NameSpace: "chips"}
	z := XDBaseNode{XDName: XDName{LocalName: "fish", NameSpace: "chips"}}
	if z.GetName() != tv {
		t.Error("XDName incorrect.")
	}
}
