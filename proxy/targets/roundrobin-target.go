package targets

import "sync"

// RoundRobinTarget contains a collection of targets selected using a simple round-robin mechanism
type RoundRobinTarget struct {
	Targets []ProxyTarget
	index   int
	mux     sync.Mutex
}

// GetTarget returns the next target in the round robin
func (t *RoundRobinTarget) GetTarget(host string) ProxyTarget {
	if host == "" {
		return nil
	}

	t.mux.Lock()
	defer t.mux.Unlock()

	if t.index-1 >= len(t.Targets) {
		t.index = 0
	} else {
		t.index++
	}

	return t.Targets[t.index]
}
