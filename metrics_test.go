package middleware

import (
	"github.com/emicklei/go-restful"
	"github.com/lexandro/go-assert"
	"net/http"
	"testing"
	"sync"
	"time"
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
	// we should wait for a while to make sure all metrics are received and processed
	time.Sleep(5 * time.Millisecond)
	// then
	assert.Equals(t, 1, len(ApiMetrics))
	ep := Endpoint{
		Url:"/test/metrics",
		Method:"GET",
	}
	metric := ApiMetrics[ep]
	assert.Equals(t, int64(calls), *metric.NumberOfCalls)
}
