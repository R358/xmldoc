package xmldoc

// XDNode all nodes implement this interface.
type XDNode interface {
	// GetType return the node type.
	GetType() XDType
	// GetParent return the parent node.
	GetParent() XDNode
	// AddChild add a child node.
	AddChild(XDNode)
	// GetName get the name of this node.
	GetName() XDName
}
