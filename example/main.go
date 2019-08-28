package main

import (
	"fmt"

	"github.com/henderjon/errloc"
)

func main() {
	fmt.Println(errloc.New("this is a serious problem"))
	fmt.Println(errloc.Here())
}
