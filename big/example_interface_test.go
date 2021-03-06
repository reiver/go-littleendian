package bigendian_test

import (
	"github.com/reiver/go-endian/big"

	"fmt"
	"bytes"
)

func ExampleWriteTo() {

	var buffer bytes.Buffer

	type Location struct{
		X uint16
		Y uint16
		Meta struct {
			Name [8]uint8
			Size uint64
		}
	}

	var value [2]Location

	value[0].X = uint16(0x0001)
	value[0].Y = uint16(0xabcd)
	value[0].Meta.Name = [8]uint8{'H','e','r','e',0x00,0x00,0x00,0x00}
	value[0].Meta.Size = uint64(0x00000000000000de)

	value[1].X = uint16(0x531f)
	value[1].Y = uint16(0x8ace)
	value[1].Meta.Name = [8]uint8{'T','h','e','r','e',0x00,0x00,0x00}
	value[1].Meta.Size = uint64(0x0123456789abcedf)


	n64, err := bigendian.WriteTo(&buffer, value)
	if nil != err {
		fmt.Printf("Problem writing: %s", err)
		return
	}

	fmt.Printf("Wrote %d bytes.\n", n64)
	fmt.Printf("Bytes: % x\n", buffer.Bytes())

	// Output:
	// Wrote 40 bytes.
	// Bytes: 00 01 ab cd 48 65 72 65 00 00 00 00 00 00 00 00 00 00 00 de 53 1f 8a ce 54 68 65 72 65 00 00 00 01 23 45 67 89 ab ce df
}
