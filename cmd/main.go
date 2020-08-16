package main

import (
	"github.com/rahulsidpatil/qlikapp/pkg/controller"
)

func main() {
	a := controller.App{}
	a.Initialize()
	a.Run()
}
