package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
)

// Args : Contains arguments for running locus
type Args struct {
	Language string `json:"language"`
	Code     string `json:"code"`
}

func main() {
	fmt.Println("***********LOCUS_RUNNER************")
	f, err := os.OpenFile("../locus_runner.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	logger := log.New(f, "LOCUS_RUNNER:", log.LstdFlags)
	input := os.Args[1]
	var args Args
	err = json.Unmarshal([]byte(input), &args)
	if err != nil {
		logger.Printf("main: Error in unmarshalling input. Reason:%v", err)
	}
	logger.Printf("Args: %v", args)
	args.convertCode(logger)
	fmt.Println("Code converted")
	fmt.Println(args.execute(logger))
}

// ConvertCode : Converts code string to File
func (s Args) convertCode(logger *log.Logger) {
	fmt.Println("Code : ", s.Code)
	file, err := os.Create("/tmp/code." + s.Language)
	if err != nil {
		logger.Fatalf("convertCode: Error in creating file. Reason:%v", err)
	}
	logger.Println("File Created!")
	line, err := file.WriteString(s.Code)
	if err != nil {
		logger.Fatalf("convertCode: Error in writing string to file. Reason:%v", err)
	}
	logger.Println("Line : ", line)
}

func (s Args) execute(logger *log.Logger) string {
	logger.Println("s.Language : ", s.Language)
	switch s.Language {
	case "py":
		op, err := exec.Command("python", "/tmp/code.py").Output()
		if err != nil {
			logger.Fatalf("execute: Error in executing code. Reason:%v", err)
		}
		return string(op)

	case "go":
		cmd := exec.Command("cd", "/bin/tmp/")
		err := cmd.Run()
		if err != nil {
			logger.Fatalf("Changing direcotry: %v", err)
		}
		_, err = exec.Command("go", "build").Output()
		if err != nil {
			logger.Fatalf("Compiler Error:%v", err)
		}
		op2, err := exec.Command("./code").Output()
		if err != nil {
			logger.Fatalf("Runtime Error:%v", err)
		}
		return string(op2)

	case "java":
		_, err := exec.Command("javac", "/tmp/main.java").Output()
		if err != nil {
			logger.Fatalf("Compiler Error:%v", err)
		}
		op2, err := exec.Command("java", "/tmp/main").Output()
		if err != nil {
			logger.Fatalf("Runtime Error:%v", err)
		}
		return string(op2)

	case "c++":
	case "c":
	}
}
