package server

import (
	"fmt"
	"github.com/chensheng0/repoB/client"
)

func ProductA() {
	fmt.Println("aa")
	client.ProductB()
}
