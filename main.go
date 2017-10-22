package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"github.com/bitly/go-simplejson"
	"os"
	"io"
)

var bingAPI string = "https://www.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1&mkt=zh-cn"
var bing string = "https://www.bing.com"

const DataRoot string  ="./tmp"

type ERRORSHOW struct {
	code int
	message string
}

func main() {
	body,err := getUrlBody(bingAPI)
	if err!= nil {
		fmt.Println(err)
		return
	}
	datas,err := simplejson.NewJson(body)
	if err!=nil{
		fmt.Println(err)
		return
	}
	imgPart,err := datas.Get("images").GetIndex(0).Get("url").String()
	if err!= nil{
		fmt.Println(err)
		return
	}
	imgurl := bing+imgPart
	if ! isDirExist(DataRoot) {
		os.Mkdir(DataRoot, 0755);
		fmt.Println("dir %s created", DataRoot)
	}
	if err := SaveImage(imgurl);err!=nil{
		fmt.Println(err)
	}
}

func getUrlBody(url string)([]byte,error){
	resp,err := http.Get(url)
	if err != nil{
		return nil,err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return nil,err
	}
	return body,nil
}

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




