package ehsan

import "net/url"

type Handler = func(query *url.URL) []byte

type URI = *url.URL

type Route struct {
	Path    string
	Handler Handler
}
