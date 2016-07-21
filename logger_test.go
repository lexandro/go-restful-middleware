package middleware

import (
	"github.com/emicklei/go-restful"
	"github.com/lexandro/go-assert"
	"net/http"
	"testing"
	"fmt"
)

type FakeLogger struct {
	Calls     int
	LastEntry string
}

func Test_ApiLogger(t *testing.T) {
	// given
	req, err := http.NewRequest("GET", "/test/logger", nil)
	req.RemoteAddr = "127.456.789.012"
	req.Header.Add("User-Agent", "fakeAgent")
	req.Header.Add("Referer", "fakeReferer")
	if err != nil {
		return
	}
	ws := new(restful.WebService)
	ws.Path("/test").Route(ws.GET("/logger").To(dummyHandleFunc))
	restful.Filter(ApiLogger)
	restful.Add(ws)
	fl := &FakeLogger{
	}
	restful.SetLogger(fl)
	// when
	SendRequest(req, nil)
	// then
	assert.Equals(t, 1, fl.Calls)
	assert.Equals(t, "127.456.789.012 - - [21/Jul/2016 10:49:32 +0000] \"GET /test/logger HTTP/1.1\" 200 0 \"fakeReferer\" \"fakeAgent\"\n", fl.LastEntry)

}

func (fl *FakeLogger) Print(v ...interface{}) {
	fl.Calls++

}
func (fl *FakeLogger) Printf(format string, v ...interface{}) {
	fl.Calls++
	v[1] = "21/Jul/2016 10:49:32 +0000" // hardcoding time
	fl.LastEntry = fmt.Sprintf(format, v...)
}

func dummyHandleFunc(_ *restful.Request, _ *restful.Response) {
}
