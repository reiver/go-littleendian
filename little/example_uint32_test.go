package littleendian_test

import (
	"github.com/reiver/go-endian/little"

	"fmt"
	"bytes"
)

func ExampleWriteUint32To() {

	var buffer bytes.Buffer

	var value uint32 = 0x76543210

	n64, err := littleendian.WriteUint32To(&buffer, value)
	if nil != err {
		fmt.Printf("Problem writing: %s", err)
		return
	}

	fmt.Printf("Wrote %d bytes.\n", n64)
	fmt.Printf("Bytes: % x\n", buffer.Bytes())

	// Output:
	// Wrote 4 bytes.
	// Bytes: 10 32 54 76
}

func ExampleReadUint32From() {

	var buffer bytes.Buffer
	buffer.Write([]byte{0x10,0x32,0x54,0x76})

	var value uint32

	n64, err := littleendian.ReadUint32From(&buffer, &value)
	if nil != err {
		fmt.Printf("Problem writing: %s", err)
		return
	}

	fmt.Printf("Read %d bytes.\n", n64)
	fmt.Printf("Value: 0x%x\n", value)

	// Output:
	// Read 4 bytes.
	// Value: 0x76543210
}
