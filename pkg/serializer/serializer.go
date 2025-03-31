package serializer

import (
	"bytes"
	"encoding/binary"
)

func Int(n int) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, int64(n))
	return buf.Bytes()
}
