package recws

import (
	"sync"
	"time"
)

type keepAliveResponse struct {
	allowDataResponse bool
	lastResponse      time.Time
	sync.RWMutex
}

func (k *keepAliveResponse) setLastResponse() {
	k.Lock()
	defer k.Unlock()

	k.lastResponse = time.Now()
}

func (k *keepAliveResponse) getAllowDataResponse() bool {
	k.RLock()
	allow := k.allowDataResponse
	k.RUnlock()
	return allow
}

func (k *keepAliveResponse) setLastDataResponse() {
	allow := k.getAllowDataResponse()
	if allow {
		k.Lock()
		k.lastResponse = time.Now()
		k.Unlock()
	}
}

func (k *keepAliveResponse) getLastResponse() time.Time {
	k.RLock()
	defer k.RUnlock()

	return k.lastResponse
}
