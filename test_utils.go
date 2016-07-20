package middleware

import (
	"fmt"
	"github.com/emicklei/go-restful"
	"net/http"
)

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

