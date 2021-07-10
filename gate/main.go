package main

import (
	"common"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("%s\n", common.ServiceConfigNote)
		return
	}

	common.ServiceWait("gate")
}
