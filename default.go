package chttp_middleware_opentracing

import (
	"fmt"
	"net/http"
)

// Opentracing is a chttp.Middleware constructor to add opentracing to the request.
// Use the default name generator: Method + Path
func Opentracing() func(request *http.Request, next func(*http.Request) (*http.Response, error)) (*http.Response, error) {
	return OpentracingCustom(defaultNameGetter)
}

func defaultNameGetter(request *http.Request) string {
	return fmt.Sprintf("%s %s", request.Method, request.URL.Path)
}
