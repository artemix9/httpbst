package tree

import "sync"

// tree decorator with synchronization
type SafeTree struct {
	tree Tree
	m    sync.RWMutex
}

func NewSafeTree(tree Tree) *SafeTree {
	return &SafeTree{tree: tree}
}

func (t *SafeTree) Search(value int) bool {
	t.m.RLock()
	defer t.m.RUnlock()

	return t.tree.Search(value)
}

func (t *SafeTree) Insert(value int) bool {
	t.m.Lock()
	defer t.m.Unlock()

	return t.tree.Insert(value)
}

func (t *SafeTree) Delete(value int) bool {
	t.m.Lock()
	defer t.m.Unlock()

	return t.tree.Delete(value)
}
