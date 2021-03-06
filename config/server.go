package config

import (
	"net/url"
)

//Server TODO: description
type Server interface {
	GetURL() string
	GetFlags() []string
}

type server struct {
	url           string
	disableParams []string
}

func createServer(address string, params []string) (server, error) {
	if _, err := url.ParseRequestURI(address); err != nil {
		return server{}, err
	}
	return server{
		url:           address,
		disableParams: params,
	}, nil
}

func (s server) GetURL() string {
	return s.url
}

//TODO: provide test
func (s server) GetFlags() []string {
	return s.disableParams
}
