package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/**
* @author wangpengcheng@ccssoft.com.cn
* @date 2018/6/11 23:30
*/
// https://www.jianshu.com/p/5973d1999f5d
var (
	hostname string
	port int
)
/* 注册命令行选项 */
func init(){
	flag.StringVar(&hostname,"hostname","0.0.0.0","指定的主机名或者IP在rest server启动后将会监听")
	flag.IntVar(&port,"port",8080,"rest server将会监听的端口")
}

func main() {
	flag.Parse()
	var address = fmt.Sprintf("%s:%d",hostname,port)
	log.Println("REST server正在监听",address)

	// 注册router
	router:=mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/service/get",MyGetHandler).Methods("GET")

	// 启动rest server监听
	err:=http.ListenAndServe(address,router)
	if err!=nil{
		log.Fatalln("ListenAndServe err:",err)
	}
	log.Println("Server end")
}

// handler 处理函数:GET请求
// curl 172.40.4.130:8080/api/service/get?servicename=wpc
// {"name":"wpc","result":"succ"}
func MyGetHandler(w http.ResponseWriter,req *http.Request){
	// 解析查询参数
	vals:=req.URL.Query()
	param,_:=vals["servicename"] // 获得查询参数,map[string][]string

	// 组合响应Body
	var res = map[string]string{"result":"succ","name":param[0]}
	response,_:=json.Marshal(res)
	w.Header().Set("Content-Type","application/json")
	w.Write(response)
}