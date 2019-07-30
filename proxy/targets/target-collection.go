package targets

import "sync"

// TargetCollection holds a buck of mappings of host headers to proxy targets
type TargetCollection struct {
	entries map[string]Target
	mux     sync.Mutex
}

// NewTargetCollection creates a properly initialised target collection
func NewTargetCollection() *TargetCollection {
	return &TargetCollection{
		entries: make(map[string]Target),
		mux:     sync.Mutex{},
	}
}

// SetTarget sets/adds a target to the collection
func (t *TargetCollection) SetTarget(originalHost string, target Target) {
	if originalHost == "" || target == nil {
		return
	}

	t.mux.Lock()
	defer t.mux.Unlock()

	t.entries[originalHost] = target
}

// GetTarget gets a registered target from a collection
func (t *TargetCollection) GetTarget(host string) Target {
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
