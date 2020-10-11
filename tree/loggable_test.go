package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLoggableTree_Search(t *testing.T) {
	tm := &treeMock{}
	lm := &loggerMock{}

	lt := NewLoggableTree(tm, lm)

	tm.On("Search", 123).Return(false).Once()
	lm.On("Infof", mock.AnythingOfType("string"), 123, false).Once()
	found := lt.Search(123)
	assert.False(t, found)
	tm.AssertExpectations(t)
	lm.AssertExpectations(t)

	tm.On("Search", 456).Return(true).Once()
	lm.On("Infof", mock.AnythingOfType("string"), 456, true).Once()
	found = lt.Search(456)
	assert.True(t, found)
	tm.AssertExpectations(t)
	lm.AssertExpectations(t)
}

func TestLoggableTree_Insert(t *testing.T) {
	tm := &treeMock{}
	lm := &loggerMock{}

	lt := NewLoggableTree(tm, lm)

	tm.On("Insert", 123).Return(false).Once()
	lm.On("Infof", mock.AnythingOfType("string"), 123, false).Once()
	found := lt.Insert(123)
	assert.False(t, found)
	tm.AssertExpectations(t)
	lm.AssertExpectations(t)

	tm.On("Insert", 456).Return(true).Once()
	lm.On("Infof", mock.AnythingOfType("string"), 456, true).Once()
	found = lt.Insert(456)
	assert.True(t, found)
	tm.AssertExpectations(t)
	lm.AssertExpectations(t)
}

func TestLoggableTree_Delete(t *testing.T) {
	tm := &treeMock{}
	lm := &loggerMock{}

	lt := NewLoggableTree(tm, lm)

	tm.On("Delete", 123).Return(false).Once()
	lm.On("Infof", mock.AnythingOfType("string"), 123, false).Once()
	found := lt.Delete(123)
	assert.False(t, found)
	tm.AssertExpectations(t)
	lm.AssertExpectations(t)

	tm.On("Delete", 456).Return(true).Once()
	lm.On("Infof", mock.AnythingOfType("string"), 456, true).Once()
	found = lt.Delete(456)
	assert.True(t, found)
	tm.AssertExpectations(t)
	lm.AssertExpectations(t)
}

type treeMock struct {
	mock.Mock
}

func (m *treeMock) Search(value int) (found bool) {
	args := m.Called(value)
	return args.Bool(0)
}

func (m *treeMock) Insert(value int) (found bool) {
	args := m.Called(value)
	return args.Bool(0)
}

func (m *treeMock) Delete(value int) (found bool) {
	args := m.Called(value)
	return args.Bool(0)
}

type loggerMock struct {
	mock.Mock
}

func (l *loggerMock) Infof(format string, args ...interface{}) {
	passargs := []interface{}{format}
	passargs = append(passargs, args...)
	l.Called(passargs...)
	return
}
