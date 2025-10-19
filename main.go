package main

import (
	"fmt"
	"os/exec"
	"strings"

	sysinfo "github.com/elastic/go-sysinfo"
	"github.com/elastic/go-sysinfo/types"
)

var (
	host types.Host
)

func getHost() {
	h, err := sysinfo.Host()
	if err != nil {
		panic(err)
	}

	host = h
}

func getLoggedInUser() string {
	cmd := exec.Command("wmic", "computersystem", "get", "username")
	output, err := cmd.Output()

	if err != nil {
		return "Unknown User"
	}
	output_lines := strings.Split(string(output), "\n")

	return output_lines[1]
}

func getOs() string {
	cmd := exec.Command("wmic", "os", "get", "Caption")
	output, err := cmd.Output()

	if err != nil {
		return "Unknown OS"
	}
	output_lines := strings.Split(string(output), "\n")

	return output_lines[1]
}

func getHostName() string {
	cmd := exec.Command("wmic", "computersystem", "get", "manufacturer,model")
	output, err := cmd.Output()

	if err != nil {
		return "Unknown Computer Model"
	}
	output_lines := strings.Split(string(output), "\n")

	return output_lines[1]
}

func getKernel() string {
	cmd := exec.Command("wmic", "os", "get", "Version")
	output, err := cmd.Output()

	if err != nil {
		return "Unknown Kernel"
	}
	output_lines := strings.Split(string(output), "\n")

	return output_lines[1]
}

func getUptime() string {
	return "Will work on this later"
}

func main() {
	getHost()
	fmt.Println(getLoggedInUser())
	println()
	fmt.Println(getOs())
	fmt.Println(getHostName())
	fmt.Println(getKernel())
	fmt.Println(getUptime())
}
