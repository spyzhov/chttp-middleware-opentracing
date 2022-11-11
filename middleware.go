package chttp_middleware_opentracing

import (
	"net/http"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

// OpentracingCustom is a chttp.Middleware constructor to add opentracing to the request.
func OpentracingCustom(getName func(*http.Request) string) func(request *http.Request, next func(*http.Request) (*http.Response, error)) (*http.Response, error) {
	if getName == nil {
		getName = defaultNameGetter
	}
	return func(request *http.Request, next func(request *http.Request) (*http.Response, error)) (response *http.Response, err error) {
		if opentracing.IsGlobalTracerRegistered() {
			span, ctx := opentracing.StartSpanFromContext(request.Context(), getName(request))
			request = request.WithContext(ctx)

			defer func(span opentracing.Span) {
				span.LogFields(
					log.String("method", request.Method),
					log.String("host", request.Host),
					log.String("path", request.URL.Path),
				)
				if response != nil {
					span.LogFields(
						log.Int("status_code", response.StatusCode),
						log.Int64("content_length", response.ContentLength),
					)
				}
				if err != nil {
					span.LogFields(log.Error(err))
				}
				span.Finish()
			}(span)

			_ = opentracing.GlobalTracer().
				Inject(span.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(request.Header))
		}
		return next(request)
	}
}
