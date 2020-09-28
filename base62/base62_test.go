package base62

import (
	"math/rand"
	"testing"
)

func TestEncodeDecode(t *testing.T) {
	input := rand.Uint32()
	output := Inst.Encode(input)
	if input != Inst.Decode(output) {
		t.Error("Inst.Encode != Inst.Decode")
		return
	}
}
