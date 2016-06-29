# Assertions for Go

Assert provides C-style assertions for Go. The assertions can be turned
on and off at run or compile time.

## Example

The assert functions are methods that operate on an `Asserter` struct
that controls whether assertions are turned on or off. If you create a
new `Asserter` with the `New` method, assertions are turned on by
default.

    package main

    import "github.com/gunnihinn/assert"

    func main() {
        a := assert.New()
        a.Assert(false, "This was not supposed to happen")
    }

This example will panic when run and the stack trace will contain both
the file name and line number of where the failing assertions was and
the message explaining what went wrong. If you turn assertions off,

    package main

    import "github.com/gunnihinn/assert"

    func main() {
        a := assert.New()
        a.Off()
        a.Assert(false, "This was not supposed to happen")
    }

the code will run and exit normally.

## Tests and coverage

The package is fully covered by tests. Run them as usual with `go test`,
and run `make cover` to get test coverage.

## Author

Gunnar Þór Magnússon <<gunnar.thor.magnusson@gmail.com>>

Licensed under GPL3.
