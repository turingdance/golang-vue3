package utils

import (
	"encoding/binary"
	"math"
)

func BytetoFloat32(bts []byte) float32 {
	bits := binary.LittleEndian.Uint32(bts)
	return math.Float32frombits(bits)
}
