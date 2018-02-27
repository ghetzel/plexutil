package client

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"strings"

	"github.com/ghetzel/bee-hotel"
	"github.com/ghetzel/plexutil"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger(`plex`)

type PlexClient struct {
	*bee.MultiClient
	IgnoreSSL bool
}

func New(address string) *PlexClient {
	mc := bee.NewMultiClient(address)
	mc.ImmediatePreRequestHooks = []bee.ImmediatePreRequestHook{
		func(request *http.Request) error {
			log.Debugf("http: %s %s", request.Method, request.URL.String())

			for name, values := range request.Header {
				for _, value := range values {
					log.Debugf("http:  %- 24s: %s", name, value)
				}
			}

			return nil
		},
	}

	return &PlexClient{
		MultiClient: mc,
	}
}

func NewFromConfig(config plexutil.Configuration) *PlexClient {
	client := New(config.URL)

	if config.Parameters != nil {
		client.RequestQueryStrings = config.Parameters
	}

	if config.Headers != nil {
		client.RequestHeaders = config.Headers
	}

	return client
}

func (self *PlexClient) Address() string {
	if len(self.Addresses) > 0 {
		return self.Addresses[0]
	}

	return ``
}

func (self *PlexClient) Initialize() error {
	self.HeaderSet(`X-Plex-Product`, plexutil.Name)
	self.HeaderSet(`X-Plex-Version`, plexutil.Version)

	if self.IgnoreSSL {
		log.Warningf("SSL certificate verification is disabled")
	}

	self.SetClient(&http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: self.IgnoreSSL,
			},
		},
	})

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

func (self *PlexClient) RecentSessions(perPage int, pageNum int) ([]Video, error) {
	sessions := MediaContainer{}

	if _, err := self.Request(`GET`, `/status/sessions/history/all`, nil, &sessions, nil, func(request *bee.MultiClientRequest) error {
		request.QuerySet(`X-Plex-Container-Size`, perPage)
		request.QuerySet(`X-Plex-Container-Start`, (perPage * (pageNum - 1)))

		return nil

	}); err == nil {
		return sessions.Videos, nil
	} else {
		return []Video{}, err
	}
}

func (self *PlexClient) ListDirectory(area string, path []string) (MediaContainer, error) {
	results := MediaContainer{}
	qs := make(map[string]interface{})

	for i, _ := range path {
		if path[i] == `folder` && len(path) > (i+1) {
			qs[`parent`] = path[i+1]
			path = append(path[:i+1], path[i+2:]...)
			break
		}
	}

	if _, err := self.Request(`GET`, fmt.Sprintf("/%s/%s", area, strings.Join(path, `/`)), nil, &results, nil, func(request *bee.MultiClientRequest) error {
		for k, v := range qs {
			request.QuerySet(k, v)
		}

		return nil
	}); err == nil {
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
		return results, fmt.Errorf("request: %v", err)
	}
}

func (self *PlexClient) GetMetadata(ratingKey int) (Video, error) {
	results := MediaContainer{}

	if _, err := self.Request(`GET`, fmt.Sprintf("/library/metadata/%d", ratingKey), nil, &results, nil); err == nil {
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
