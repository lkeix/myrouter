package bench

import (
	"fmt"
	"testing"
)

type trienode struct {
	b        byte
	origin   string
	value    int
	parent   *trienode
	children []*trienode
}

func (n *trienode) findLongestCommonPrefix(b byte) (*trienode, int) {
	max := 0
	var next *trienode
	for _, child := range n.children {
		if child.b == b {
			return child, 1
		}
	}
	return next, max
}

type TrieRouter struct {
	root *trienode
}

func (r *TrieRouter) Insert(path string, value int) {
	search := path
	now := r.root
	for {
		if len(search) == 0 {
			return
		}
		next, _ := now.findLongestCommonPrefix(search[0])
		if next == nil {
			nn := &trienode{
				b:        search[0],
				origin:   search,
				parent:   now,
				value:    value,
				children: make([]*trienode, 0),
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

func (r *TrieRouter) Search(path string) int {
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

type RadixRouter struct {
	root *node
}

func (r *RadixRouter) Insert(path string, value int) {
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

		if next.prefix[index:] == "" {
			now = next
			search = search[index:]
			continue
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

func (r *RadixRouter) Search(path string) int {
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

func (r *RadixRouter) SearchWithIndex(path string) int {
	search := path
	now := r.root
	for {
		next := now.findWithIndex(search[0])
		if next == nil {
			return -1
		}

		min := len(search)
		if len(next.prefix) < min {
			min = len(next.prefix)
		}
		i := 0
		for ; i < min && search[i] == next.prefix[i]; i++ {
		}
		search = search[i:]
		if len(search) == 0 {
			return next.value
		}
		now = next
	}
}

func TestInsertAndSearch(t *testing.T) {
	r := &RadixRouter{
		root: &node{
			children: make([]*node, 0),
		},
	}

	paths := []string{
		"/health",
		"/hogehoge",
		"/fuga",
		"/piyo",
		"/piyopiyo",
		"/bar",
		"/bazbaz",
	}

	values := []int{
		1, 2, 3, 4, 5, 6, 7,
	}

	for i, path := range paths {
		r.Insert(path, values[i])
	}

	for i, path := range paths {
		if values[i] != r.Search(path) {
			panic(fmt.Sprintf("invalid search data: %s", path))
		}
	}
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

func (n *node) findWithIndex(b byte) *node {
	for _, child := range n.children {
		if child.prefix[0] == b {
			return child
		}
	}
	return nil
}

func BenchmarkSearch(b *testing.B) {
	b.ResetTimer()
	paths := []string{
		"/health",
		"/hogehoge",
		"/fuga",
		"/piyo",
		"/piyopiyo",
		"/bar",
		"/bazbaz",
	}

	values := []int{
		1, 2, 3, 4, 5, 6, 7,
	}

	r := &RadixRouter{
		root: &node{
			children: make([]*node, 0),
		},
	}

	for i, path := range paths {
		r.Insert(path, values[i])
	}
	for i := 0; i < b.N; i++ {
		for j, path := range paths {
			v := r.Search(path)
			if values[j] != v {
				fmt.Println(v, values[j])
				panic("invalid search data")
			}
		}
	}
}

func BenchmarkSearchWithInitial(b *testing.B) {
	b.ResetTimer()
	paths := []string{
		"/health",
		"/hogehoge",
		"/fuga",
		"/piyo",
		"/piyopiyo",
		"/bar",
		"/bazbaz",
	}

	values := []int{
		1, 2, 3, 4, 5, 6, 7,
	}

	r := &RadixRouter{
		root: &node{
			children: make([]*node, 0),
		},
	}

	for i, path := range paths {
		r.Insert(path, values[i])
	}
	for i := 0; i < b.N; i++ {
		for j, path := range paths {
			v := r.SearchWithIndex(path)
			if values[j] != v {
				panic("invalid search data")
			}
		}
	}
}

func BenchmarkTrie(b *testing.B) {
	b.ResetTimer()
	paths := []string{
		"/health",
		"/hogehoge",
		"/fuga",
		"/piyo",
		"/piyopiyo",
		"/bar",
		"/bazbaz",
	}

	values := []int{
		1, 2, 3, 4, 5, 6, 7,
	}

	r := &TrieRouter{
		root: &trienode{
			children: make([]*trienode, 0),
		},
	}

	for i, path := range paths {
		r.Insert(path, values[i])
	}
	for i := 0; i < b.N; i++ {
		for j, path := range paths {
			v := r.Search(path)
			if values[j] != v {
				panic("invalid search data")
			}
		}
	}
}
