package main

import (
	"fmt"
	"github.com/rmasci/gotabulate"
)

func main() {
	row_1 := []interface{}{"John", 20, "John smith was here."}
	row_2 := []interface{}{"Joe", 23, "Joe Smith was here too."}
	// Create an object from 2D interface array
	t := gotabulate.Create([][]interface{}{row_1, row_2})
	fmt.Println(t.Render("mysql"))
}
