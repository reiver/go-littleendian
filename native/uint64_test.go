package nativeendian_test

import (
	"github.com/reiver/go-endian"
	"github.com/reiver/go-endian/native"

	"bytes"
	"math/rand"
	"time"

	"testing"
)

func TestWriteUint64To(t *testing.T) {

	var tests []struct{
		Value uint64
		Expected []byte
	}


	var endianness endian.Type = endian.NativeEndianness()

	switch endianness {
	default:
		panic(endianness)
	case endian.Little():
		tests = append(tests, []struct{
			Value uint64
			Expected []byte
		}{
			{
				Value: 0x0000000000000000,
				Expected:     []byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00},
			},
			{
				Value: 0x0000000000000001,
				Expected:     []byte{0x01,0x00,0x00,0x00,0x00,0x00,0x00,0x00},
			},
			{
				Value: 0x0000000000000002,
				Expected:     []byte{0x02,0x00,0x00,0x00,0x00,0x00,0x00,0x00},
			},
			{
				Value: 0x0000000000000003,
				Expected:     []byte{0x03,0x00,0x00,0x00,0x00,0x00,0x00,0x00},
			},
			{
				Value: 0x0000000000000004,
				Expected:     []byte{0x04,0x00,0x00,0x00,0x00,0x00,0x00,0x00},
			},
			{
				Value: 0x0000000000000005,
				Expected:     []byte{0x05,0x00,0x00,0x00,0x00,0x00,0x00,0x00},
			},
			{
				Value: 0x0000000000000006,
				Expected:     []byte{0x06,0x00,0x00,0x00,0x00,0x00,0x00,0x00},
			},
			{
				Value: 0x0000000000000007,
				Expected:     []byte{0x07,0x00,0x00,0x00,0x00,0x00,0x00,0x00},
			},
			{
				Value: 0x0000000000000008,
				Expected:     []byte{0x08,0x00,0x00,0x00,0x00,0x00,0x00,0x00},
			},
			{
				Value: 0x0000000000000009,
				Expected:     []byte{0x09,0x00,0x00,0x00,0x00,0x00,0x00,0x00},
			},
			{
				Value: 0x000000000000000a,
				Expected:     []byte{0x0a,0x00,0x00,0x00,0x00,0x00,0x00,0x00},
			},
			{
				Value: 0x000000000000000b,
				Expected:     []byte{0x0b,0x00,0x00,0x00,0x00,0x00,0x00,0x00},
			},
			{
				Value: 0x000000000000000c,
				Expected:     []byte{0x0c,0x00,0x00,0x00,0x00,0x00,0x00,0x00},
			},
			{
				Value: 0x000000000000000d,
				Expected:     []byte{0x0d,0x00,0x00,0x00,0x00,0x00,0x00,0x00},
			},
			{
				Value: 0x000000000000000e,
				Expected:     []byte{0x0e,0x00,0x00,0x00,0x00,0x00,0x00,0x00},
			},
			{
				Value: 0x000000000000000f,
				Expected:     []byte{0x0f,0x00,0x00,0x00,0x00,0x00,0x00,0x00},
			},
			{
				Value: 0x0000000000000010,
				Expected:     []byte{0x10,0x00,0x00,0x00,0x00,0x00,0x00,0x00},
			},
			{
				Value: 0x0000000000000011,
				Expected:     []byte{0x11,0x00,0x00,0x00,0x00,0x00,0x00,0x00},
			},
			{
				Value: 0x0000000000000012,
				Expected:     []byte{0x12,0x00,0x00,0x00,0x00,0x00,0x00,0x00},
			},
			{
				Value: 0x0000000000000013,
				Expected:     []byte{0x13,0x00,0x00,0x00,0x00,0x00,0x00,0x00},
			},

			{
				Value: 0x00000000000000ef,
				Expected:     []byte{0xef,0x00,0x00,0x00,0x00,0x00,0x00,0x00},
			},
			{
				Value: 0x00000000000000f0,
				Expected:     []byte{0xf0,0x00,0x00,0x00,0x00,0x00,0x00,0x00},
			},
			{
				Value: 0x00000000000000f1,
				Expected:     []byte{0xf1,0x00,0x00,0x00,0x00,0x00,0x00,0x00},
			},
				{
				Value: 0x00000000000000f2,
				Expected:     []byte{0xf2,0x00,0x00,0x00,0x00,0x00,0x00,0x00},
			},
			{
				Value: 0x00000000000000f3,
				Expected:     []byte{0xf3,0x00,0x00,0x00,0x00,0x00,0x00,0x00},
			},
			{
				Value: 0x00000000000000f4,
				Expected:     []byte{0xf4,0x00,0x00,0x00,0x00,0x00,0x00,0x00},
			},
			{
				Value: 0x00000000000000f5,
				Expected:     []byte{0xf5,0x00,0x00,0x00,0x00,0x00,0x00,0x00},
			},
			{
				Value: 0x00000000000000f6,
				Expected:     []byte{0xf6,0x00,0x00,0x00,0x00,0x00,0x00,0x00},
			},
			{
				Value: 0x00000000000000f7,
				Expected:     []byte{0xf7,0x00,0x00,0x00,0x00,0x00,0x00,0x00},
			},
			{
				Value: 0x00000000000000f8,
				Expected:     []byte{0xf8,0x00,0x00,0x00,0x00,0x00,0x00,0x00},
			},
			{
				Value: 0x00000000000000f9,
				Expected:     []byte{0xf9,0x00,0x00,0x00,0x00,0x00,0x00,0x00},
			},
			{
				Value: 0x00000000000000fa,
				Expected:     []byte{0xfa,0x00,0x00,0x00,0x00,0x00,0x00,0x00},
			},
			{
				Value: 0x00000000000000fb,
				Expected:     []byte{0xfb,0x00,0x00,0x00,0x00,0x00,0x00,0x00},
			},
			{
				Value: 0x00000000000000fc,
				Expected:     []byte{0xfc,0x00,0x00,0x00,0x00,0x00,0x00,0x00},
			},
			{
				Value: 0x00000000000000fd,
				Expected:     []byte{0xfd,0x00,0x00,0x00,0x00,0x00,0x00,0x00},
			},
			{
				Value: 0x00000000000000fe,
				Expected:     []byte{0xfe,0x00,0x00,0x00,0x00,0x00,0x00,0x00},
			},
			{
				Value: 0x00000000000000ff,
				Expected:     []byte{0xff,0x00,0x00,0x00,0x00,0x00,0x00,0x00},
			},
			{
				Value: 0x0000000000000100,
				Expected:     []byte{0x00,0x01,0x00,0x00,0x00,0x00,0x00,0x00},
			},
			{
				Value: 0x0000000000000101,
				Expected:     []byte{0x01,0x01,0x00,0x00,0x00,0x00,0x00,0x00},
			},
			{
				Value: 0x0000000000000102,
				Expected:     []byte{0x02,0x01,0x00,0x00,0x00,0x00,0x00,0x00},
			},
			{
				Value: 0x0000000000000103,
				Expected:     []byte{0x03,0x01,0x00,0x00,0x00,0x00,0x00,0x00},
			},

			{
				Value: 0x0000000000000245,
				Expected:     []byte{0x45,0x02,0x00,0x00,0x00,0x00,0x00,0x00},
			},

			{
				Value: 0x00000000000037ae,
				Expected:     []byte{0xae,0x37,0x00,0x00,0x00,0x00,0x00,0x00},
			},

			{
				Value: 0x000000000000fffa,
				Expected:     []byte{0xfa,0xff,0x00,0x00,0x00,0x00,0x00,0x00},
			},
			{
				Value: 0x000000000000fffb,
				Expected:     []byte{0xfb,0xff,0x00,0x00,0x00,0x00,0x00,0x00},
			},
			{
				Value: 0x000000000000fffc,
				Expected:     []byte{0xfc,0xff,0x00,0x00,0x00,0x00,0x00,0x00},
			},
			{
				Value: 0x000000000000fffd,
				Expected:     []byte{0xfd,0xff,0x00,0x00,0x00,0x00,0x00,0x00},
			},
			{
				Value: 0x000000000000fffe,
				Expected:     []byte{0xfe,0xff,0x00,0x00,0x00,0x00,0x00,0x00},
			},
			{
				Value: 0x000000000000ffff,
				Expected:     []byte{0xff,0xff,0x00,0x00,0x00,0x00,0x00,0x00},
			},
			{
				Value: 0x0000000000010000,
				Expected:     []byte{0x00,0x00,0x01,0x00,0x00,0x00,0x00,0x00},
			},
			{
				Value: 0x0000000000010001,
				Expected:     []byte{0x01,0x00,0x01,0x00,0x00,0x00,0x00,0x00},
			},
			{
				Value: 0x0000000000010002,
				Expected:     []byte{0x02,0x00,0x01,0x00,0x00,0x00,0x00,0x00},
			},
			{
				Value: 0x0000000000010003,
				Expected:     []byte{0x03,0x00,0x01,0x00,0x00,0x00,0x00,0x00},
			},

			{
				Value: 0x000000001357bd24,
				Expected:     []byte{0x24,0xbd,0x57,0x13,0x00,0x00,0x00,0x00},
			},

			{
				Value: 0x00000000ffffffff,
				Expected:     []byte{0xff,0xff,0xff,0xff,0x00,0x00,0x00,0x00},
			},

			{
				Value: 0xfedcba9876543210,
				Expected:     []byte{0x10,0x32,0x54,0x76,0x98,0xba,0xdc,0xfe},
			},

			{
					Value: 0xffffffffffffffff,
				Expected:     []byte{0xff,0xff,0xff,0xff,0xff,0xff,0xff,0xff},
			},
		}...)
	case endian.Big():
		tests = append(tests, []struct{
			Value uint64
			Expected []byte
		}{
			{
				Value: 0x0000000000000000,
				Expected:     []byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00},
			},
			{
				Value: 0x0000000000000001,
				Expected:     []byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x01},
			},
			{
				Value: 0x0000000000000002,
				Expected:     []byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x02},
			},
			{
				Value: 0x0000000000000003,
				Expected:     []byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x03},
			},
			{
				Value: 0x0000000000000004,
				Expected:     []byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x04},
			},
			{
				Value: 0x0000000000000005,
				Expected:     []byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x05},
			},
			{
				Value: 0x0000000000000006,
				Expected:     []byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x06},
			},
			{
				Value: 0x0000000000000007,
				Expected:     []byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x07},
			},
			{
				Value: 0x0000000000000008,
				Expected:     []byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x08},
			},
			{
				Value: 0x0000000000000009,
				Expected:     []byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x09},
			},
			{
				Value: 0x000000000000000a,
				Expected:     []byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x0a},
			},
			{
				Value: 0x000000000000000b,
				Expected:     []byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x0b},
			},
			{
				Value: 0x000000000000000c,
				Expected:     []byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x0c},
			},
			{
				Value: 0x000000000000000d,
				Expected:     []byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x0d},
			},
			{
				Value: 0x000000000000000e,
				Expected:     []byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x0e},
			},
			{
				Value: 0x000000000000000f,
				Expected:     []byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x0f},
			},
			{
				Value: 0x0000000000000010,
				Expected:     []byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x10},
			},
			{
				Value: 0x0000000000000011,
				Expected:     []byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x11},
			},
			{
				Value: 0x0000000000000012,
				Expected:     []byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x12},
			},
			{
				Value: 0x0000000000000013,
				Expected:     []byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x13},
			},

			{
				Value: 0x00000000000000ef,
				Expected:     []byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0xef},
			},
			{
				Value: 0x00000000000000f0,
				Expected:     []byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0xf0},
			},
			{
				Value: 0x00000000000000f1,
				Expected:     []byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0xf1},
			},
			{
				Value: 0x00000000000000f2,
				Expected:     []byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0xf2},
			},
			{
				Value: 0x00000000000000f3,
				Expected:     []byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0xf3},
			},
			{
				Value: 0x00000000000000f4,
				Expected:     []byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0xf4},
			},
			{
				Value: 0x00000000000000f5,
				Expected:     []byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0xf5},
			},
			{
				Value: 0x00000000000000f6,
				Expected:     []byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0xf6},
			},
			{
				Value: 0x00000000000000f7,
				Expected:     []byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0xf7},
			},
			{
				Value: 0x00000000000000f8,
				Expected:     []byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0xf8},
			},
			{
				Value: 0x00000000000000f9,
				Expected:     []byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0xf9},
			},
			{
				Value: 0x00000000000000fa,
				Expected:     []byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0xfa},
			},
			{
				Value: 0x00000000000000fb,
				Expected:     []byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0xfb},
			},
			{
				Value: 0x00000000000000fc,
				Expected:     []byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0xfc},
			},
			{
				Value: 0x00000000000000fd,
				Expected:     []byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0xfd},
			},
			{
				Value: 0x00000000000000fe,
				Expected:     []byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0xfe},
			},
			{
				Value: 0x00000000000000ff,
				Expected:     []byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0xff},
			},
			{
				Value: 0x0000000000000100,
				Expected:     []byte{0x00,0x00,0x00,0x00,0x00,0x00,0x01,0x00},
			},
			{
				Value: 0x0000000000000101,
				Expected:     []byte{0x00,0x00,0x00,0x00,0x00,0x00,0x01,0x01},
			},
			{
				Value: 0x0000000000000102,
				Expected:     []byte{0x00,0x00,0x00,0x00,0x00,0x00,0x01,0x02},
			},
			{
				Value: 0x0000000000000103,
				Expected:     []byte{0x00,0x00,0x00,0x00,0x00,0x00,0x01,0x03},
			},

			{
				Value: 0x0000000000000245,
				Expected:     []byte{0x00,0x00,0x00,0x00,0x00,0x00,0x02,0x45},
			},

			{
				Value: 0x00000000000037ae,
				Expected:     []byte{0x00,0x00,0x00,0x00,0x00,0x00,0x37,0xae},
			},

			{
				Value: 0x000000000000fffa,
				Expected:     []byte{0x00,0x00,0x00,0x00,0x00,0x00,0xff,0xfa},
			},
			{
				Value: 0x000000000000fffb,
				Expected:     []byte{0x00,0x00,0x00,0x00,0x00,0x00,0xff,0xfb},
			},
			{
				Value: 0x000000000000fffc,
				Expected:     []byte{0x00,0x00,0x00,0x00,0x00,0x00,0xff,0xfc},
			},
			{
				Value: 0x000000000000fffd,
				Expected:     []byte{0x00,0x00,0x00,0x00,0x00,0x00,0xff,0xfd},
			},
			{
				Value: 0x000000000000fffe,
				Expected:     []byte{0x00,0x00,0x00,0x00,0x00,0x00,0xff,0xfe},
			},
			{
				Value: 0x000000000000ffff,
				Expected:     []byte{0x00,0x00,0x00,0x00,0x00,0x00,0xff,0xff},
			},
			{
				Value: 0x0000000000010000,
				Expected:     []byte{0x00,0x00,0x00,0x00,0x00,0x01,0x00,0x00},
			},
			{
				Value: 0x0000000000010001,
				Expected:     []byte{0x00,0x00,0x00,0x00,0x00,0x01,0x00,0x01},
			},
			{
				Value: 0x0000000000010002,
				Expected:     []byte{0x00,0x00,0x00,0x00,0x00,0x01,0x00,0x02},
			},
			{
				Value: 0x0000000000010003,
				Expected:     []byte{0x00,0x00,0x00,0x00,0x00,0x01,0x00,0x03},
			},

			{
				Value: 0x000000001357bd24,
				Expected:     []byte{0x00,0x00,0x00,0x00,0x13,0x57,0xbd,0x24},
			},

			{
				Value: 0x00000000ffffffff,
				Expected:     []byte{0x00,0x00,0x00,0x00,0xff,0xff,0xff,0xff},
			},

			{
				Value: 0xfedcba9876543210,
				Expected:     []byte{0xfe,0xdc,0xba,0x98,0x76,0x54,0x32,0x10},
			},

			{
				Value: 0xffffffffffffffff,
				Expected:     []byte{0xff,0xff,0xff,0xff,0xff,0xff,0xff,0xff},
			},
		}...)
	}


	{
		randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))

		const max = 50
		for i:=0; i<max; i++ {
			x64 := randomness.Int63()

			u64 := uint64(x64)

			var b0 byte
			var b1 byte
			var b2 byte
			var b3 byte
			var b4 byte
			var b5 byte
			var b6 byte
			var b7 byte

			switch endianness {
			default:
                                panic(endianness)
			case endian.Little():
				b0 = uint8( 0x00000000000000ff & u64       )
				b1 = uint8((0x000000000000ff00 & u64) >>  8)
				b2 = uint8((0x0000000000ff0000 & u64) >> 16)
				b3 = uint8((0x00000000ff000000 & u64) >> 24)
				b4 = uint8((0x000000ff00000000 & u64) >> 32)
				b5 = uint8((0x0000ff0000000000 & u64) >> 40)
				b6 = uint8((0x00ff000000000000 & u64) >> 48)
				b7 = uint8((0xff00000000000000 & u64) >> 56)
			case endian.Big():
				b0 = uint8((0xff00000000000000 & u64) >> 56)
				b1 = uint8((0x00ff000000000000 & u64) >> 48)
				b2 = uint8((0x0000ff0000000000 & u64) >> 40)
				b3 = uint8((0x000000ff00000000 & u64) >> 32)
				b4 = uint8((0x00000000ff000000 & u64) >> 24)
				b5 = uint8((0x0000000000ff0000 & u64) >> 16)
				b6 = uint8((0x000000000000ff00 & u64) >>  8)
				b7 = uint8( 0x00000000000000ff & u64       )
			}

			test := struct{
				Value uint64
				Expected []byte
			}{
				Value:           u64,
				Expected: []byte{b0, b1, b2, b3, b4, b5, b6, b7},
			}

			tests = append(tests, test)
		}
	}


	for testNumber, test := range tests {

		var buffer bytes.Buffer

		n64, err := nativeendian.WriteUint64To(&buffer, test.Value)
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
