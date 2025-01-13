package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	env := append(os.Environ(), "GO_BALEK_SHELL=1")
	env = append(os.Environ(), "PS1=gobalekshell> ")
	// Start a balek shell
	cmd := exec.Command("bash")
	// Connect the standard input, output and error of the balek process to the current process and then balek
	cmd.Env = env
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the main balek command
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Balek: %s\n", err)
	} else {
		fmt.Printf("Bye and thanks for using GoBalekShell !\n")
	}
}
