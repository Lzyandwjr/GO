package main

import "fmt"

func d(){
	fmt.Println("func a")
}
func b(){
	defer func(){
		err := recover()
		if err != nil {
			fmt.Println("func b error")
		}
	}()
	fmt.Println("func b")
}
func c(){
	fmt.Println("func c")
}

func  main(){

	d()
	b()
	c()

}

