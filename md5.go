package zutil

import (
	"crypto/md5"
	"io"
	"os"
)
func GeneratorMd5(p string)([]byte,error){
	f,err:=os.Open(p)
	if err!=nil{
		return nil,err
	}

	defer f.Close()
	md5hash:=md5.New()
	if _,err:=io.Copy(md5hash,f);err!=nil{
		return nil,err
	}
	return md5hash.Sum(nil),nil
}