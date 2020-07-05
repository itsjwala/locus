package python

import (
	"os/exec"
	"path/filepath"
	"bytes"
)

func Run(codeFilePath string) (string, string , error){
	cmd := exec.Command("python")
	var stdout  bytes.Buffer	
	var stderr  bytes.Buffer	
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.Dir = filepath.Dir(codeFilePath)	
	cmd.Args = append(cmd.Args,filepath.Base(codeFilePath))
	err := cmd.Run()
	return stdout.String(),stderr.String(), err
}
