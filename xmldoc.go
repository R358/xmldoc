package xmldoc

import (
	"encoding/xml"
	"io"
)

// XDType defines the type of the xml node.
type XDType int

var (
	// ElementType is an xml element.
	ElementType = XDType(0)
	// CDataType is xml character data.
	CDataType = XDType(1)
	// CommentType is xml comment.
	CommentType = XDType(2)
	// DirectiveType is an xml directive.
	DirectiveType = XDType(3)
	// ProcessingInstructionType is an xml processing instruction <!bang>
	ProcessingInstructionType = XDType(4)
	// DocumentType is the type from the whole document.
	DocumentType = XDType(5)
)

// Parse read and parse data from the reader.
func Parse(r io.Reader) (doc *XDCDocument, err error) {
	doc = NewXDCDocument()
	parser := xml.NewDecoder(r)
	var token xml.Token

	var parent XDNode = doc

	for token, err = parser.Token(); err == nil && token != nil; token, err = parser.Token() {

		switch tl := token.(type) {
		case xml.StartElement:
			j := NewXDElement(XDName{LocalName: tl.Name.Local, NameSpace: tl.Name.Space})

			if tl.Attr != nil {
				for _, v := range tl.Attr {
					j.Attributes[XDName{LocalName: v.Name.Local, NameSpace: v.Name.Space}] = v.Value
				}
			}
			if doc.Root == nil {
				doc.Root = j
			}
			parent.AddChild(j)
			j.Parent = parent
			parent = j
			break

		case xml.EndElement:
			parent = parent.GetParent()
			break

		case xml.CharData:
			ch := NewXDCData(tl.Copy())
			parent.AddChild(ch)
			ch.Parent = parent

		case xml.Comment:
			ch := NewXDComment(tl.Copy())
			ch.Parent = parent
			parent.AddChild(ch)

		case xml.ProcInst:
			pi := NewXDProcessingInstruction(tl.Target, tl.Inst)
			pi.Parent = parent
			parent.AddChild(pi)

		case xml.Directive:
			dir := NewXDirective(tl.Copy())
			dir.Parent = parent
			parent.AddChild(dir)
		}
	}

	return doc, err
}
