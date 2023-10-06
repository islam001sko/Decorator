package main

import "fmt"

func main() {

	developer := &SeniorGoDeveloper{&GoDeveloper{}}

	fmt.Println(developer.writeCode())
}
