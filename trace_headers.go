package traceheaders

import (
	"context"
	"net/http"
)

const requestId = "x-request-id"
const traceId = "x-b3-traceid"
const spanId = "x-b3-spanid"
const parentSpanId = "x-b3-parentspanid"
const flags = "x-b3-flags"
const sampled = "x-b3-sampled"

type TraceHeaders struct {
	RequestId    string
	TraceId      string
	SpanId       string
	ParentSpanId string
	Sampled      string
	Flags        string
}

func (h TraceHeaders) Inject(headers http.Header) {
	if h.RequestId != "" {
		headers.Set(requestId, h.RequestId)
	}
	if h.TraceId != "" {
		headers.Set(traceId, h.TraceId)
	}
	if h.SpanId != "" {
		headers.Set(spanId, h.SpanId)
	}
	if h.ParentSpanId != "" {
		headers.Set(parentSpanId, h.ParentSpanId)
	}
	if h.Sampled != "" {
		headers.Set(sampled, h.Sampled)
	}
	if h.Flags != "" {
		headers.Set(flags, h.Flags)
	}
}

type keyType int

const key keyType = 0

func FromRequest(req *http.Request) TraceHeaders {
	return TraceHeaders{
		RequestId:    req.Header.Get(requestId),
		TraceId:      req.Header.Get(traceId),
		SpanId:       req.Header.Get(spanId),
		ParentSpanId: req.Header.Get(parentSpanId),
		Sampled:      req.Header.Get(sampled),
		Flags:        req.Header.Get(flags),
	}
}

func NewContext(ctx context.Context, headers TraceHeaders) context.Context {
	return context.WithValue(ctx, key, headers)
}

func FromContext(ctx context.Context) (TraceHeaders, bool) {
	headers, ok := ctx.Value(key).(TraceHeaders)
	return headers, ok
}
