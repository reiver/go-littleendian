package littleendian

import (
	"bytes"
	"math/rand"
	"time"

	"testing"
)

func TestWriteUint32To(t *testing.T) {

	tests := []struct{
		Value uint32
		Expected []byte
	}{
		{
			Value:     0x00000000,
			Expected: []byte{0x00,0x00,0x00,0x00},

		},
		{
			Value:     0x00000001,
			Expected: []byte{0x01,0x00,0x00,0x00},

		},
		{
			Value:     0x00000002,
			Expected: []byte{0x02,0x00,0x00,0x00},

		},
		{
			Value:     0x00000003,
			Expected: []byte{0x03,0x00,0x00,0x00},

		},
		{
			Value:     0x00000004,
			Expected: []byte{0x04,0x00,0x00,0x00},

		},
		{
			Value:     0x00000005,
			Expected: []byte{0x05,0x00,0x00,0x00},

		},
		{
			Value:     0x00000006,
			Expected: []byte{0x06,0x00,0x00,0x00},

		},
		{
			Value:     0x00000007,
			Expected: []byte{0x07,0x00,0x00,0x00},

		},
		{
			Value:     0x00000008,
			Expected: []byte{0x08,0x00,0x00,0x00},

		},
		{
			Value:     0x00000009,
			Expected: []byte{0x09,0x00,0x00,0x00},

		},
		{
			Value:     0x0000000a,
			Expected: []byte{0x0a,0x00,0x00,0x00},

		},
		{
			Value:     0x0000000b,
			Expected: []byte{0x0b,0x00,0x00,0x00},

		},
		{
			Value:     0x0000000c,
			Expected: []byte{0x0c,0x00,0x00,0x00},

		},
		{
			Value:     0x0000000d,
			Expected: []byte{0x0d,0x00,0x00,0x00},

		},
		{
			Value:     0x0000000e,
			Expected: []byte{0x0e,0x00,0x00,0x00},

		},
		{
			Value:     0x0000000f,
			Expected: []byte{0x0f,0x00,0x00,0x00},

		},
		{
			Value:     0x00000010,
			Expected: []byte{0x10,0x00,0x00,0x00},

		},
		{
			Value:     0x00000011,
			Expected: []byte{0x11,0x00,0x00,0x00},

		},
		{
			Value:     0x00000012,
			Expected: []byte{0x12,0x00,0x00,0x00},

		},
		{
			Value:     0x00000013,
			Expected: []byte{0x13,0x00,0x00,0x00},

		},

		{
			Value:     0x000000ef,
			Expected: []byte{0xef,0x00,0x00,0x00},

		},
		{
			Value:     0x000000f0,
			Expected: []byte{0xf0,0x00,0x00,0x00},

		},
		{
			Value:     0x000000f1,
			Expected: []byte{0xf1,0x00,0x00,0x00},

		},
		{
			Value:     0x000000f2,
			Expected: []byte{0xf2,0x00,0x00,0x00},

		},
		{
			Value:     0x000000f3,
			Expected: []byte{0xf3,0x00,0x00,0x00},

		},
		{
			Value:     0x000000f4,
			Expected: []byte{0xf4,0x00,0x00,0x00},

		},
		{
			Value:     0x000000f5,
			Expected: []byte{0xf5,0x00,0x00,0x00},

		},
		{
			Value:     0x000000f6,
			Expected: []byte{0xf6,0x00,0x00,0x00},

		},
		{
			Value:     0x000000f7,
			Expected: []byte{0xf7,0x00,0x00,0x00},

		},
		{
			Value:     0x000000f8,
			Expected: []byte{0xf8,0x00,0x00,0x00},

		},
		{
			Value:     0x000000f9,
			Expected: []byte{0xf9,0x00,0x00,0x00},

		},
		{
			Value:     0x000000fa,
			Expected: []byte{0xfa,0x00,0x00,0x00},

		},
		{
			Value:     0x000000fb,
			Expected: []byte{0xfb,0x00,0x00,0x00},

		},
		{
			Value:     0x000000fc,
			Expected: []byte{0xfc,0x00,0x00,0x00},

		},
		{
			Value:     0x000000fd,
			Expected: []byte{0xfd,0x00,0x00,0x00},

		},
		{
			Value:     0x000000fe,
			Expected: []byte{0xfe,0x00,0x00,0x00},

		},
		{
			Value:     0x000000ff,
			Expected: []byte{0xff,0x00,0x00,0x00},

		},
		{
			Value:     0x00000100,
			Expected: []byte{0x00,0x01,0x00,0x00},

		},
		{
			Value:     0x00000101,
			Expected: []byte{0x01,0x01,0x00,0x00},

		},
		{
			Value:     0x00000102,
			Expected: []byte{0x02,0x01,0x00,0x00},

		},
		{
			Value:     0x00000103,
			Expected: []byte{0x03,0x01,0x00,0x00},

		},

		{
			Value:     0x00000245,
			Expected: []byte{0x45,0x02,0x00,0x00},

		},

		{
			Value:     0x000037ae,
			Expected: []byte{0xae,0x37,0x00,0x00},

		},

		{
			Value:     0x0000fffa,
			Expected: []byte{0xfa,0xff,0x00,0x00},

		},
		{
			Value:     0x0000fffb,
			Expected: []byte{0xfb,0xff,0x00,0x00},

		},
		{
			Value:     0x0000fffc,
			Expected: []byte{0xfc,0xff,0x00,0x00},

		},
		{
			Value:     0x0000fffd,
			Expected: []byte{0xfd,0xff,0x00,0x00},

		},
		{
			Value:     0x0000fffe,
			Expected: []byte{0xfe,0xff,0x00,0x00},

		},
		{
			Value:     0x0000ffff,
			Expected: []byte{0xff,0xff,0x00,0x00},

		},
		{
			Value:     0x00010000,
			Expected: []byte{0x00,0x00,0x01,0x00},

		},
		{
			Value:     0x00010001,
			Expected: []byte{0x01,0x00,0x01,0x00},

		},
		{
			Value:     0x00010002,
			Expected: []byte{0x02,0x00,0x01,0x00},

		},
		{
			Value:     0x00010003,
			Expected: []byte{0x03,0x00,0x01,0x00},

		},

		{
			Value:     0x1357bd24,
			Expected: []byte{0x24,0xbd,0x57,0x13},

		},

		{
			Value:     0xffffffff,
			Expected: []byte{0xff,0xff,0xff,0xff},

		},
	}


	{
		randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))

		const max = 50
		for i:=0; i<max; i++ {
			x64 := randomness.Int63n(0xffffffff)

			u32 := uint32(x64)

			var b0 byte = uint8( 0x000000ff & u32      )
			var b1 byte = uint8((0x0000ff00 & u32) >> 8)
			var b2 byte = uint8((0x00ff0000 & u32) >> 16)
			var b3 byte = uint8((0xff000000 & u32) >> 24)

			test := struct{
				Value uint32
				Expected []byte
			}{
				Value:           u32,
				Expected: []byte{b0, b1, b2, b3},
			}

			tests = append(tests, test)
		}
	}


	for testNumber, test := range tests {

		var buffer bytes.Buffer

		n64, err := WriteUint32To(&buffer, test.Value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually go one: (%T) %q", testNumber, err, err)
			continue
		}

		if expected, actual := int64(len(test.Expected)), n64; expected != actual {
			t.Errorf("For test #%d, expected %d, but actually got %d.", testNumber, expected, actual)
			continue
		}

		if expected, actual := test.Expected, buffer.Bytes(); !bytes.Equal(expected, actual) {
			t.Errorf("For test #%d,....", testNumber)
			t.Errorf("\tEXPECTED: % v", expected)
			t.Errorf("\tACTUAL:   % v", actual)
			continue
		}
	}
}