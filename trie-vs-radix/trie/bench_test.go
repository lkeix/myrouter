package trie

import (
	"fmt"
	"testing"
)

type Router struct {
	root *node
}

func (r *Router) Insert(path string, value int) {
	search := path
	now := r.root
	for {
		if len(search) == 0 {
			return
		}
		next, _ := now.findLongestCommonPrefix(search[0])
		if next == nil {
			nn := &node{
				b:        search[0],
				origin:   search,
				parent:   now,
				value:    value,
				children: make([]*node, 0),
			}
			search = search[1:]
			now.children = append(now.children, nn)
			now = nn
			continue
		}

		search = search[1:]
		now = next
	}
}

func (r *Router) Search(path string) int {
	search := path
	now := r.root
	for {
		next, _ := now.findLongestCommonPrefix(search[0])
		if next == nil {
			return -1
		}

		search = search[1:]
		if len(search) == 0 {
			return next.value
		}
		now = next
	}
}

func TestInsertAndSearch(t *testing.T) {
	r := &Router{
		root: &node{
			children: make([]*node, 0),
		},
	}

	paths := []string{
		"/health",
		"/fuga",
	}

	values := []int{
		1, 2,
	}

	for i, path := range paths {
		r.Insert(path, values[i])
	}

	fmt.Println(string(r.root.children[0].b))
	// fmt.Println(r.root.children[0].children[0])
	fmt.Println(r.Search("/health"))
	fmt.Println(r.Search("/fuga"))
}

type node struct {
	b        byte
	origin   string
	value    int
	parent   *node
	children []*node
}

func (n *node) findLongestCommonPrefix(b byte) (*node, int) {
	max := 0
	var next *node
	for _, child := range n.children {
		if child.b == b {
			return child, 1
		}
	}
	return next, max
}

func BenchmarkTrie(b *testing.B) {
	b.ResetTimer()
	paths := []string{
		"/health",
		"/fuga",
	}

	values := []int{
		1, 2,
	}

	r := &Router{
		root: &node{
			children: make([]*node, 0),
		},
	}

	for i, path := range paths {
		r.Insert(path, values[i])
	}
	for i := 0; i < b.N; i++ {
		for _, path := range paths {
			r.Search(path)
		}
	}
}
