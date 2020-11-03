package client

import (
	"fmt"
	"os"
	"testing"
)

func newPlexClient(doInit bool) (*PlexClient, error) {
	addr := `http://127.0.0.1:32400`

	if v := os.Getenv(`PLEX_API`); v != `` {
		addr = v
	}

	client := New(addr)

	if v := client.Address(); v == addr {
		if doInit {
			if err := client.init(); err != nil {
				return nil, err
			}
		}

		return client, nil
	} else {
		return nil, fmt.Errorf("bad address; expected: %q, got: %q", addr, v)
	}
}

func TestNewClient(t *testing.T) {
	if _, err := newPlexClient(false); err != nil {
		t.Error(err)
	}
}

func TestGetStatus(t *testing.T) {
	if client, err := newPlexClient(true); err == nil {
		if details, err := client.GetStatus(); err == nil {
			if details.MachineIdentifier == `` {
				t.Errorf("status: MachineIdentifier is empty")
			}
		} else {
			t.Error(err)
		}
	} else {
		t.Error(err)
	}
}
