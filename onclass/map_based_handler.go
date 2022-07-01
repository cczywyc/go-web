package onclass

import "C"
import (
	"net/http"
)

type Routable interface {
	Route(method string, pattern string, handleFunc func(ctx *Context))
}

type Handler interface {
	ServerHTTP(c *Context)
	Routable
}

type HandlerBasedOnMap struct {
	// key: method + url
	handlers map[string]func(ctx *Context)
}

func (s *HandlerBasedOnMap) Route(method string, pattern string, handleFunc func(ctx *Context)) {
	key := s.key(method, pattern)
	s.handlers[key] = handleFunc
}

func (h *HandlerBasedOnMap) ServerHTTP(c *Context) {
	key := h.key(c.R.Method, C.R.URL.Path)
	if handler, ok := h.handlers[key]; ok {
		handler(c)
	} else {
		c.W.WriteHeader(http.StatusNotFound)
		c.W.Write([]byte("Not Found"))
	}
}

func (h *HandlerBasedOnMap) key(method string, pattern string) string {
	return method + "#" + pattern
}

var _ Handler = &HandlerBasedOnMap{}

func NewHandlerBasedOnMap() Handler {
	return &HandlerBasedOnMap{
		handlers: make(map[string]func(ctx *Context)),
	}
}
