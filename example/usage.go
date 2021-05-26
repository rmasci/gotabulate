package main

import (
	"fmt"
	"codecloud.web.att.com/st_cloudutils/gotabulate"
)

func main() {
	row_1 := []interface{}{"john", 20, "John smith was here."}
	row_2 := []interface{}{"bndr", 23, "bndr was here."}
	row_3 := []interface{}{"frank", 29, "frank was here."}

	// Create an object from 2D interface array
	t := gotabulate.Create([][]interface{}{row_1, row_2, row_3})

	// Set the Headers (optional)
	//t.SetHeaders([]string{"age", "status"})
	//t.SetNoHeader()
	fmt.Printf("NoHeader? %v\n",t.NoHeader)
	// Set the Empty String (optional)
	t.SetEmptyString("None")

	// Set Align (Optional)
	t.SetAlign("left")
	t.SetWrapStrings(false)
	t.SetRemEmptyLines(true)
	// Print the result: grid, or simple
	fmt.Println("Simple")
	//hl:=[]string {"bottomLine","top"}
	//t.Data[0].Continuos=false
	fmt.Println(t.Render("simple"))
	fmt.Println("Plain")
	fmt.Println(t.Render("plain"))
	fmt.Println("Grid")
	fmt.Println(t.Render("grid"))
	fmt.Println("csv")
	fmt.Println(t.Render("csv"))
	//	fmt.Println("html")
	//	fmt.Println(t.Render("html"))
}
