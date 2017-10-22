package main

import (
	"fmt"
	"os"
)

const (
	DataRoot   string = "./tmp"
	BingURL    string = "https://www.bing.com"
	QueryParam string = "HPImageArchive.aspx?format=js&idx=0&n=1&mkt=zh-cn"
)

func main() {
	body, err := getUrlBody(fmt.Sprintf("%v/%v", BingURL, QueryParam))
	if err != nil {
		fmt.Println(err)
		return
	}
	imgPart, err := getimgurl(body)
	if err != nil {
		fmt.Println(err)
		return
	}
	imgurl := BingURL + imgPart
	if !isDirExist(DataRoot) {
		os.Mkdir(DataRoot, 0755)
		fmt.Println("dir %s created", DataRoot)
	}
	if err := SaveImage(imgurl); err != nil {
		fmt.Println(err)
	}
}
