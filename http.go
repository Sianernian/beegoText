package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	/*
		net / http 包

		127.0.0.1  特指当前编程所在的电脑本机
	*/
	//err :=http.ListenAndServe("127.0.0.1：8080", nil)
	//程序执行时， 该程序会在本地计算机的8080端口，一直监听 ， 等待浏览器的链接
	//if err !=nil{
	//	fmt.Println(err.Error())
	//}
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("天王盖地虎，宝giti塔镇河妖"))
	})
	err := http.ListenAndServe("127.0.0.1:8080",nil)
	if err !=nil{
		fmt.Println(err.Error())
	}

	http.HandleFunc("/wuhan/jiayou",jiayou)
	err1 :=http.ListenAndServe("8080",nil)
	if err1 !=nil{
		log.Fatal(err1.Error())
	}
}

func jiayou(w http.ResponseWriter,r *http.Request){
	fmt.Println("武汉加油")
	fmt.Fprintf(w,"武汉加油")
}
