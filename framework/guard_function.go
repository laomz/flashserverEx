package framework

import (
	"fmt"
)

func GuardFunction(info ...interface{}) {
	if err := recover(); err != nil {
		fmt.Printf("GuardFunction:err=%+v,info=%+v\n", err, info)
	}
}
