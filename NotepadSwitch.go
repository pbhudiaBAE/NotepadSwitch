package main

import (
	"fmt"
	"io/ioutil"
)



func One() {

	first,err1 := ioutil.ReadFile("A.txt")

	if err1 != nil {
		fmt.Print(err1)
	}

	fmt.Println(first)
	str1 := string(first)
	fmt.Println(str1)
}

func Two() {

	second,err2 := ioutil.ReadFile("B.txt")

	if err2 != nil {
		fmt.Print(err2)
	}

	fmt.Println(second)
	str2 := string(second)
	fmt.Println(str2)
}


func main(){
	One()
	Two()
}