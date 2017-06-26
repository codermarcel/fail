package fail

import (
	"fmt"
	"io"
	"runtime/debug"
)

type failer interface {
	GetCurrent() error
	GetInner() error
	GetTrace() []byte
}

//Type check.
var _ failer = &myFail{}

type myFail struct {
	Current error
	Inner   error
	Trace   []byte
}

func (f myFail) Error() string {
	return f.Current.Error()
}

func (f myFail) GetCurrent() error {
	return f.Current
}

func (f myFail) GetInner() error {
	return f.Inner
}

func (f myFail) GetTrace() []byte {
	return f.Trace
}

//TODO: Clean this up
func (f myFail) Format(s fmt.State, verb rune) {
	normal := f.Current.Error()
	extended := f.Trace

	switch verb {
	case 's':
		if s.Flag('+') {
			io.WriteString(s, string(extended))
			return
		}
		io.WriteString(s, normal)
	case 'v':
		if s.Flag('+') {
			io.WriteString(s, string(extended))
			return
		}
		io.WriteString(s, normal)
	}
}

//Wrap is used to wrap an original error with a new error to provide more context.
func Wrap(original, new error) error {
	return myFail{new, original, debug.Stack()} //TODO: Is there a better way for the stack trace?
}

//First returns the first error in the stack (The most recent error)
//If err is the only error in the stack then it will return itself.
func First(err error) error {
	failer, ok := err.(failer)

	if ok {
		return failer.GetCurrent()
	}

	return err
}

//Second returns the next error in the stack after the first.
//If err is the only error in the stack then it will return itself.
func Second(err error) error {
	failer, ok := err.(failer)

	if ok {
		return failer.GetInner()
	}

	return err
}

//Last returns the original error (also called "cause")
//If err is the only error in the stack then it will return itself.
//This may be the raw error that wasn't wrapped yet.
func Last(err error) error {
	for err != nil {
		failer, ok := err.(failer)

		if !ok {
			break
		}

		err = failer.GetInner()
	}

	return err
}
