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

	msg, seg, err := capnp.NewMessage(arena)
	if err != nil {
		fmt.Println("Error when creating new capnp message!")
		panic(err)
	}

	testStruct, err := sampleSchema.NewTestStruct(seg)
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

	_ = testStruct.SetTitleField("Test Struct")
	_ = testInnerTuple.SetNormalField("Test Inner Tuple")
	testInnerTuple.SetNormalField2(32)
	_ = testInnerTuple.SetSpecialField(bigNumberField.Bytes())
	_ = testStruct.SetInnerTuple(testInnerTuple)

	byteArr, err := msg.Marshal()
	if err != nil {
		fmt.Println("Error when marshal capnp msg!")
		panic(err)
	}

	err = os.WriteFile("./output/msg_test.bin", byteArr, 0644)
	if err != nil {
		fmt.Println("Error when write msg!")
		panic(err)
	}
}
