package main

import "fmt"

func main() {
	l := &linked{}

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			l.append(i)
		} else {
			l.insert(i)
		}
	}

	fmt.Println(l.length)
	l.show()
}
