Go Package that makes it easy to work with errors.

`go get github.com/codermarcel/fail`

> The traditional error handling idiom in Go is roughly akin to

```go
if err != nil {
        return err
}
```
> which applied recursively up the call stack results in error reports without context or debugging information. The errors package allows programmers to add context to the failure path in their code in a way that does not destroy the original value of the error.

This is a great description of the current state and problem of errors in go as described by the https://github.com/pkg/errors package. 

**However**, the current solutions to this problem didn't really satisfy my use cases. I was looking for a simple solution that allows me to easily type check my errors directly without the need to compare strings, be able to retrieve the First or Last error, easily Wrap an existing error with a new one to provide some additional context, have a simple stack trace that helps me debugging and use something that doesn't break my existing code.

If you have the same goals as me, then this package might be for you! Let me show you some examples:

## Wrapping an existing error

`fail.Wrap` helps you to wrap an existing error with another error to add more context. This method is different from other solutions because instead of making you pass in an error message, you use your own error which allows you to typecheck for it later!


```go
var (
	ErrAdditionalContext = errors.New("Some new error for additional context")
	ErrOriginal          = errors.New("The Original error!")
)

func main() {
	err := wrapped()
}

func wrapped() error {
	return fail.Wrap(ErrOriginal, ErrAdditionalContext)
}
```

## Getting the first error in the stack (The most recent error)

`fail.First` helps you to easily retrieve the first error in the stack and allows you to do **type checks** on the error. It also does not break your code, if you pass in an error that hasn't been wraped yet, it will simply return the passed in error itself.

```go

var (
	ErrAdditionalContext = errors.New("Some new error for additional context")
	ErrOriginal          = errors.New("The Original error!")
)

func main() {
	err := wrapped()
	first := fail.First(err)
    
	if first == ErrAdditionalContext {
		fmt.Println("Yay! typechecking works!")
	}
}

func wrapped() error {
	err := Original()
	
	return fail.Wrap(err, ErrAdditionalContext)
}

func Original() error {
	return ErrOriginal
}
```

The above code would print `"Yay! typechecking works!"`


## Getting the last error in the stack

`fail.Last` helps you to easily retrieve the last error (The "Cause" or "Original" error) in the stack.
Like `fail.First` it allows you to do **type checks** on the error and does not break your code if you pass in an error that hasn't been wraped yet.

```go

var (
	ErrAdditionalContext = errors.New("Some new error for additional context")
	ErrOriginal          = errors.New("The Original error!")
)

func main() {
	err := wrapped()
	last := fail.Last(err)
    
	if last == ErrOriginal {
		fmt.Println("Yay! typechecking works!")
	}
}

func wrapped() error {
	err := Original()
	
	return fail.Wrap(err, ErrAdditionalContext)
}

func Original() error {
	return ErrOriginal
}
```

The above code would print `"Yay! typechecking works!"`


## Printing the error messages

The `fail` package helps you to seamlessly use your errors like you would normally use them.

```go

var (
	ErrAdditionalContext = errors.New("Some new error for additional context")
	ErrOriginal          = errors.New("The Original error!")
)

func main() {
	err := wrapped()
	
	fmt.Printf("%s \n", err)
	fmt.Printf("%s \n", fail.First(err))
	fmt.Printf("%s \n", fail.Last(err))
}

func wrapped() error {
	err := Original()
	
	return fail.Wrap(err, ErrAdditionalContext)
}

func Original() error {
	return ErrOriginal
}
```
**Hint** You can use `%v` instead of `%s` to get the same result.

The above code would print the following:

`"Some new error for additional context"`

`"Some new error for additional context"`

`"The Original error!"`

The printed messages are exactly what the normal error would print.

You might also have noticed that the wrapped error itself `fmt.Printf("%s", err)` also returns the First error message (`"Some new error for additional context"`) which means that you can easily pass the wrapped error into a third party library without worring about it breaking either. 

This also means that using the `.Error()` method returns your expected message, of course.

```go
var (
	ErrAdditionalContext = errors.New("Some new error for additional context")
	ErrOriginal          = errors.New("The Original error!")
)

func main() {
	err := wrapped()

	if (ErrOriginal.Error() == fail.Last(err).Error()) {
		fmt.Println("Yay! This is true!")
	}
}

func wrapped() error {
	err := Original()
	
	return fail.Wrap(err, ErrAdditionalContext)
}

func Original() error {
	return ErrOriginal
}
```

The above code would print `"Yay! This is true!"`


## Printing the stack trace

For debugging purposes, it is also really helpful to get a stack trace that helps you hunt down bugs faster. Here is how you do it:

```go

var (
	ErrAdditionalContext = errors.New("Some new error for additional context")
	ErrOriginal          = errors.New("The Original error!")
)

func main() {
	err := wrapped()
	
	fmt.Printf("%+s", err)
}

func wrapped() error {
	err := Original()
	
	return fail.Wrap(err, ErrAdditionalContext)
}

func Original() error {
	return ErrOriginal
}
```
It is very similar to the previous example, the only difference is the added `+` sign in the format.
**Hint** You can use `%+v` instead of `%+s` to get the same result.


# Version

### This Package is still in active development and I would advice against using it in production before it reaches a stable version (1.0)

# License 

See the LICENSE file.