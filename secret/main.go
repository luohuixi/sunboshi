package main

import (
	"log"
	"net/http"
	"io/ioutil"
	"path/filepath"
	"tool/getDecryptedPaper"
	"tool/savePaper"
)

func main() {
	// 目标根URL
	url := "http://121.43.151.190:8000/"
	// 发送 GET 请求,返回的结果还需要进行处理才能得到你需要的结果
	response, err := http.Get(url + "paper")
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}
	defer response.Body.Close()
	paper,error1:=ioutil.ReadAll(response.Body)
	if error1!=nil{log.Fatalf("%v",error1)}
	response2,err2:=http.Get(url+"secret")
	if err2!=nil {log.Fatalf("%v",err2)}   
	defer response2.Body.Close()
	secret,error2:=ioutil.ReadAll(response2.Body)
	if error2!=nil {log.Fatalf("%v",error2)}
    article:=getDecryptedPaper.GetDecryptedPaper(string(paper), string(secret))
    path:=filepath.Join("..", "paper", "Academician Sun's papers.txt")
	savePaper.SavePaper(path,article)
}
