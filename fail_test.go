package fail

import (
	"errors"
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)

var (
	ErrOriginal = errors.New("The original error ('cause')")
	ErrMiddle   = errors.New("The error in the middle")
	ErrFirst    = errors.New("The very first error")
)

func TestFirstAndLast(t *testing.T) {
	err := makeWrappedFirstError()
	first := First(err)
	last := Last(err)

	assert.Equal(t, first, ErrFirst)
	assert.Equal(t, first.Error(), ErrFirst.Error())
	assert.Equal(t, fmt.Sprintf("%s", first), ErrFirst.Error())
	assert.Equal(t, fmt.Sprintf("%v", first), ErrFirst.Error())
	assert.Equal(t, fmt.Sprintf("%+v", first), ErrFirst.Error())

	assert.Equal(t, ErrOriginal, last)
	assert.Equal(t, ErrOriginal.Error(), last.Error())
	assert.Equal(t, fmt.Sprintf("%s", last), ErrOriginal.Error())
	assert.Equal(t, fmt.Sprintf("%v", last), ErrOriginal.Error())
	assert.Equal(t, fmt.Sprintf("%+v", last), ErrOriginal.Error())
}

func TestReturnsTrace(t *testing.T) {
	err := makeWrappedFirstError()

	assert.NotEqual(t, fmt.Sprintf("%+s", err), fmt.Sprintf("%s", err))
	assert.NotEqual(t, fmt.Sprintf("%+v", err), fmt.Sprintf("%v", err))
}

//Return itself if you pass in an unwrapped error.
func TestFirstReturnsItself(t *testing.T) {
	err := ErrFirst
	first := First(err)
	last := Last(err)

	assert.Equal(t, err, first)
	assert.Equal(t, first, last)
}

//Return itself if you pass in an unwrapped error.
func TestLastReturnsItself(t *testing.T) {
	err := ErrFirst
	first := First(err)
	last := Last(err)

	assert.Equal(t, err, last)
	assert.Equal(t, first, last)
}


func makeWrappedFirstError() error {
	return Wrap(makeWrappedMiddleError(), ErrFirst)
}
func makeWrappedMiddleError() error {
	return Wrap(makeOriginalError(), ErrMiddle)
}

func makeOriginalError() error {
	return ErrOriginal
}
