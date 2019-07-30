package targets

import (
	"io"
	"net/http"
)

// Target of a reverse proxt request
type Target interface {
	GetResponse(client *http.Client, method, path string, body io.Reader) (*http.Response, error)
}
