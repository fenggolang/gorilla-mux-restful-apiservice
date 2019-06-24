package health

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheckHandle(t *testing.T) {
	// 创建一个请求传递给我们的处理程序。我们现在没用任何查询参数，所以我们会传递"nil"作为第三个参数
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}
	// 我们创建一个ResponseRecorder(它满足http.ResponseWriter)来记录响应。
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheckHandle)

	// 我们的处理程序满足http.Handler,所以我们可以调用他们的ServeHTTP方法直接传入我们的Request和ResponseRecorder
	handler.ServeHTTP(rr, req)

	// 检查状态码是否是我们期望的
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler 返回了错误的状态码: 获得 %v,期望是 %v", status, http.StatusOK)
	}

	// 检查响应body是否是我们期望的
	expected := `{"alive":true}`
	if rr.Body.String() != expected {
		t.Errorf("handler 返回了不期望的body: 获得 %v,期望是 %v", rr.Body.String(), expected)
	}
	// 在我们的路线有变量的情况下，我们可以在请求中传递这些变量。我们可以编写表驱动测试来根据需要测试多个可能的路由变量。
	// 略。。。
}
