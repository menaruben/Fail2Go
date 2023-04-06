package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("powershell.exe", "Get-Process | Select-Object -First 3")
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(stdout))
}
