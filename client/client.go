package client

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/ghetzel/go-stockutil/httputil"
	"github.com/ghetzel/plexutil"
)

type PlexClient struct {
	*httputil.Client
	IgnoreSSL bool
	Timeout   time.Duration
	inited    bool
}

func New(address string) *PlexClient {
	var client = &PlexClient{
		Client:  httputil.MustClient(address),
		Timeout: 5 * time.Second,
	}

	client.SetPreRequestHook(func(_ *http.Request) (interface{}, error) {
		if client.inited {
			return nil, nil
		} else {
			return nil, client.init()
		}
	})

	return client
}

func NewFromConfig(config plexutil.Configuration) *PlexClient {
	var client = New(config.URL)

	for k, v := range config.Parameters {
		client.SetParam(k, v)
	}

	for k, v := range config.Headers {
		client.SetHeader(k, v)
	}

	return client
}

func (self *PlexClient) Address() string {
	return self.URI().String()
}

func (self *PlexClient) init() error {
	self.SetHeader(`X-Plex-Product`, plexutil.Name)
	self.SetHeader(`X-Plex-Version`, plexutil.Version)
	self.SetInsecureTLS(self.IgnoreSSL)
	self.inited = true

	return nil
}

func (self *PlexClient) GetStatus() (MediaContainer, error) {
	var status MediaContainer

	if res, err := self.Get(`/`, nil, nil); err == nil {
		if err := self.Decode(res.Body, &status); err == nil {
			return status, nil
		} else {
			return status, err
		}
	} else {
		return status, err
	}
}

func (self *PlexClient) CurrentSessions() ([]Video, error) {
	var sessions MediaContainer

	if res, err := self.Get(`/status/sessions`, nil, nil); err == nil {
		if err := self.Decode(res.Body, &sessions); err == nil {
			return sessions.Videos, nil
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}

func (self *PlexClient) RecentSessions(perPage int, pageNum int) ([]Video, error) {
	var sessions MediaContainer

	if res, err := self.Get(
		`/status/sessions/history/all`,
		map[string]interface{}{
			`X-Plex-Container-Size`:  perPage,
			`X-Plex-Container-Start`: (perPage * (pageNum - 1)),
		},
		nil,
	); err == nil {
		if err := self.Decode(res.Body, &sessions); err == nil {
			return sessions.Videos, nil
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}

func (self *PlexClient) ListDirectory(area string, path []string) (MediaContainer, error) {
	var results MediaContainer
	var qs = make(map[string]interface{})

	for i, _ := range path {
		if path[i] == `folder` && len(path) > (i+1) {
			qs[`parent`] = path[i+1]
			path = append(path[:i+1], path[i+2:]...)
			break
		}
	}

	if res, err := self.Get(
		fmt.Sprintf("/%s/%s", area, strings.Join(path, `/`)),
		qs,
		nil,
	); err == nil {
		if err := self.Decode(res.Body, &results); err == nil {
			// post-process directories to provide logical path components without
			// dealing with querystrings
			for i, directory := range results.Directories {
				if strings.Contains(directory.Key, `?parent=`) {
					parts := strings.Split(directory.Key, `?`)
					parentKV := strings.SplitN(parts[1], `=`, 2)

					if len(parentKV) == 2 {
						results.Directories[i].PathKey = parentKV[1]
					}
				} else {
					results.Directories[i].PathKey = results.Directories[i].Key
				}
			}

			return results, nil
		} else {
			return results, nil
		}
	} else {
		return results, fmt.Errorf("request: %v", err)
	}
}

func (self *PlexClient) GetMetadata(ratingKey int) (Video, error) {
	var results MediaContainer

	if res, err := self.Get(
		fmt.Sprintf("/library/metadata/%d", ratingKey),
		nil,
		nil,
	); err == nil {
		if err := self.Decode(res.Body, &results); err != nil {
			return Video{}, err
		}

		if l := len(results.Videos); l == 1 {
			video := results.Videos[0]
			video.LibrarySectionID = results.LibrarySectionID
			video.LibrarySectionTitle = results.LibrarySectionTitle

			return video, nil
		} else {
			return Video{}, fmt.Errorf("Too many results; expected: 1, got: %d", l)
		}
	} else {
		return Video{}, err
	}
}
