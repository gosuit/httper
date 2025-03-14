package httper

import (
	"net/http"
	"time"
)

const (
	// HTTP methods defined as constants for easy reference.
	GetMethod     method = "GET"
	PostMethod    method = "POST"
	PutMethod     method = "PUT"
	DeleteMethod  method = "DELETE"
	OptionsMethod method = "OPTIONS"
	HeadMethod    method = "HEAD"
	TraceMethod   method = "TRACE"
	ConnectMethod method = "CONNECT"

	// Content types defined as constants for easy reference.
	JsonType contentType = "application/json"
	XmlType  contentType = "application/xml"
	TextType contentType = "text/plain"
	HtmlType contentType = "text/html"
)

const (
	// HTTP status codes defined as constants for easy reference.

	StatusContinue           = 100
	StatusSwitchingProtocols = 101
	StatusProcessing         = 102
	StatusEarlyHints         = 103

	StatusOK                   = 200
	StatusCreated              = 201
	StatusAccepted             = 202
	StatusNonAuthoritativeInfo = 203
	StatusNoContent            = 204
	StatusResetContent         = 205
	StatusPartialContent       = 206
	StatusMultiStatus          = 207
	StatusAlreadyReported      = 208
	StatusIMUsed               = 226

	StatusMultipleChoices  = 300
	StatusMovedPermanently = 301
	StatusFound            = 302
	StatusSeeOther         = 303
	StatusNotModified      = 304
	StatusUseProxy         = 305

	StatusTemporaryRedirect = 307
	StatusPermanentRedirect = 308

	StatusBadRequest                   = 400
	StatusUnauthorized                 = 401
	StatusPaymentRequired              = 402
	StatusForbidden                    = 403
	StatusNotFound                     = 404
	StatusMethodNotAllowed             = 405
	StatusNotAcceptable                = 406
	StatusProxyAuthRequired            = 407
	StatusRequestTimeout               = 408
	StatusConflict                     = 409
	StatusGone                         = 410
	StatusLengthRequired               = 411
	StatusPreconditionFailed           = 412
	StatusRequestEntityTooLarge        = 413
	StatusRequestURITooLong            = 414
	StatusUnsupportedMediaType         = 415
	StatusRequestedRangeNotSatisfiable = 416
	StatusExpectationFailed            = 417
	StatusTeapot                       = 418
	StatusMisdirectedRequest           = 421
	StatusUnprocessableEntity          = 422
	StatusLocked                       = 423
	StatusFailedDependency             = 424
	StatusTooEarly                     = 425
	StatusUpgradeRequired              = 426
	StatusPreconditionRequired         = 428
	StatusTooManyRequests              = 429
	StatusRequestHeaderFieldsTooLarge  = 431
	StatusUnavailableForLegalReasons   = 451

	StatusInternalServerError           = 500
	StatusNotImplemented                = 501
	StatusBadGateway                    = 502
	StatusServiceUnavailable            = 503
	StatusGatewayTimeout                = 504
	StatusHTTPVersionNotSupported       = 505
	StatusVariantAlsoNegotiates         = 506
	StatusInsufficientStorage           = 507
	StatusLoopDetected                  = 508
	StatusNotExtended                   = 510
	StatusNetworkAuthenticationRequired = 511
)

type method string

type contentType string

// ServeMux is an alias for http.ServeMux, which is an HTTP request multiplexer.
type ServeMux = http.ServeMux

// DefaultClient is the default HTTP client.
var DefaultClient = NewClient(&ClientCfg{Prefix: "", Timeout: 5 * time.Second})

// Get performs a GET request to the specified URL with DefaultClient and returns the response.
func Get(url string) (*Resp, error) {
	return DefaultClient.Get(url)
}

// GetJson performs a GET request to the specified URL with DefaultClient and decodes the JSON response into the provided interface.
func GetJson(url string, to interface{}) (*Resp, error) {
	return DefaultClient.GetJson(url, to)
}

// Post performs a POST request with JSON data to the specified URL with DefaultClient and returns the response.
func Post(url string, data interface{}) (*Resp, error) {
	return DefaultClient.Post(url, data)
}

// PostJson performs a POST request with JSON data to the specified URL with DefaultClient and binds the response to the provided interface.
func PostJson(url string, data interface{}, to interface{}) (*Resp, error) {
	return DefaultClient.PostJson(url, data, to)
}

// Do executes a custom HTTP request with DefaultClient and returns the response.
func Do(req *Req) (*Resp, error) {
	return DefaultClient.Do(req)
}
