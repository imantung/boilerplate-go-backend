package cmdkit

import (
	"os"
	"os/exec"
)

func GoImports(targetDir string) error {
	cmd := exec.Command("go", "run", "golang.org/x/tools/cmd/goimports@latest", "-w", targetDir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func GoMock(dest, src string) error {
	cmd := exec.Command("go", "run", "github.com/golang/mock/mockgen@v1.6.0",
		"-destination", dest,
		"-source", src,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
