package targets

// SimpleTarget contains a straighforward Host and Scheme that can be targeted by a reverse proxy
type SimpleTarget struct {
	Hostname  string
	URLScheme string
}

// Host that the request should be directed at
func (t *SimpleTarget) Host() string {
	return t.Hostname
}

// Scheme that the URL should use
func (t *SimpleTarget) Scheme() string {
	return t.URLScheme
}
