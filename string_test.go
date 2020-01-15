package zutil

import (
	"fmt"
	"testing"
)

func TestMapToString(t *testing.T) {
	val:=make(map[string]interface{})
	val["int"]=12
	val["str"]="str"
	fmt.Println(MapToString(val))
}
