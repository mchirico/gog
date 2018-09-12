package main

import (
	"github.com/mchirico/gog/rpkg"
)

func main() {
	a := rpkg.App{}
	a.Initilize()
	a.Run("8082")
}
