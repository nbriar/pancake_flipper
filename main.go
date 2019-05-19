package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Pancakes holds the state of the stack
type Pancakes struct {
	Stack string
	Count int
}

func main() {
	input := validateInput(os.Args[1])

	fmt.Printf("Input Stack: %v \n", input)

	p := Pancakes{Stack: input}
	p.SortStack()
}

// SortStack is used to sort the stack of pancakes
// It then returns the modified Pancake stack until all pancakes are face up
func (p *Pancakes) SortStack() *Pancakes {
	l := len(p.Stack)
	fdi := strings.Index(p.Stack, "-")
	fui := strings.Index(p.Stack, "+")

	if l == 0 {
		panic("You must have a stack of at least 1 pancake!")
	}

	if fui == -1 {
		p.Stack = flipstack(p.Stack)
		p.Count++
		fmt.Printf("Flip#: %d, Result: %s \n", p.Count, p.Stack)
		return p
	}

	c := 0
	for fdi > -1 {
		// Let's limit the number of times we can flip the stack. We don't want to get carried away.
		c++
		if c > 100 {
			panic("Something really went wrong and your function ran away! Better go catch it.")
		}

		// Take all the face down pancakes at the start of the stack and flip them
		if fui > fdi {
			fds := p.Stack[:fui]
			flipped := flipstack(fds)

			p.Stack = flipped + p.Stack[fui:]
			p.Count++

			fmt.Printf("Flip#: %d, Result: %s \n", p.Count, p.Stack)
			return p.SortStack()
		}

		// Take all the face up pancakes at the start of the stack and flip them
		if fdi > fui {
			fus := p.Stack[:fdi]
			flipped := flipstack(fus)

			p.Stack = flipped + p.Stack[fdi:]
			p.Count++

			fmt.Printf("Flip#: %d, Result: %s \n", p.Count, p.Stack)
			return p.SortStack()
		}
	}

	return p
}

// flipstack simulates flipping the given stack
// It just reverses the sign on the string given
// Only the head of the stack should ever be given
func flipstack(s string) string {
	fd := strings.Index(s, "-")
	if fd > -1 {
		return strings.Replace(s, "-", "+", -1)
	}

	return strings.Replace(s, "+", "-", -1)
}

// validateInput validates the stack input from the command line is only comprised of `+` or `-`
func validateInput(s string) string {
	match, _ := regexp.MatchString("^[\\-\\+]+$", s)
	if !match {
		panic("Input must be a string made up of only `+` or `-`!")
	}
	return s
}
