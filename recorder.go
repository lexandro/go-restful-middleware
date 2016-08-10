//
package middleware

import (
	"github.com/emicklei/go-restful"
)

func ApiRecorderFilter(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	//
	chain.ProcessFilter(req, resp)
}
