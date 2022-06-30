package onclass

import (
	"net/http"
)

type Routable interface {
	Route(method string, pattern string, handleFunc func(ctx *Context))
}

type Handler interface {
	http.Handler
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

func (h *HandlerBasedOnMap) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	key := h.key(request.Method, request.URL.Path)
	if handler, ok := h.handlers[key]; ok {
		handler(NewContext(writer, request))
	} else {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("Not Found"))
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
