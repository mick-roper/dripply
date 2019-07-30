package targets

import (
	"io"
	"net/http"
	"net/url"
)

// SimpleTarget that can be used for a proxy
type SimpleTarget struct {
	Hostname string
	Scheme   string
}

// GetResponse using the host and scheme
func (t *SimpleTarget) GetResponse(client *http.Client, method, path string, body io.Reader) (*http.Response, error) {
	u := url.URL{
		Host:   t.Hostname,
		Scheme: t.Scheme,
		Path:   path,
	}

	req, err := http.NewRequest(method, u.String(), body)

	if err != nil {
		return nil, err
	}

	return client.Do(req)
}
