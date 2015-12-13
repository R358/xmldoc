package xmldoc

// NewXDComment creates a new XDComment.
func NewXDComment(value []byte) *XDComment {
	xd := &XDComment{
		XDBaseNode: &XDBaseNode{Type: CommentType},
		Comment:    value,
	}
	return xd
}

// XDComment represents an xml comment <!-- Foo -->
type XDComment struct {
	*XDBaseNode
	Comment []byte
	Type    XDType
}
