package roundrobin

import (
	"errors"

	"github.com/hlts2/lock-free"
)

// ErrServersNotExists is the error that servers dose not exists
var ErrServersNotExists = errors.New("servers dose not exist")

// Servers is custom type of servers
type Servers []string

// RoundRobin returns round-robin closure
func RoundRobin(servers Servers) (func() string, error) {
	if len(servers) == 0 {
		return nil, ErrServersNotExists
	}

	lf := lockfree.New()

	idx := 0

	var server string

	return func() string {
		lf.Wait()

		if idx >= len(servers) {
			idx = 0
		}

		server = servers[idx]

		idx++

		// I do not use defer, decause defer is slow.
		lf.Signal()
		return server
	}, nil
}