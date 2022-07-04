package onclass

import (
	"net/http"
	"strings"
)

type HandlerBasedOnTree struct {
	root *node
}

type node struct {
	path     string
	children []*node

	handler handlerFunc
}

func (h *HandlerBasedOnTree) ServerHTTP(c *Context) {
	handler, found := h.findRouter(c.R.URL.Path)
	if !found {
		c.W.WriteHeader(http.StatusNotFound)
		_, _ = c.W.Write([]byte("Not Found"))
		return
	}
	handler(c)
}

func (h *HandlerBasedOnTree) findRouter(path string) (handlerFunc, bool) {
	paths := strings.Split(strings.Trim(path, "/"), "/")
	cur := h.root
	for _, p := range paths {
		matchChild, found := h.findMatchChild(cur, p)
		if !found {
			return nil, false
		}
		cur = matchChild
	}
	if cur.handler == nil {
		return nil, false
	}
	return cur.handler, true
}

func (h *HandlerBasedOnTree) Route(method string, pattern string, handleFunc handlerFunc) {
	pattern = strings.Trim(pattern, "/")
	paths := strings.Split(pattern, "/")
	cur := h.root
	for index, path := range paths {
		mathChild, ok := h.findMatchChild(cur, path)
		if ok {
			cur = mathChild
		} else {
			h.createSubTree(cur, paths[index:], handleFunc)
			return
		}
	}
	cur.handler = handleFunc
}

func (h *HandlerBasedOnTree) createSubTree(root *node, paths []string, handlerFn handlerFunc) {
	cur := root
	for _, path := range paths {
		nn := newNode(path)
		cur.children = append(cur.children, nn)
		cur = nn
	}
	cur.handler = handlerFn
}

func newNode(path string) *node {
	return &node{
		path:     path,
		children: make([]*node, 0, 2),
	}
}

func (h *HandlerBasedOnTree) findMatchChild(root *node, path string) (*node, bool) {
	for _, child := range root.children {
		if child.path == path || child.path == "*" {
			return child, true
		}
	}
	return nil, false
}
