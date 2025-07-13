package main

import (
	"fmt"

	"github.com/DrGMark7/Go-100-day/cmd/package/test1"
)


func main(){
	test1.SayHello()
	fmt.Println(test1.PublicFunc(-1))
}