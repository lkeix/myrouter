package loop

import (
	"fmt"
	"testing"
)

const path = "/hoge/hoge/hoge"

type Router struct {
	root *node
}

func (r *Router) Insert(path string, value int) {
	search := path
	now := r.root
	for {
		next, index := now.findLongestCommonPrefix(search)
		if next == nil {
			nn := &node{
				prefix:   search,
				parent:   now,
				value:    value,
				children: make([]*node, 0),
			}
			now.children = append(now.children, nn)
			return
		}

		prefix := search[:index]
		nn := &node{
			prefix:   prefix,
			parent:   now,
			children: make([]*node, 0),
		}
		next.prefix = next.prefix[index:]
		next.parent = nn
		if now == r.root {
			now.children = make([]*node, 0)
		}
		now.children = append(now.children, nn)
		nn.children = append(nn.children, next)
		search = search[index:]
		now = nn
	}
}

func (r *Router) Search(path string) int {
	search := path
	now := r.root
	for {
		next, index := now.findLongestCommonPrefix(search)
		if next == nil {
			return -1
		}

		search = search[index:]
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

	fmt.Println(r.root.children)
	fmt.Println(r.root.children[0])
	fmt.Println(r.root.children[0].children[0])
	fmt.Println(r.root.children[0].children[1])
	// fmt.Println(r.root.children[0].children[0])
	fmt.Println(r.Search("/health"))
	fmt.Println(r.Search("/fuga"))
}

type node struct {
	prefix   string
	value    int
	parent   *node
	children []*node
}

func (n *node) findLongestCommonPrefix(prefix string) (*node, int) {
	max := 0
	var next *node
	for _, child := range n.children {
		index := 0
		min := len(prefix)
		if min > len(child.prefix) {
			min = len(child.prefix)
		}
		for i := 0; i < min; i++ {
			if child.prefix[i] == prefix[i] {
				index++
			}
		}
		if index > max {
			next = child
			max = index
		}
	}
	return next, max
}

func BenchmarkRadix(b *testing.B) {
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
