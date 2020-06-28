package languages

import (

	"github.com/itsjwala/locus/runner/languages/python"

	"fmt"

)

type runFn func(string) (string, string, error)

var languagesMap = map[string] runFn {
	"python" : python.Run,
}

func IsSupported(language string) bool {
	fmt.Println(language, languagesMap[language])
	return true
}


func Run(language string, code string) (string, string, error){
	return languagesMap[language](code)
}
