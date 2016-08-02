// All based on: https://gist.github.com/Tantas/1fc00c5eb7c291e2a34b
//
package middleware

import (
	"github.com/emicklei/go-restful"
	"net/http"
	"time"
	"github.com/emicklei/go-restful/log"
)

// Based on: https://httpd.apache.org/docs/2.4/logs.html#accesslog
const apacheFormatPattern = "%s - - [%s] \"%s %s %s\" %d %d \"%s\" \"%s\"\n"

type ApacheLogRecord struct {
	ip                    string
	time                  string
	method, uri, protocol string
	status                int
	responseBytes         int
	referer               string
	userAgent             string
}

func ApiLoggerFilter(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	//
	chain.ProcessFilter(req, resp)
	r := req.Request
	lr := createLogRecordFrom(r, resp)
	log.Logger.Printf(apacheFormatPattern, lr.ip, lr.time, lr.method, lr.uri, lr.protocol, lr.status, lr.responseBytes, lr.referer, lr.userAgent)
}

func createLogRecordFrom(req *http.Request, resp *restful.Response) (result *ApacheLogRecord) {
	//127.0.0.1 - frank [10/Oct/2000:13:55:36 -0700] "GET /apache_pb.gif HTTP/1.0" 200 2326 "http://www.example.com/start.html" "Mozilla/4.08 [en] (Win98; I ;Nav)"
	//
	finishTime := time.Now()
	timeFormatted := finishTime.UTC().Format("02/Jan/2006 03:04:05 -0700")
	//
	result = &ApacheLogRecord{
		ip:             nvl(req.RemoteAddr),
		time:          timeFormatted,
		method:        nvl(req.Method),
		uri:           nvl(req.URL.Path),
		protocol:      nvl(req.Proto),
		status:        resp.StatusCode(),
		responseBytes: resp.ContentLength(),
		referer:       nvl(req.Referer()),
		userAgent:     nvl(req.UserAgent()),
	}

	return result
}

func nvl(content string) string {
	if len(content) > 0 {
		return content
	} else {
		return "-"
	}
}