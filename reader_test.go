package bit

import (
	"encoding/hex"
	"encoding/json"
	"testing"
)

type Cmd struct {
	Table   int   // 8
	SSI     int   // 1
	Priv    int   // 1
	Res     int   // 2
	Len     int   // 12
	Ver     int   // 8
	Enc     int   // 1
	EncAlg  int   // 6
	PTSA    int64 // 33
	CWI     int   // 8
	Tier    int   // 12
	CmdLen  int   // 12
	CmdType int   // 8
}

func TestReader(t *testing.T) {
	data, _ := hex.DecodeString("FC302500000000000000FFF01405000003E77FEFFE0011FB9EFE002932E00001010100004D192A59")
	r := NewReader(data)
	c := Cmd{}
	r.Decode(&c.Table, 8)
	r.Decode(&c.SSI, 1)
	r.Decode(&c.Priv, 1)
	r.Decode(&c.Res, 2)
	r.Decode(&c.Len, 12)
	r.Decode(&c.Ver, 8) // 4 bytes total
	r.Decode(&c.Enc, 1)
	r.Decode(&c.EncAlg, 6)
	r.Decode(&c.PTSA, 33)
	r.Decode(&c.CWI, 8) // 6 bytes total
	r.Decode(&c.Tier, 12)
	r.Decode(&c.CmdLen, 12)
	r.Decode(&c.CmdType, 8)
	t.Log(js(c))
	if c.CmdType != 5 {
		t.Fatalf("bad bitstream; splice command type should be 0x5 (splice_insert) have=%x", c.CmdType)
	}
}

func js(v any) string {
	data, _ := json.Marshal(v)
	return string(data)
}
