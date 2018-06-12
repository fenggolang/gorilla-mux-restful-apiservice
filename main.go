package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"io/ioutil"
	"time"
)

/**
* @author wangpengcheng@ccssoft.com.cn
* @date 2018/6/11 23:30
*/
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

	// GET routes
	{
		// curl -XGET http://172.40.4.130:8080/api/service/get?servicename=wpc
		// {"name":"wpc","result":"succ"}
		router.HandleFunc("/api/service/get",MyGetHandler2).Methods("GET")
	}

	// POST routes
	{
		// curl -XPOST -d '{"servicetype":"wpctype"}' http://172.40.4.130:8080/api/service/wpcservicename/post
		// {"name":"wpcservicename","result":"succ","type":"wpctype"}
		router.HandleFunc("/api/service/{servicename}/post",MyPostHandler).Methods("POST")
	}

	// 启动rest server监听
	err:=http.ListenAndServe(address,router)
	if err!=nil{
		log.Fatalln("ListenAndServe err:",err)
	}
	log.Println("Server end")
}

// handler 处理函数:GET请求
// curl 172.40.4.130:8080/api/service/get?name=wangpengcheng
// {"name":"wpc","result":"succ"}
func MyGetHandler(w http.ResponseWriter,req *http.Request){
	// 如果路径请求错误需要返回状态码并返回api请求错误原因
	res := make(map[string]string)
	status := http.StatusOK

	// 解析查询参数
	vals:=req.URL.Query()
	//fmt.Println("req.URL.Query()=",vals)
	param,ok:=vals["name"] // 获得查询参数,map[string][]string
	if !ok {
		res["result"] = "fail"
		res["error"] = "required parameter name is missing"
		status = http.StatusBadRequest
	}else{
		res["result"] = "succ"
		res["name"] = param[0]
		status = http.StatusOK
	}
	// 组合响应Body
	response,_:=json.Marshal(res)
	w.WriteHeader(status)
	w.Header().Set("Content-Type","application/json")
	w.Write(response)
}

/**
  上面的测试中我们明白：w http.ResponseWriter可以用来设置返回信息的Status
 值，Header信息，以及Body内容，如果这个函数什么都不做如下
 */
func MyGetHandler2(w http.ResponseWriter,req *http.Request){
	// do nothing
	// 看到client正常的收到了200 OK的返回码

	// 让handler休眠5秒，然后postman客户端请求这个server,发现客户端响应也是需要等待5秒，说明handler是一个同步调用函数
	time.Sleep(5*time.Second)

	// handler函数是单线程的还是多线程的
	// :不同的handler可以同时进来，因为不同的handler的执行方式是多线程的,多个请求可以冲入
	// 同一个handler也是可以重入的。
}
/**
 client发送POST消息，并使用path变量，同时附带JSON格式body消息体，
 server端解析body内容，并返回JSON信息。
 */
 func MyPostHandler(w http.ResponseWriter, req *http.Request){
 	// 解析path变量
 	vars:=mux.Vars(req)
 	servicename:=vars["servicename"]

 	// 解析JSON Body
 	var r map[string]interface{}
 	body,_:=ioutil.ReadAll(req.Body)
 	json.Unmarshal(body,&r)
 	servicetype:=r["servicetype"].(string)

 	// 组装响应body
 	var res = map[string]string{"result":"succ","name":servicename,"type":servicetype}
 	response,_:=json.Marshal(res)
 	w.Header().Set("Content-Type","application/json")
 	w.Write(response)
 }

 // 测试依次访问下面的2个handler，发现MyGetHandler3先得到相应但是后退出，MyPostHandler3后得到相应先退出，说明不同的Handler是多线程执行
 func MyGetHandler3(w http.ResponseWriter, req *http.Request){
 	log.Println("进入MyGetHandler3")
 	time.Sleep(10*time.Second)
 	log.Println("退出MyGetHandler3")
 }
 func MyPostHandler3(w http.ResponseWriter, req *http.Request){
 	log.Println("进入MyPostHandler3")
 	time.Sleep(2*time.Second)
 	log.Println("退出MyPostHandler3")
 }