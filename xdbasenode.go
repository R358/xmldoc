package xmldoc

// XDBaseNode is the base node for all XD nodes.
type XDBaseNode struct {
	Type     XDType
	Parent   XDNode
	Children []XDNode
	XDName   XDName
}

// GetType return the type of the node.
func (xd *XDBaseNode) GetType() XDType {
	return xd.Type
}

// GetParent return the parent of the node.
func (xd *XDBaseNode) GetParent() XDNode {
	return xd.Parent
}

// AddChild return the child of the node.
func (xd *XDBaseNode) AddChild(n XDNode) {
	xd.Children = append(xd.Children, n)
}

// GetName return the name of the node.
func (xd *XDBaseNode) GetName() XDName {
	return xd.XDName
}