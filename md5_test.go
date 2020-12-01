package zutil

import "testing"

func TestGenerateStringMd5(t *testing.T) {
	str:=GenerateStringMd5("wenchangshou")

	if str!="a8cba0500bec4ec529ac3f8de86db392"{
		t.Errorf("md5 计算错误,期望:%s,实际:%s","a8cba0500bec4ec529ac3f8de86db392",str)
	}
}