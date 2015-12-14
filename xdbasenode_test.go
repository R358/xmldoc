package xmldoc

import (
	"reflect"
	"testing"
)

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

func TestXDBaseNode_Traverse(t *testing.T) {

	var expected = []XDNode{
		NewXDCData([]byte("Chips")),
		NewXDElement(XDName{"cats", "fish"}),
		NewXDComment([]byte("Bananas")),
	}

	var bn = &XDBaseNode{Children: make([]XDNode, 0)}

	for _, v := range expected {
		bn.AddChild(v)
	}

	var result = make([]XDNode, 0)

	bn.TraverseChildren(func(child XDNode) (stop bool) {
		result = append(result, child)
		return stop
	})

	if !reflect.DeepEqual(result, expected) {
		t.Error("Children traversed did not match.")
	}

}

func TestXDBaseNode_TraverseStop(t *testing.T) {

	expected := []XDNode{
		NewXDCData([]byte("Chips")),
		NewXDElement(XDName{"cats", "fish"}),
		NewXDComment([]byte("Bananas")),
	}

	var bn = &XDBaseNode{Children: make([]XDNode, 0)}

	for _, v := range expected {
		bn.AddChild(v)
	}

	var result = make([]XDNode, 0)

	bn.TraverseChildren(func(child XDNode) bool {
		result = append(result, child)
		return true
	})

	if reflect.DeepEqual(result, expected) {
		t.Error("Children traversed should not match.")
	}

	if !reflect.DeepEqual(result, expected[:1]) {
		t.Error("Children traversed should match slice.")
	}

}
