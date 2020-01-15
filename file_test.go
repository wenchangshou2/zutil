package zutil

import (
	"fmt"
	"testing"
)

func TestGetWorkDirFileList(t *testing.T) {
	files,err:=GetWorkDirFileList("d:/Resource/videoPlayer")
	fmt.Println("111",files,err)
}