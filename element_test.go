package xmldoc

import (
	"reflect"
	"testing"
)

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

func TestXDElement_AttributeTraversal(t *testing.T) {
	var name = XDName{LocalName: "fish", NameSpace: "fingers"}

	var n = NewXDElement(name)

	var names = []XDName{XDName{"a", "b"}, XDName{"c", "d"}}
	var values = []string{"1", "2"}

	for i, v := range names {
		n.Attributes[v] = values[i]
	}

	var rname = make([]XDName, 0)
	var rval = make([]string, 0)

	n.TraverseAttributes(func(name XDName, value string) (stop bool) {
		rname = append(rname, name)
		rval = append(rval, value)
		return stop
	})

	if !reflect.DeepEqual(names, rname) {
		t.Error("Names not equal.")
	}

	if !reflect.DeepEqual(rval, values) {
		t.Error("Values not equal.")
	}

}

func TestXDElement_AttributeTraversalStop(t *testing.T) {
	var name = XDName{LocalName: "fish", NameSpace: "fingers"}

	var n = NewXDElement(name)

	var names = []XDName{XDName{"a", "b"}, XDName{"c", "d"}}
	var values = []string{"1", "2"}

	for i, v := range names {
		n.Attributes[v] = values[i]
	}

	var rname = make([]XDName, 0)
	var rval = make([]string, 0)

	n.TraverseAttributes(func(name XDName, value string) bool {
		rname = append(rname, name)
		rval = append(rval, value)
		return true
	})

	if reflect.DeepEqual(names, rname) {
		t.Error("Names should not equal.")
	}

	if reflect.DeepEqual(rval, values) {
		t.Error("Values should not equal.")
	}

	if !reflect.DeepEqual(names[:1], rname) {
		t.Error("Names should  equal.")
	}

	if !reflect.DeepEqual(values[:1], rval) {
		t.Error("Values should equal.")
	}

}
