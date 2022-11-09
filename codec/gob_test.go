package codec_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitub.com/lxygwqf9527/demo-rpc/codec"
)

type TestStruct struct {
	F1 string
	F2 int
}

func TestGob(t *testing.T) {
	should := assert.New(t)
	gobBytes, err := codec.GobEncode(&TestStruct{F1: "test_f1", F2: 10})
	if should.NoError(err) {
		fmt.Println(gobBytes)
	}
	obj := TestStruct{}
	err = codec.GobDecode(gobBytes, &obj)
	if should.NoError(err) {
		fmt.Println(obj)
	}
}
