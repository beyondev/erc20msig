package debugutils

import "fmt"

func Println(_type string, a ...interface{}) {
	if printsConfigMap[_type] {
		fmt.Printf("【%s】", _type)
		fmt.Println(a...)
	}
}

func Printf(_type string, msg string, a ...interface{}) {
	if printsConfigMap[_type] {
		fmt.Printf("【%s】", _type)
		fmt.Printf(msg, a...)
	}
}
