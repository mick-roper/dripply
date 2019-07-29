package targets

import "sync"

// ProxyTarget that requests can be routed to
type ProxyTarget struct {
	Host   string
	Scheme string
}

// TargetCollection holds a buck of mappings of host headers to proxy targets
type TargetCollection struct {
	entries map[string]*ProxyTarget
	mux     sync.Mutex
}

// SetTarget sets/adds a target to the collection
func (t *TargetCollection) SetTarget(originalHost string, target *ProxyTarget) {
	if originalHost == "" || target == nil {
		return
	}

	t.mux.Lock()
	defer t.mux.Unlock()

	t.entries[originalHost] = target
}

// GetTarget gets a registered target from a collection
func (t *TargetCollection) GetTarget(host string) *ProxyTarget {
	if host == "" {
		return nil
	}

	t.mux.Lock()
	defer t.mux.Unlock()

	return t.entries[host]
}

// RemoveTarget removes a target from the collection
func (t *TargetCollection) RemoveTarget(host string) {
	if host == "" {
		return
	}

	t.mux.Lock()
	defer t.mux.Unlock()

	delete(t.entries, host)
}
