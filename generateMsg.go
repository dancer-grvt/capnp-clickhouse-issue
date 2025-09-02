package main

import (
	"fmt"
	"os"

	"capnproto.org/go/capnp/v3"
	"github.com/dancer-grvt/capnp-clickhouse-issue/sampleSchema"
	"github.com/holiman/uint256"
)

func main() {
	arena := capnp.SingleSegment(nil)

	_, seg, err := capnp.NewMessage(arena)
	if err != nil {
		fmt.Println("Error when creating new capnp message!")
		panic(err)
	}

	testStruct, err := sampleSchema.NewRootTestStruct(seg)
	if err != nil {
		fmt.Println("Error when creating new struct!")
		panic(err)
	}

	testInnerTuple, err := sampleSchema.NewInnerTuple(seg)
	if err != nil {
		fmt.Println("Error when creating new inner tuple!")
		panic(err)
	}

	bigNumberField := uint256.NewInt(59)

	// init the inner tuple first
	_ = testInnerTuple.SetNormalField("Test Inner Tuple")
	testInnerTuple.SetNormalField2(32)
	_ = testInnerTuple.SetSpecialField(ToClickhouseUInt256(bigNumberField.Bytes()))

	// init the root struct
	_ = testStruct.SetTitleField("Test Struct")
	_ = testStruct.SetInnerTuple(testInnerTuple)

	outputFile, err := os.Create("./msg/sample_msg_01.bin")
	if err != nil {
		fmt.Println("Error when open output file!")
		panic(err)
	}
	defer outputFile.Close()

	encoder := capnp.NewEncoder(outputFile)
	err = encoder.Encode(testStruct.Message())
	if err != nil {
		panic(err)
	}
}

// Clickhouse need little endian while uint256 lib of golang return
// We also need to fill UInt256 with empty byte even if it's just 0
func ToClickhouseUInt256(in []byte) []byte {
	inSize := len(in)
	out := make([]byte, 32)
	for i := 0; i < inSize; i++ {
		out[i] = in[(inSize-1)-i]
	}
	return out
}
