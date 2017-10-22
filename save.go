package main

import (
	"os"
	"net/http"
	"io"
)


func isDirExist(path string) bool {
	p, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	} else {
		return p.IsDir()
	}
}

func SaveImage(url string)error{
	res, err := http.Get(url)
	defer res.Body.Close()
	if err != nil {
		return err
	}
	dst, err := os.Create(DataRoot +"/"+ "test.jpg")
	if err != nil {
		return err
	}
	// 写入文件
	io.Copy(dst, res.Body)
	return nil
}