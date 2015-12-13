package xmldoc

// NewXDCData creates a new XDCData, parent is nil.
func NewXDCData(value []byte) *XDCData {
	xd := &XDCData{
		XDBaseNode: &XDBaseNode{Type: CDataType},
		Data:       value,
	}
	return xd
}

// XDCData represents character data.
type XDCData struct {
	*XDBaseNode
	Data []byte
	Type XDType
}
