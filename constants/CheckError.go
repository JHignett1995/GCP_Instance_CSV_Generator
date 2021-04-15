package constants

import "fmt"

func ErrorCheck(msg string, err error){
	if err !=nil{
		fmt.Println(msg+ ": " + err.Error())
	}else{
		fmt.Println(msg)
	}
}
