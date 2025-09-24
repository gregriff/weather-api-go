// package shared provides structs and configurations shared between the http.Client's used for the mapbox and nws services of this API.
package shared

import (
	"fmt"
	"net/http"
	"strings"
)

// Transport allows custom attributes to be added to each HTTP request sent by an http.Client that uses this transport
type Transport struct {
	BaseURL string
	Headers map[string]string
}

// RoundTrip adds upon the normal http.Transport.RoundTrip() behavior to add headers and a base url to each request.
// Reference: https://cs.opensource.google/go/x/oauth2/+/refs/tags/v0.31.0:transport.go
func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	url := req.URL.String()
	if !strings.HasPrefix(url, "http") {
		baseURL := strings.TrimSuffix(t.BaseURL, "/")
		path := "/" + strings.TrimPrefix(url, "/")
		newURL, err := req.URL.Parse(baseURL + path)
		if err != nil {
			panic(fmt.Errorf("URL PARSE ERROR: %v", err))
		}
		req.URL = newURL
	}

	for k, v := range t.Headers {
		req.Header.Add(k, v)
	}
	return http.DefaultTransport.RoundTrip(req)
}
