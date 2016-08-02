package middleware

import (
	"github.com/emicklei/go-restful"
	"github.com/lexandro/go-assert"
	"net/http"
	"testing"
	"fmt"
	"sync"
)

func Test_MetricsLogger(t *testing.T) {
	// given
	restful.DefaultContainer = restful.NewContainer()
	req, err := http.NewRequest("GET", "/test/metrics", nil)
	req.RemoteAddr = "111.111.111.111"
	req.Header.Add("User-Agent", "fakeAgent")
	req.Header.Add("Referer", "fakeReferer")
	if err != nil {
		return
	}
	ws := new(restful.WebService)
	ws.Path("/test").Route(ws.GET("/metrics").To(DummyHandleFunc))

	restful.Filter(ApiMetricsFilter)
	restful.Add(ws)
	ml := &MockLogger{
	}
	restful.SetLogger(ml)
	// when in async way
	wg := &sync.WaitGroup{}
	calls := 4
	wg.Add(calls)
	for i := 0; i < calls; i++ {
		go SendAsyncRequest(wg, req, nil)

	}
	wg.Wait()
	// then
	fmt.Printf("%d\n", ml.Calls)
	assert.Equals(t, 1, ml.Calls)
	assert.Equals(t, "127.456.789.012 - - [21/Jul/2016 10:49:32 +0000] \"GET /test/logger HTTP/1.1\" 200 0 \"fakeReferer\" \"fakeAgent\"\n", ml.LastEntry)
}
