package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"syscall"
	"time"
	"unsafe"
)

func downLoad(url string) {
	req1, _ := http.NewRequest("GET", url, nil)
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Timeout: 30 * time.Second,
	}
	resp1, err := client.Do(req1)
	// 获得跳转以后的URL, 图片网址
	picURL := resp1.Header.Get("location")
	if err != nil {
		fmt.Printf("err: %s\n", err)
	}
	defer resp1.Body.Close()

	picFileName := picName(picURL)
	fmt.Printf("%v\n", picFileName)

	// 获得图片网址的response
	req2, _ := http.NewRequest("GET", url, nil)
	client2 := http.Client{}
	resp2, err := client2.Do(req2)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	defer resp2.Body.Close()
	// 存入文件
	if !isExist("D:/wallPapers") {
		os.MkdirAll("D:/wallPapers", 0777)
	}
	picPath := "D:/wallPapers/" + picFileName + ".jpg"
	localFile, _ := os.OpenFile(picPath, os.O_CREATE|os.O_RDWR, 0777)
	if _, err := io.Copy(localFile, resp2.Body); err != nil {
		panic("failed save " + picFileName)
	}
	localFile.Close()
	setWindowsWallpaper(picPath)
}

// 调用user32.dll 里的 API， 设置桌面背景。
func setWindowsWallpaper(imagePath string) error {
	dll, _ := syscall.LoadDLL("user32.dll")
	proc, _ := dll.FindProc("SystemParametersInfoW")
	defer syscall.FreeLibrary(dll.Handle)
	ptr, _ := syscall.UTF16PtrFromString(imagePath)
	ret, _, _ := proc.Call(uintptr(20), uintptr(0), uintptr(unsafe.Pointer(ptr)), uintptr(3))
	if ret != 1 {
		return errors.New("系统调用失败")
	}
	return nil
}

// 判断文件夹是否存在
func isExist(dir string) bool {
	_, err := os.Stat(dir)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

// 获取 保存用的图片名
func picName(s string) string {
	arr := strings.Split(s, ".")
	return arr[3]
}
func main() {
	url := "https://area.sinaapp.com/bingImg/"
	downLoad(url)
}
