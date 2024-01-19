package main

import "fmt"

func main() {

	i, j := 42, 2701
	fmt.Println(&i, &j)

	p := &i
	fmt.Println(p)
	fmt.Println(*p)

	a := 4
	squareVal(a)
	fmt.Println("a has a value of", a)

	squareAdd(&a)
	fmt.Println("a has a value of", a)

}

func squareVal(v int) { //v = parameter is an int
	v *= v
	fmt.Println("Using regular variables", &v, v)
}

func squareAdd(p *int) { //p = parameter is a pointer
	*p *= *p
	fmt.Println("Using pointers", p, *p)
}
