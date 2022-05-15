package httplib

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/MC2BP/MicroS-Go/lib/errorlib"
	"github.com/MC2BP/MicroS-Go/lib/loglib"
)

const (
	ContentJS    = "application/javascript"
	ContentOGG   = "application/ogg"
	ContentPDF   = "application/pdf"
	ContentXHTML = "application/xhtml+xml"
	ContentJSON  = "application/json"
	ContentXML   = "application/xml"
	ContentZIP   = "application/zip"
)

type Request struct {
	Method      string
	URL         string
	ContentType string
	Headers     []Header
	Parameter   []Parameter
	Body        []byte
	Status      int
}

type Header struct {
	Key   string
	Value string
}

type Parameter struct {
	Key   string
	Value interface{}
}

func (r *Request) Send(ctx context.Context) (response []byte, err error) {
	req, err := r.getHttpRequest()
	if err != nil {
		return
	}

	req.WithContext(ctx)

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	r.Status = resp.StatusCode

	return
}

func (r *Request) SendWithoutContext() (response []byte, err error) {
	req, err := r.getHttpRequest()
	if err != nil {
		return
	}

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	response, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	r.Status = resp.StatusCode

	return
}

func (r *Request) getHttpRequest() (req *http.Request, err error) {
	req, err = http.NewRequest(r.Method, r.URL, bytes.NewBuffer(r.Body))
	if err != nil {
		return
	}

	// build url
	q := req.URL.Query()
	for _, parameter := range r.Parameter {
		q.Add(parameter.Key, fmt.Sprint(parameter.Value))
	}
	req.URL.RawQuery = q.Encode()

	for _, header := range r.Headers {
		req.Header.Add(header.Key, header.Value)
	}

	return
}

func ReadBody(r *http.Request) (body []byte, err error) {
	body, err = io.ReadAll(r.Body)
	return
}

func ParseBody(r *http.Request, body *interface{}) error {
	rawBody, err := ReadBody(r)
	if err != nil {
		return err
	}

	response := Response{
		Data: body,
	}

	err = json.Unmarshal(rawBody, &response)
	if err != nil {
		return err
	}

	return nil
}

func SendError(w http.ResponseWriter, err error) {
	loglib.Warning(err)
	response := Response{
		Success:   false,
		ErrorCode: errorlib.GetErrorCode(err),
		Message:   errorlib.GetError(err).Error(),
	}
	responseBody, err := json.Marshal(response)
	if err != nil {
		loglib.Error(err)
	}
	fmt.Fprintf(w, string(responseBody))
}

func Send(w http.ResponseWriter, body interface{}) {
	response := Response{
		Success: true,
		Data:    body,
	}
	responseBody, err := json.Marshal(response)
	if err != nil {
		loglib.Error(err)
	}
	fmt.Fprintf(w, string(responseBody))
}
