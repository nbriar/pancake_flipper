# Pancake Flipper

The Pancake Flipper takes a string made up of `+` and `-` which represent faceup (`+`) and face down (`-`) pancakes in a stack.

Example: `---+--+-+++++--`

## Use

To flip a stack run `make run stack=<stack_string>` from the command line

Example:

```
make run stack=+++----+--+++----++--
```


## Other Commands

`make cover` will print out a test coverage report

`make cyclo` will check for cyclomatic complexity issues in the code

`make tests` will run the full test suite

`make vet` will run `go vet` on all relevant packages

`make build` will run `go build` and result in a binary named `pancake_flipper`