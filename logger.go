// All based on: https://gist.github.com/Tantas/1fc00c5eb7c291e2a34b
//
package middleware

import (
	"time"
	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful/log"
	"net/http"
)

// https://httpd.apache.org/docs/2.2/logs.html#combined + execution time.
const apacheFormatPattern = "%s - - [%s] \"%s %s %s\" %d %d \"%s\" \"%s\" %.4f\n"

type ApacheLogRecord struct {
	http.ResponseWriter

	ip                    string
	time                  time.Time
	method, uri, protocol string
	status                int
	responseBytes         int64
	referer               string
	userAgent             string
	elapsedTime           time.Duration
}


func ApiLogger(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	startTime := time.Now()
	//
	chain.ProcessFilter(req, resp)
	//
	finishTime := time.Now()
	timeFormatted := finishTime.UTC().Format("02/Jan/2006 03:04:05")
	//

	r := req.Request

	referer := r.Referer()
	if referer == "" {
		referer = "-"
	}

	userAgent := r.UserAgent()
	if userAgent == "" {
		userAgent = "-"
	}

	log.Logger.Printf(apacheFormatPattern, r.RemoteAddr, timeFormatted, r.Method, r.RequestURI, r.Proto, resp.StatusCode(), resp.ContentLength(), referer, userAgent,
		finishTime.Sub(startTime).Seconds())

}
