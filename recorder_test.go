package middleware

import (
	"github.com/emicklei/go-restful"
	"github.com/lexandro/go-assert"
	"net/http"
	"testing"
)

func Test_ApiRecorder(t *testing.T) {
	// given
	req, err := http.NewRequest("GET", "/test/recorder", nil)
	req.RemoteAddr = "321.321.321.321"
	req.Header.Add("User-Agent", "fakeAgent")
	req.Header.Add("Referer", "fakeReferer")
	if err != nil {
		return
	}
	ws := new(restful.WebService)
	ws.Path("/test").Route(ws.GET("/recorder").To(DummyHandleFunc))
	restful.Filter(ApiRecorderFilter)
	restful.Add(ws)
	ml := &MockLogger{
	}
	restful.SetLogger(ml)
	// when
	SendRequest(req, nil)
	// then
	assert.Equals(t, 1, ml.Calls)
	assert.Equals(t, "127.456.789.012 - - [21/Jul/2016 10:49:32 +0000] \"GET /test/logger HTTP/1.1\" 200 0 \"fakeReferer\" \"fakeAgent\"\n", ml.LastEntry)

}



