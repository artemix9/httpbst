package tree

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPureTree_Search_NotFoundWhenEmpty(t *testing.T) {
	p := &PureTree{}
	expected := deepCopyNode(p.root)

	assert.False(t, p.Search(rand.Int()))
	assert.True(t, deepEqualNode(p.root, expected))
}

func TestPureTree_Search_NotFound(t *testing.T) {
	p := &PureTree{
		root: &node{
			value: 86,
			left:  &node{value: -17},
			right: &node{value: 576},
		},
	}
	expected := deepCopyNode(p.root)

	assert.False(t, p.Search(20))
	assert.True(t, deepEqualNode(p.root, expected))
}

func TestPureTree_Search_AllFound(t *testing.T) {
	p := &PureTree{
		root: &node{
			value: 86,
			left: &node{
				value: -17,
				left:  &node{value: -90},
				right: &node{value: 35},
			},
			right: &node{
				value: 576,
				left:  &node{value: 221},
				right: &node{value: 600},
			},
		},
	}
	expected := deepCopyNode(p.root)

	assert.True(t, p.Search(-90))
	assert.True(t, p.Search(-17))
	assert.True(t, p.Search(35))
	assert.True(t, p.Search(86))
	assert.True(t, p.Search(221))
	assert.True(t, p.Search(576))
	assert.True(t, p.Search(600))

	assert.True(t, deepEqualNode(p.root, expected))
}

func TestPureTree_Insert(t *testing.T) {
	p := &PureTree{}

	assert.False(t, p.Insert(86))
	assert.False(t, p.Insert(35))
	assert.False(t, p.Insert(221))

	expected := &node{
		value: 86,
		left:  &node{value: 35},
		right: &node{value: 221},
	}
	assert.True(t, deepEqualNode(p.root, expected))
}

func TestPureTest_Insert_Duplication(t *testing.T) {
	p := &PureTree{}

	assert.False(t, p.Insert(86))
	assert.True(t, p.Insert(86))

	expected := &node{value: 86}
	assert.True(t, deepEqualNode(p.root, expected))
}

func TestPureTree_Delete_NotFoundWhenEmpty(t *testing.T) {
	p := &PureTree{}
	expected := deepCopyNode(p.root)

	assert.False(t, p.Delete(rand.Int()))
	assert.True(t, deepEqualNode(p.root, expected))
}

func TestPureTree_Delete_NotFound(t *testing.T) {
	p := &PureTree{
		root: &node{
			value: 86,
			left:  &node{value: -17},
			right: &node{value: 576},
		},
	}
	expected := deepCopyNode(p.root)

	assert.False(t, p.Delete(50))
	assert.True(t, deepEqualNode(p.root, expected))
}

func TestPureTree_Delete_WithNoChildrenFromRight(t *testing.T) {
	p := &PureTree{
		root: &node{
			value: 86,
			left:  &node{value: -17},
			right: &node{value: 576},
		},
	}
	assert.True(t, p.Delete(576))

	expected := &node{
		value: 86,
		left:  &node{value: -17},
	}
	assert.True(t, deepEqualNode(p.root, expected))
}

func TestPureTree_Delete_WithNoChildrenFromLeft(t *testing.T) {
	p := &PureTree{
		root: &node{
			value: 86,
			left:  &node{value: -17},
			right: &node{value: 576},
		},
	}
	assert.True(t, p.Delete(-17))

	expected := &node{
		value: 86,
		right: &node{value: 576},
	}
	assert.True(t, deepEqualNode(p.root, expected))
}

func TestPureTree_Delete_WithLeftChild(t *testing.T) {
	p := &PureTree{
		root: &node{
			value: 86,
			left: &node{
				value: -17,
				left:  &node{value: -90},
				right: &node{value: 35},
			},
			right: &node{
				value: 576,
				left:  &node{value: 221},
			},
		},
	}
	assert.True(t, p.Delete(576))

	expected := &node{
		value: 86,
		left: &node{
			value: -17,
			left:  &node{value: -90},
			right: &node{value: 35},
		},
		right: &node{value: 221},
	}
	assert.True(t, deepEqualNode(p.root, expected))
}

func TestPureTree_Delete_WithRightChild(t *testing.T) {
	p := &PureTree{
		root: &node{
			value: 86,
			left: &node{
				value: -17,
				left:  &node{value: -90},
				right: &node{value: 35},
			},
			right: &node{
				value: 576,
				right: &node{value: 600},
			},
		},
	}
	assert.True(t, p.Delete(576))

	expected := &node{
		value: 86,
		left: &node{
			value: -17,
			left:  &node{value: -90},
			right: &node{value: 35},
		},
		right: &node{value: 600},
	}
	assert.True(t, deepEqualNode(p.root, expected))
}

func TestPureTree_Delete_WithBothChildren(t *testing.T) {
	p := &PureTree{
		root: &node{
			value: 86,
			left: &node{
				value: -17,
				left:  &node{value: -90},
				right: &node{value: 35},
			},
			right: &node{
				value: 576,
				left:  &node{value: 221},
				right: &node{value: 600},
			},
		},
	}
	assert.True(t, p.Delete(86))

	expected := &node{
		value: 221,
		left: &node{
			value: -17,
			left:  &node{value: -90},
			right: &node{value: 35},
		},
		right: &node{
			value: 576,
			right: &node{value: 600},
		},
	}
	assert.True(t, deepEqualNode(p.root, expected))
}

func TestNewPureTree_Empty(t *testing.T) {
	p := NewPureTree([]int{})
	assert.True(t, deepEqualNode(p.root, nil))
}

func TestNewPureTree_One(t *testing.T) {
	p := NewPureTree([]int{3})
	expected := &node{value: 3}
	assert.True(t, deepEqualNode(p.root, expected))
}

func TestNewPureTree_Two(t *testing.T) {
	p := NewPureTree([]int{5, 4})
	expected := &node{value: 5, left: &node{value: 4}}
	assert.True(t, deepEqualNode(p.root, expected))
}

func TestNewPureTree_Three(t *testing.T) {
	p := NewPureTree([]int{5, 4, 6})
	expected := &node{
		value: 5,
		left:  &node{value: 4},
		right: &node{value: 6},
	}
	assert.True(t, deepEqualNode(p.root, expected))
}

func TestNewPureTree_Many(t *testing.T) {
	p := NewPureTree([]int{600, 600, 576, 221, 221, 221, 86, 35, 86, -17, -90, 86})
	expected := &node{
		value: 86,
		left: &node{
			value: -17,
			left:  &node{value: -90},
			right: &node{value: 35},
		},
		right: &node{
			value: 576,
			left:  &node{value: 221},
			right: &node{value: 600},
		},
	}
	assert.True(t, deepEqualNode(p.root, expected))
}

func deepEqualNode(first, second *node) bool {
	if first == nil && second != nil ||
		first != nil && second == nil {

		return false
	}

	if first == nil && second == nil {
		return true
	}

	return (first.value == second.value) &&
		deepEqualNode(first.left, second.left) &&
		deepEqualNode(first.right, second.right)
}

func deepCopyNode(n *node) *node {
	if n == nil {
		return nil
	}

	return &node{
		value: n.value,
		left:  deepCopyNode(n.left),
		right: deepCopyNode(n.right),
	}
}
