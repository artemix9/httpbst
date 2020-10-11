package tree

type Tree interface {
	Search(value int) (found bool)
	Insert(value int) (found bool)
	Delete(value int) (found bool)
}
