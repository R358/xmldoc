package xmldoc

import (
	"bytes"
	"testing"
)

func Test_Parse_SingleElement(t *testing.T) {
	var xml = `<?xml version="1.0" encoding="UTF-8"?><content></content>`

	doc, e := Parse(bytes.NewBufferString(xml))

	if e != nil {
		t.Error("Unexpected error: ", e.Error())
	}

	if len(doc.Children) != 2 {
		t.Error("Child length incorrect.", len(doc.Children))
	}

	if doc.Children[0].GetType() != ProcessingInstructionType {
		t.Error("Expecing processing instruction at 0")
	}

	if doc.Children[1].GetType() != ElementType {
		t.Error("Expecing element at 1")
	}

	if doc.Children[1].GetName() != (XDName{"content", ""}) {
		t.Error("Incorrect name.")
	}

}

func Test_Parse_SingleElement_Attributes(t *testing.T) {
	var xml = `<?xml version="1.0" encoding="UTF-8"?><content fish="chip" cow="milk"></content>`

	doc, e := Parse(bytes.NewBufferString(xml))

	if e != nil {
		t.Error("Unexpected error: ", e.Error())
	}

	if len(doc.Children) != 2 {
		t.Error("Child length incorrect.", len(doc.Children))
	}

	if doc.Children[0].GetType() != ProcessingInstructionType {
		t.Error("Expecing processing instruction at 0")
	}

	if doc.Children[1].GetType() != ElementType {
		t.Error("Expecing element at 1")
	}

	if doc.Children[1].GetName() != (XDName{"content", ""}) {
		t.Error("Incorrect name.")
	}

	attr := doc.Children[1].(*XDElement).Attributes

	if a, ok := attr[(XDName{LocalName: "fish"})]; !ok || a != "chip" {
		t.Error("Incorrect fish attribute.")
	}

	if a, ok := attr[(XDName{LocalName: "cow"})]; !ok || a != "milk" {
		t.Error("Incorrect cow attribute.")
	}
}

func Test_Parse_SingleElement_CData(t *testing.T) {
	var xml = `<?xml version="1.0" encoding="UTF-8"?><content>Fish</content>`

	doc, e := Parse(bytes.NewBufferString(xml))

	if e != nil {
		t.Error("Unexpected error: ", e.Error())
	}

	if len(doc.Children) != 2 {
		t.Error("Child length incorrect.", len(doc.Children))
	}

	if doc.Children[0].GetType() != ProcessingInstructionType {
		t.Error("Expecing processing instruction at 0")
	}

	if doc.Children[1].GetType() != ElementType {
		t.Error("Expecing element at 1")
	}

	if doc.Children[1].GetName() != (XDName{"content", ""}) {
		t.Error("Incorrect name.")
	}

	var el = doc.Children[1].(*XDElement)

	if el.Children[0].GetType() != CDataType {
		t.Error("Expecting Cdata type.")
	}

	if !bytes.Equal(el.Children[0].(*XDCData).Data, []byte("Fish")) {
		t.Error("Incorrect value.", string(el.Children[0].(*XDCData).Data))
	}

}

func Test_Parse_SingleElement_Comment(t *testing.T) {
	var xml = `<?xml version="1.0" encoding="UTF-8"?><content><!--  Banana --></content>`

	doc, e := Parse(bytes.NewBufferString(xml))

	if e != nil {
		t.Error("Unexpected error: ", e.Error())
	}

	if len(doc.Children) != 2 {
		t.Error("Child length incorrect.", len(doc.Children))
	}

	if doc.Children[0].GetType() != ProcessingInstructionType {
		t.Error("Expecing processing instruction at 0")
	}

	if doc.Children[1].GetType() != ElementType {
		t.Error("Expecing element at 1")
	}

	if doc.Children[1].GetName() != (XDName{"content", ""}) {
		t.Error("Incorrect name.")
	}

	var el = doc.Children[1].(*XDElement)

	if el.Children[0].GetType() != CommentType {
		t.Error("Expecting Cdata type.")
	}

	if !bytes.Equal(el.Children[0].(*XDComment).Comment, []byte("  Banana ")) {
		t.Error("Incorrect value.", string(el.Children[0].(*XDComment).Comment))
	}
}

func Test_Parse_SingleElement_Directivet(t *testing.T) {
	var xml = `<?xml version="1.0" encoding="UTF-8"?><content><!Aardvark></content>`

	doc, e := Parse(bytes.NewBufferString(xml))

	if e != nil {
		t.Error("Unexpected error: ", e.Error())
	}

	if len(doc.Children) != 2 {
		t.Error("Child length incorrect.", len(doc.Children))
	}

	if doc.Children[0].GetType() != ProcessingInstructionType {
		t.Error("Expecing processing instruction at 0")
	}

	if doc.Children[1].GetType() != ElementType {
		t.Error("Expecing element at 1")
	}

	if doc.Children[1].GetName() != (XDName{"content", ""}) {
		t.Error("Incorrect name.")
	}

	var el = doc.Children[1].(*XDElement)

	if el.Children[0].GetType() != DirectiveType {
		t.Error("Expecting Cdata type.")
	}

	if !bytes.Equal(el.Children[0].(*XDDirective).Directive, []byte("Aardvark")) {
		t.Error("Incorrect value.", string(el.Children[0].(*XDDirective).Directive))
	}
}
