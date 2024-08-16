package json

import (
	"fmt"
	"testing"
)

type TestStruct struct {
	BoolStr    bool    `json:""`
	IntStr     int64   `json:""`
	OmitEncode string  `json:",omit_decode"`
	UintptrStr uintptr `json:""`
	StrStr     string  `json:""`
}

func TestOmitEncode(t *testing.T) {

	var x = TestStruct{
		BoolStr:    true,
		IntStr:     66,
		OmitEncode: "must be omitted",
		UintptrStr: 56,
		StrStr:     "dev",
	}
	marshal, _ := Marshal(x)
	fmt.Println(string(marshal))
	var y TestStruct
	err := Unmarshal(marshal, &y)
	fmt.Printf("%+v", y)
	fmt.Println("error:", err)
}
