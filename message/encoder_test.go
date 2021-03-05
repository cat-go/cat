package message

import (
	"bytes"
	"testing"
)

func TestEncoderBase_EncodeHeader(t *testing.T) {
	_ = ReadableProtocol
	_ = BinaryProtocol
}


func BenchmarkWriteI64(b *testing.B) {
	var err error
	for i := 0; i < b.N; i++ {
		buf := bytes.NewBuffer([]byte("hello"))
		if err = writeI64(buf, -10000000); err != nil {
			return
		}
	}
}
