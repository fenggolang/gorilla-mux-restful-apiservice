package health

/**
* @author wangpengcheng@ccssoft.com.cn
* @date 2018/6/18 22:11
 */
import (
	"io"
	"net/http"
)

/**
* @author wangpengcheng@ccssoft.com.cn
* @date 2018/6/18 22:05
 */

/**
在Go Web应用程序中测试处理程序非常简单，多路复用器不会使得这个问题更复杂。
*/

// 一个简单的健康检查handler
func HealthCheckHandle(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	// 我们可以报告我们的数据库的状态或我们的缓存
	// 例如Redis，执行一个简单的PING,并将它们包含在响应中
	io.WriteString(w, `{"alive": true}`) // 浏览器页面输出：{"alive": true}
}
