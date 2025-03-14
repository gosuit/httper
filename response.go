package httper

import (
	"net/http"
)

// Resp is a wrapper around http.Response that includes the response body as a byte slice.
type Resp struct {
	ByteBody []byte
	*http.Response
}

func newResp(resp *http.Response, body []byte) *Resp {
	return &Resp{
		body,
		resp,
	}
}
