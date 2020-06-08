package python


import (
	"fmt"
)


func Run(code string) (string, string , error){

	fmt.Println("Running python code")
	fmt.Println(code)
	return "","", nil
}