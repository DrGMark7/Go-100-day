package test1

import "fmt"

func SayHello(){ //. ถ้า function เริ่มด้วยตัว Upper Case หมายถึงเป็น Public Funciton
	fmt.Println("Hello World")
}

func privateFunc()(string){ //. ถ้า function เริ่มต้น ด้วย Lower Case หมายถึงเป็น Private function 
	return "string"
}

func PublicFunc(a int)(string){
	if a == 1{
		return privateFunc()
	}else{
		return privateFunc2()
	}
}