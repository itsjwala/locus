package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type result struct {
	Stdout string `json:"stdout"`
	Stderr string `json:"stderr"`
	Error  string `json:"error"`
}

var CODE_DIR string = os.Getenv("CODE_DIR")

func jsonifyResult(stdout string, stderr string, err error) string {

	res := &result{}

	res.Stdout = stdout
	res.Stderr = stderr
	if err != nil {
		res.Error = fmt.Sprintf("%v", err)
	}
	if jsonBytes, err := json.Marshal(res); err != nil {
		fmt.Printf("Error occured while marshalling result %s", err)
		panic(err)
	} else {
		return string(jsonBytes)
	}

}

func createCodeFile(code string) (string, error) {
	filename := CODE_DIR + "/code"
	if err := ioutil.WriteFile(filename, []byte(code), 0644); err != nil {
		return "", err
	}
	return filename, nil
}
