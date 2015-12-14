package xmldoc

// NewXDElement creates a new XDElement, parent may be nil.
func NewXDElement(name XDName) *XDElement {
	xd := &XDElement{
		XDBaseNode: &XDBaseNode{Type: ElementType, Children: make([]XDNode, 0), XDName: name},
		Attributes: make(map[XDName]string),
	}
	return xd
}

// XDElement represents an xml element. <camel legs="4" speed="run">Very Bumpy!</camel>
type XDElement struct {
	*XDBaseNode
	Attributes map[XDName]string
}

// TraverseAttributes traverses the attributes passing each one to the supplied function
// traversal stops if the function returns true.
func (ele *XDElement) TraverseAttributes(onAttribute func(XDName, string) bool) {
	for n, v := range ele.Attributes {
		if onAttribute(n, v) {
			break
		}
	}
}
