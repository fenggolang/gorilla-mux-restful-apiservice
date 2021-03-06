package health

import (
	"net/http"
)

/**
在Go Web应用程序中测试处理程序非常简单，多路复用器不会使得这个问题更复杂。
*/

// 一个简单的健康检查handler
func HealthCheckHandle(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	// 我们可以报告我们的数据库的状态或我们的缓存
	// 例如Redis，执行一个简单的PING,并将它们包含在响应中
	//io.WriteString(w, `{"alive": true}`) // 浏览器页面输出：{"alive": true}
	w.Write([]byte("Gorilla mux example!\n")) // Gorilla mux example!
	w.Write([]byte(`{"alive": true}`))        // 浏览器页面输出：{"alive": true}
}
