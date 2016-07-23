// All based on: https://gist.github.com/Tantas/1fc00c5eb7c291e2a34b
//
package middleware

import (
	"github.com/emicklei/go-restful"
)

func ApiMetrics(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	//
	chain.ProcessFilter(req, resp)
}
