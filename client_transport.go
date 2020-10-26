package godaddy

import (
	"net"
	"net/http"
	"sync"
	"time"
)

func newTransport(rateLimit time.Duration) http.RoundTripper {
	return &rateLimitedTransport{
		delegate: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: 10 * time.Second,
			}).DialContext,
			TLSHandshakeTimeout: 10 * time.Second,
		},
		rateLimit: rateLimit,
		throttle:  time.Now().Add(-(rateLimit)),
	}
}

// rateLimitedTransport throttles API calls to GoDaddy. It appears that
// the rate limit is 60 requests per minute, which can be throttled and
// enforced at a maximum of one request/second.
type rateLimitedTransport struct {
	delegate  http.RoundTripper
	throttle  time.Time
	rateLimit time.Duration
	sync.Mutex
}

func (t *rateLimitedTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	t.Lock()
	defer t.Unlock()

	if t.throttle.After(time.Now()) {
		delta := t.throttle.Sub(time.Now())
		time.Sleep(delta)
	}

	t.throttle = time.Now().Add(t.rateLimit)
	return t.delegate.RoundTrip(req)
}
