package main

import (
	"github.com/bitly/go-simplejson"
	"net/http"
	"io/ioutil"
)

func getimgurl(body []byte)(string,error){
	datas,err := simplejson.NewJson(body)
	if err!=nil{
		return "false",err
	}
	imgPart,err := datas.Get("images").GetIndex(0).Get("url").String()
	if err!= nil{
		return "false",err
	}
	return imgPart,nil
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
