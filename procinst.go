package xmldoc

// NewXDProcessingInstruction creates a new XDProcessingInstruction.
func NewXDProcessingInstruction(target string, instruction []byte) *XDProcessingInstruction {
	xd := &XDProcessingInstruction{
		XDBaseNode:  &XDBaseNode{Type: ProcessingInstructionType},
		Target:      target,
		Instruction: instruction,
	}
	return xd
}

// XDProcessingInstruction represents a processing instruction. <?xml ... ?>
type XDProcessingInstruction struct {
	*XDBaseNode
	Instruction []byte
	Target      string
	Type        XDType
}
