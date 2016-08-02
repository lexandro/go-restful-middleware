//
package middleware

import (
	"github.com/emicklei/go-restful"
	"time"
	"sync/atomic"
)

type Endpoint struct {
	Url    string
	Method string
}

type EndpointMetrics struct {
	AvgDurationInNano *int64
	//
	NumberOfCalls     *int64
	AllDurationInNano *int64
	//
	StatusCodes       map[int]*int64
}

type CallMetrics struct {
	Url            string
	Method         string
	DurationInNano int64
	Status         int
}

var ApiMetrics map[Endpoint]*EndpointMetrics = make(map[Endpoint]*EndpointMetrics)

var metricsChan chan CallMetrics

func init() {
	metricsChan = make(chan CallMetrics)
	go metricsWriter(metricsChan)
}

// single writer model
func metricsWriter(metrCh chan CallMetrics) {
	for callM := range metrCh {
		ep := Endpoint{
			Url:callM.Url,
			Method:callM.Method,
		}
		epM, ok := ApiMetrics[ep]
		if !ok {
			epM = &EndpointMetrics{
				AvgDurationInNano:new(int64),
				AllDurationInNano:new(int64),
				NumberOfCalls:new(int64),
				StatusCodes: make(map[int]*int64),
			}
			ApiMetrics[ep] = epM
		}

		allDur := atomic.AddInt64(epM.AllDurationInNano, callM.DurationInNano)
		allNumCalls := atomic.AddInt64(epM.NumberOfCalls, 1)
		avgDur := int64(0)
		if allNumCalls != 0 {
			avgDur = allDur / allNumCalls
		}
		// calculated value, no need for thread safety
		epM.AvgDurationInNano = &avgDur
		//
		status := callM.Status

		statusCounter, ok := epM.StatusCodes[status]
		if !ok {
			statusCounter = new(int64)
			epM.StatusCodes[status] = statusCounter
		}
		atomic.AddInt64(statusCounter, 1)
	}
}
func ApiMetricsFilter(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	//
	start := time.Now()
	chain.ProcessFilter(req, resp)
	end := time.Now()
	//
	cm := CallMetrics{
		Url : req.SelectedRoutePath(),
		Method : req.Request.Method,
		Status:resp.StatusCode(),
		DurationInNano: end.Sub(start).Nanoseconds(),
	}
	metricsChan <- cm
}
