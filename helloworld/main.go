package main //Need's to be named main if you want to built into an executable

import "fmt" //give this package access to a standard library package in order to be linked to any of the standard library it needs to be imported. We can see more packages and descriptions at this link. https://golang.org/pkg

//func is a function. Declare a function, name it and go through the process of creating the function the same way I create them in javascript
func main() {
	fmt.Println("Hi there!")
}
