package tree

import "sort"

type PureTree struct {
	root *node
}

func NewPureTree(values []int) *PureTree {
	v := distinct(values)
	sort.Ints(v)

	return &PureTree{root: newNode(v)}
}

func (t *PureTree) Search(value int) bool {
	return t.search(t.root, value)
}

func (t *PureTree) Insert(value int) bool {
	return t.insert(&t.root, value)
}

func (t *PureTree) Delete(value int) bool {
	return t.delete(&t.root, value)
}

func (t *PureTree) search(n *node, value int) bool {
	if n == nil {
		return false
	}

	if value > n.value {
		return t.search(n.right, value)
	}

	if value < n.value {
		return t.search(n.left, value)
	}

	return true
}

// always ptr != nil
func (t *PureTree) insert(ptr **node, value int) bool {
	n := *ptr
	if n == nil {
		*ptr = &node{value: value}
		return false
	}

	if value > n.value {
		return t.insert(&n.right, value)
	}

	if value < n.value {
		return t.insert(&n.left, value)
	}

	return true
}

// always ptr != nil
func (t *PureTree) delete(ptr **node, value int) bool {
	n := *ptr
	if n == nil {
		return false
	}

	if value > n.value {
		return t.delete(&n.right, value)
	}

	if value < n.value {
		return t.delete(&n.left, value)
	}

	// value found

	if n.left == nil && n.right == nil {
		*ptr = nil
		return true
	}

	if n.left == nil && n.right != nil {
		*ptr = n.right
		return true
	}

	if n.left != nil && n.right == nil {
		*ptr = n.left
		return true
	}

	// both children

	min := n.right.minimum()
	n.value = min
	t.delete(&n.right, min)

	return true
}

type node struct {
	value       int
	left, right *node
}

func newNode(sortedSet []int) *node {
	if len(sortedSet) == 0 {
		return nil
	}

	if len(sortedSet) == 1 {
		return &node{value: sortedSet[0]}
	}

	if len(sortedSet) == 2 {
		return &node{
			value: sortedSet[1],
			left:  &node{value: sortedSet[0]},
		}
	}

	if len(sortedSet) == 3 {
		return &node{
			value: sortedSet[1],
			left:  &node{value: sortedSet[0]},
			right: &node{value: sortedSet[2]},
		}
	}

	midIdx := len(sortedSet) / 2
	return &node{
		value: sortedSet[midIdx],
		left:  newNode(sortedSet[0:midIdx]),
		right: newNode(sortedSet[midIdx+1:]),
	}
}

func (n *node) minimum() int {
	if n.left == nil {
		return n.value
	}

	return n.left.minimum()
}

func distinct(values []int) []int {
	vm := make(map[int]struct{})
	for _, value := range values {
		vm[value] = struct{}{}
	}

	out := make([]int, len(vm))
	i := 0
	for v := range vm {
		out[i] = v
		i++
	}

	return out
}
