package littleendian_test

import (
	"github.com/reiver/go-endian/little"

	"fmt"
	"bytes"
)

func ExampleWriteUint16To() {

	var buffer bytes.Buffer

	var value uint16 = 0x3210

	n64, err := littleendian.WriteUint16To(&buffer, value)
	if nil != err {
		fmt.Printf("Problem writing: %s", err)
		return
	}

	fmt.Printf("Wrote %d bytes.\n", n64)
	fmt.Printf("Bytes: % x\n", buffer.Bytes())

	// Output:
	// Wrote 2 bytes.
	// Bytes: 10 32
}