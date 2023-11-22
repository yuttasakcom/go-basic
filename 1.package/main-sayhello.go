package main

import "gitlab.com/yuttasakcom/go-basic/palm"

func main() {
	name := palm.SayHello()

	palm.HelloPalm(name)
}
