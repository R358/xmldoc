package xmldoc

// NewXDCDocument creates a new NewXDCDocument.
func NewXDCDocument() *XDCDocument {
	xd := &XDCDocument{
		XDBaseNode: &XDBaseNode{Type: DocumentType},
	}
	return xd
}

// XDCDocument represents an xml document.
type XDCDocument struct {
	*XDBaseNode
	Root *XDElement
}
