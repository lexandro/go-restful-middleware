package middleware

import (
	"testing"
	"github.com/emicklei/go-restful"
	"net/http/httptest"
	"net/http"
	"fmt"
)

type FakeLogger struct {

}

func Test_ApiLogger(t *testing.T) {
	resp := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/test/logger", nil)
	if err != nil {
		return
	}
	ws := new(restful.WebService)
	ws.Path("/test").Route(ws.GET("/logger").To(loggerFunc))
	restful.Filter(ApiLogger)
	restful.Add(ws)
	restful.SetLogger(&FakeLogger{})
	//
	SendRequest(req, resp)

}

func (fl *FakeLogger) Print(v ...interface{}) {
	fmt.Println("1")

}
func (fl *FakeLogger) Printf(format string, v ...interface{}) {
	fmt.Println("2")
}

func loggerFunc(_ *restful.Request, _ *restful.Response) {
}