package xmldoc

// NewXDirective creates a new XDDirective node.
func NewXDirective(value []byte) *XDDirective {
	xd := &XDDirective{
		XDBaseNode: &XDBaseNode{Type: DirectiveType},
		Directive:  value,
	}
	return xd
}

// XDDirective represents an xml directive. Eg <!bang>
type XDDirective struct {
	*XDBaseNode
	Directive []byte
	Type      XDType
}
