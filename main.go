package main

import (
	"fmt"
	"os/exec"

	sysinfo "github.com/elastic/go-sysinfo"
)

func getLoggedInUser() string {
	cmd := exec.Command("whoami")
	output, err := cmd.Output()

	if err != nil {
		return "Unknown User"
	}

	return string(output)
}

func getOs() string {
	host, err := sysinfo.Host()
	if err != nil {
		panic(err)
	}

	value := host.Info()
	return fmt.Sprintf("%s %s", string(value.OS.Name), string(value.OS.Version))
}

func main() {
	fmt.Println(getLoggedInUser())
	fmt.Print(getOs())
}
