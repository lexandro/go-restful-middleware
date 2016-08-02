package middleware

import (
	"fmt"
	"github.com/emicklei/go-restful"
	"net/http"
	"sync"
)

type MockLogger struct {
	Calls     int
	LastEntry string
}

func SendAsyncRequest(wg *sync.WaitGroup, req *http.Request, rw http.ResponseWriter) (err error) {
	err = SendRequest(req, rw)
	defer wg.Done()
	return
}
func SendRequest(req *http.Request, rw http.ResponseWriter) (err error) {
	// handle panic
	defer func() {
		switch t := recover().(type) {
		case string:
			err = fmt.Errorf(t)
		case error:
			err = t
		}
	}()
	restful.DefaultContainer.ServeHTTP(rw, req)
	return
}

func DummyHandleFunc(_ *restful.Request, _ *restful.Response) {

}

func (ml *MockLogger) Print(v ...interface{}) {
	ml.Calls++

}
func (ml *MockLogger) Printf(format string, v ...interface{}) {
	ml.Calls++
	v[1] = "21/Jul/2016 10:49:32 +0000" // hardcoding time
	ml.LastEntry = fmt.Sprintf(format, v...)
}