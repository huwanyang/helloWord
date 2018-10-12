package helloworld

import "fmt"

func SayHelloWorld() {
	fmt.Printf("%s", "Hello world!")
}

func Version() string {
	return "v1.0.0"
}

func SayOther(str string) {
	fmt.Printf("Saying %s", str)
}
