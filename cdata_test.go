package xmldoc

import (
	"bytes"
	"testing"
)

func TestNewXDCData(t *testing.T) {
	tv := []byte("Cats")
	n := NewXDCData(tv)

	if n.GetType() != CDataType {
		t.Error("Incorrect type for CDATA")
	}

	if !bytes.Equal(n.Data, tv) {
		t.Error("Incorrect data.")
	}
}
