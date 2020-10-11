package tree

// tree decorator with logging
type LoggableTree struct {
	tree   Tree
	logger Logger
}

type Logger interface {
	Infof(format string, args ...interface{})
}

func NewLoggableTree(tree Tree, logger Logger) *LoggableTree {
	return &LoggableTree{tree: tree, logger: logger}
}

func (t *LoggableTree) Search(value int) bool {
	found := t.tree.Search(value)
	t.logger.Infof("search value %d found %t", value, found)
	return found
}

func (t *LoggableTree) Insert(value int) bool {
	found := t.tree.Insert(value)
	t.logger.Infof("insert value %d found %t", value, found)
	return found
}

func (t *LoggableTree) Delete(value int) bool {
	found := t.tree.Delete(value)
	t.logger.Infof("insert value %d found %t", value, found)
	return found
}
