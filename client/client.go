package client

import (
	"github.com/ghetzel/bee-hotel"
	"github.com/ghetzel/plexutil"
)

type PlexClient struct {
	*bee.MultiClient
}

func New(address string) *PlexClient {
	mc := bee.NewMultiClient(address)

	return &PlexClient{
		MultiClient: mc,
	}
}

func NewFromConfig(config plexutil.Configuration) *PlexClient {
	client := New(config.URL)

	return client
}

func (self *PlexClient) Address() string {
	if len(self.Addresses) > 0 {
		return self.Addresses[0]
	}

	return ``
}

func (self *PlexClient) Initialize() error {
	return nil
}

func (self *PlexClient) GetStatus() (MediaContainer, error) {
	status := MediaContainer{}

	if _, err := self.Request(`GET`, `/`, nil, &status, nil); err == nil {
		return status, nil
	} else {
		return status, err
	}
}

func (self *PlexClient) CurrentSessions() ([]Video, error) {
	sessions := MediaContainer{}

	if _, err := self.Request(`GET`, `/status/sessions`, nil, &sessions, nil); err == nil {
		return sessions.Videos, nil
	} else {
		return []Video{}, err
	}
}
