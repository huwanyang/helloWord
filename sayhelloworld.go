package helloworld

import "fmt"

func SayHelloWorld() {
	fmt.Println("Hello world!")
}

func Version() string {
	return "v1.0.0"
}

func SayOther(str string) {
	fmt.Printf("Saying: %s\n", str)
}
