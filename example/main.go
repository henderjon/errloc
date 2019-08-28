package main

import (
	"fmt"

	"github.com/henderjon/errmsg"
)

func main() {
	fmt.Println(errmsg.New("this is a serious problem"))
	fmt.Println(errmsg.Here())
}
