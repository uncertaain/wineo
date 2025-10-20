package main

import (
	"fmt"
	"math"
	"os/exec"
	"strings"
	"time"

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

// makes a word plural if it should be and not if it shouldn't
func makePlural(quantity int, word string) string {
	if quantity == 1 {
		return word
	}
	return word + "s"
}

func getUptime() string {
	// 						  YYYYMMDDHHMMSS
	time_layout := "20060102150405"
	location := time.Now().Location()

	cmd := exec.Command("wmic", "os", "get", "lastbootuptime")
	output, err := cmd.Output()

	if err != nil {
		return "Unknown Boot up time"
	}
	output_lines := strings.Split(string(output), "\n")

	raw_date := strings.Split(output_lines[1], ".")[0]
	last_boot_time, err := time.ParseInLocation(time_layout, raw_date, location)
	if err != nil {
		panic(err)
	}

	now := time.Now()
	time_since_last_boot := now.Sub(last_boot_time).Seconds()

	hours := int(math.Floor(time_since_last_boot / 3600))
	time_since_last_boot -= float64(hours) * 3600
	minutes := int(math.Floor(time_since_last_boot / 60))
	time_since_last_boot -= float64(minutes) * 60

	var hour = makePlural(hours, "hour")
	var min = makePlural(minutes, "min")

	var date_str string
	if hours > 0 {
		date_str = fmt.Sprintf("%d %s, %d %s", hours, hour, minutes, min)
	} else {
		date_str = fmt.Sprintf("%d %s", minutes, min)
	}

	return date_str
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
