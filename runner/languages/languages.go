package languages

import (
	"github.com/itsjwala/locus/runner/languages/python"
	"github.com/itsjwala/locus/runner/languages/nodejs"
)

type runFn func(string) (string, string, error)

var languagesMap = map[string] runFn {
	"python" : python.Run,
	"nodejs" : nodejs.Run,
}

func IsSupported(language string) bool {
	_,present := languagesMap[language]
	return present
}

func Run(language string, code string) (string, string, error){
	return languagesMap[language](code)
}
