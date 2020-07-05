package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/itsjwala/locus/runner/languages"
	"os"
)

type capability struct {
	Code     string `json:"code"`
	Language string `json:"language"`
}

func main() {
	caps := &capability{}
	jsonArg := os.Args[1]
	if err := json.Unmarshal([]byte(jsonArg), caps); err != nil {
		errString := fmt.Sprintf("%s : %s", err, jsonArg)
		fmt.Println(jsonifyResult("", "", errors.New(errString)))
	}

	if !languages.IsSupported(caps.Language) {
		errString := fmt.Sprintf("%s is not supported\n", caps.Language)
		fmt.Println(jsonifyResult("", "", errors.New(errString)))
	} else {
		if codeFilePath, err := createCodeFile(caps.Code); err != nil {
			errString := fmt.Sprintf("%s", err)
			fmt.Println(jsonifyResult("", "", errors.New(errString)))
		} else {
			stdout, stderr, error := languages.Run(caps.Language, codeFilePath)
			fmt.Println(jsonifyResult(stdout, stderr, error))
		}
	}
}
