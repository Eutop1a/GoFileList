package judge_python

import (
	"os"
	"os/exec"
)

func Compiler(path string, fileName string) error {

	cmd := exec.Command("python", path+"//"+fileName+".py")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	return err

}
