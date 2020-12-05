package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	path := "C:Users/SolitudeRain/AppData/Local/Packages/Microsoft.Windows.ContentDeliveryManager_cw5n1h2txyewy/LocalState/Assets"
	dirList, e := ioutil.ReadDir(path)
	if e != nil {
		fmt.Println("read dir error")
		return
	}
	destPath := "D:/wallPapers/"
	for _, v := range dirList {
		CopyFile(destPath+v.Name()+".jpg", path+"/"+v.Name())
	}
}

// CopyFile 复制
func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		println(err)
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}
