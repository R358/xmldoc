package xmldoc

import (
	"bytes"
	"testing"
)

func TestNewXDProcessingInstruction(t *testing.T) {

	var tgt = "foo"
	var instr = []byte("Cata")

	n := NewXDProcessingInstruction(tgt, instr)

	if n.GetType() != ProcessingInstructionType {
		t.Error("Incorrect type for processing instruction.")
	}

	if n.Target != tgt {
		t.Error("Target not set.")
	}

	if !bytes.Equal(n.Instruction, instr) {
		t.Error("Instruction not equal.")
	}

}
