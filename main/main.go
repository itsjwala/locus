package main

import (
	"fmt"
	"log"
	"os"
)

type Args struct {
	Language string
	Code     string
}

func main() {
	fmt.Println("***********LOCUS_RUNNER************")
	f, err := os.OpenFile("../locus_runner.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	logger := log.New(f, "LOCUS_RUNNER:", log.LstdFlags)
	a := &Args{
		Language: os.Args[1],
		Code:     os.Args[2],
	}
	a.convertCode(logger)
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
	fmt.Println("Line : ", line)
}
