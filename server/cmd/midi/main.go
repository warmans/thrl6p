package main

import (
	"encoding/hex"
	"fmt"
	"golang.org/x/text/encoding"
)

func main() {

//	sysex := `f0 00010c 24024d 01 08 000d01 08000000000200020100000001000000000200505352500000000000040008000000004561726c790056480001000000000200000000000200000000020000000000000300000003000000000000020020000000000000004000020001000000080100000002005401525447030000000007003c010000051000000007000801020000060000000200000b0000001801000000010000000000001b010000010000000000001a0100020003005b0b2b3e001e010000010001000000001d0100000803004363593e21000100000100010000000028010000030000185b483e230100000001000100000000260100000200 f7`
	payload := `08000000000200020100000001000000000200505352500000000000040008000000004561726c790056480001000000000200000000000200000000020000000000000300000003000000000000020020000000000000004000020001000000080100000002005401525447030000000007003c010000051000000007000801020000060000000200000b0000001801000000010000000000001b010000010000000000001a0100020003005b0b2b3e001e010000010001000000001d0100000803004363593e21000100000100010000000028010000030000185b483e230100000001000100000000260100000200`
	//eightBit := `10000000008002020000002000000008050A74A800000000008004000000045C3CB679015A400020000000080000000002000000004000000000003000000300000000000100400000000000008000100020000010040000008054034AA470600000001C03C020000520000000E004010400006000000200005800000C01000000100000000000D810000080000000000680800400180B62D5BE0078080000400100000003A04000100C043C765F2100040000200080000002802000030000C5B90F91810000001000400000013010000100`

	b, err := hex.DecodeString(payload)
	if err != nil {
		panic(err)
	}
	for _, by := range b {
		fmt.Println(hex.EncodeToString([]byte{by}), " -> ", string(by << 00))
	}
}

func Decode(enc []byte, en encoding.Encoding) string {
	dec := en.NewDecoder()
	out, err := dec.Bytes(enc)
	if err != nil {
		panic(err)
	}
	return string(out)
}