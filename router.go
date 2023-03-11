package myrouter

import (
	"net/http"
)

type Node struct {
	isRoot    bool
	character byte
	prefix    string
	children  []*Node
	parent    *Node
	handlers  map[string]http.Handler
}

func newNode(parent *Node, prefix string) *Node {
	return &Node{
		prefix:   prefix,
		children: []*Node{},
		handlers: make(map[string]http.Handler),
	}
}

func (n *Node) longestCommonChild(prefix string) *Node {
	var nextChild *Node
	maxLcpIndex := 0
	for i := 0; i < len(n.children); i++ {
		lcpIndex := 0
		maxLen := len(n.children[i].prefix)
		if len(prefix) < maxLen {
			maxLen = len(prefix)
		}

		for lcpIndex < maxLen && prefix[lcpIndex] == n.children[i].prefix[lcpIndex] {
			lcpIndex++
		}

		if maxLcpIndex < lcpIndex {
			maxLcpIndex = lcpIndex
			nextChild = n.children[i]
		}
	}

	return nextChild
}

func (n *Node) RemoveChild(child *Node) {
	for i := 0; i < len(n.children); i++ {
		if n.children[i] == child {
			n.children = n.children[:i+copy(n.children[i:], n.children[i+1:])]
		}
	}
}

type Router struct {
	tree *Node
}

func NewRouter() *Router {
	return &Router{
		tree: &Node{
			isRoot:   true,
			handlers: nil,
		},
	}
}

func (r *Router) GET(endpoint string, handler http.Handler) {
	r.insert(http.MethodGet, endpoint, handler)
}

func (r *Router) insert(method, endpoint string, handler http.Handler) {
	currentNode := r.tree

	for i := 0; i < len(endpoint); i++ {
		nextNode := currentNode.longestCommonChild(endpoint)

		if nextNode == nil {
			node := newNode(currentNode, endpoint)
			currentNode.children = append(currentNode.children, node)
			currentNode = node
			break
		}

		lcpIndex := 0
		endpointLen := len(endpoint)
		nodeLen := len(nextNode.prefix)
		maxLen := endpointLen
		if nodeLen < endpointLen {
			maxLen = nodeLen
		}

		for lcpIndex < maxLen && endpoint[lcpIndex] == nextNode.prefix[lcpIndex] {
			lcpIndex++
		}

		if nodeLen == lcpIndex {
			endpoint = endpoint[lcpIndex:]
			currentNode = nextNode
			continue
		}

		parent := nextNode.parent
		// ノードをアップデートする
		node := newNode(parent, endpoint[:lcpIndex])

		nextNode.parent = node
		nextNode.prefix = nextNode.prefix[lcpIndex:]
		node.children = append(node.children, nextNode)
		currentNode.children = append(currentNode.children, node)
		currentNode.RemoveChild(nextNode)

		endpoint = endpoint[lcpIndex:]
		currentNode = node
	}

	currentNode.handlers[method] = handler
}

func (r *Router) Search(method, endpoint string) http.Handler {
	currentNode := r.tree
	lcpIndex := 0
	for {
		nextNode := currentNode.longestCommonChild(endpoint)
		if nextNode == nil {
			return nil
		}

		maxLen := len(nextNode.prefix)
		if len(endpoint) < maxLen {
			maxLen = len(endpoint)
		}

		for lcpIndex < maxLen && endpoint[lcpIndex] == nextNode.prefix[lcpIndex] {
			lcpIndex++
		}

		currentNode = nextNode
		endpoint = endpoint[lcpIndex:]
		if len(endpoint) == 0 {
			break
		}
	}

	return currentNode.handlers[method]
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	handler := r.Search(req.Method, req.URL.Path)
	if handler != nil {
		handler.ServeHTTP(w, req)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 Not Found"))
}
