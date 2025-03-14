package httper

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
)

// Req represents an HTTP request with additional parameters and flags for unmarshalling.
type Req struct {
	params        *Params
	NeedUnmarshal bool
	*http.Request
}

// Params holds the configuration for creating an HTTP request, including method, URL, body, and marshaling options.
type Params struct {
	// HTTP method
	Method method

	// URL for the request.
	Url string

	// Body of the request to be sent.
	Body interface{}

	// Predefined byte array for the body if not marshaling.
	ByteBody []byte

	// Flag indicating if the body should be marshaled.
	Marshal bool

	// Type of content for marshaling
	MarshalType contentType

	// Flag indicating if the response should be unmarshalled.
	Unmarshal bool

	// Destination for unmarshalling the response body.
	UnmarshalTo interface{}

	// Type of content for unmarshalling.
	UnmarshalType contentType
}

// NewReq creates a new Req instance based on the provided Params. It marshals the body if required.
func NewReq(params *Params) (*Req, error) {
	var body []byte
	var err error

	if params.Marshal {
		body, err = marshal(params)
		if err != nil {
			return nil, err
		}
	} else {
		body = params.ByteBody
	}

	reader := bytes.NewReader(body)

	base, err := http.NewRequest(string(params.Method), params.Url, reader)
	if err != nil {
		return nil, err
	}

	return &Req{
		params,
		params.Unmarshal,
		base,
	}, nil
}

func (r *Req) unmarshal(body []byte) error {
	switch r.params.UnmarshalType {

	case JsonType:
		err := json.Unmarshal(body, r.params.UnmarshalTo)
		if err != nil {
			return err
		}
	case XmlType:
		err := xml.Unmarshal(body, r.params.UnmarshalTo)
		if err != nil {
			return err
		}
	case TextType:
		if ptr, ok := r.params.UnmarshalTo.(*string); ok {
			*ptr = string(body)
		} else {
			return errors.New("incorrect type")
		}
	case HtmlType:
		if ptr, ok := r.params.UnmarshalTo.(*string); ok {
			*ptr = string(body)
		} else {
			return errors.New("incorrect type")
		}
	default:
		return errors.New("incorrect type")
	}

	return nil
}

func marshal(params *Params) ([]byte, error) {
	var body []byte

	switch params.MarshalType {

	case JsonType:
		enc, err := json.Marshal(params.Body)
		if err != nil {
			return nil, err
		}
		body = enc
	case XmlType:
		enc, err := xml.Marshal(params.Body)
		if err != nil {
			return nil, err
		}
		body = enc
	case TextType:
		body = []byte(fmt.Sprintf("%v", params.Body))
	case HtmlType:
		body = []byte(fmt.Sprintf("%v", params.Body))
	default:
		return nil, errors.New("incorrect type")
	}

	return body, nil
}
