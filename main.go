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

const (
	c1        = "\033[38;2;219;30;52m"
	c2        = "\033[38;2;219;30;82m"
	c3        = "\033[38;2;219;30;112m"
	c4        = "\033[38;2;219;30;142m"
	c5        = "\033[38;2;219;30;172m"
	c6        = "\033[38;2;219;30;202m"
	no_colour = "\033[0m"
)

var (
	host  types.Host
	logos = map[string]string{
		"dragon": `
${c3}⠀⠀⠀⠀  ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⠢⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
${c1}⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣶⠋⡆⢹⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
${c5}⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⡆⢀⣤⢛⠛⣠⣿⠀⡏⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
${c6}⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣶⣿⠟⣡⠊⣠⣾⣿⠃⣠⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
${c2}⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣴⣯⣿⠀⠊⣤⣿⣿⣿⠃⣴⣧⣄⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
${c1}⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣤⣶⣿⣿⡟⣠⣶⣿⣿⣿⢋⣤⠿⠛⠉⢁⣭⣽⠋⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
${c4}  ⠀⠀⠀⠀⠀⠀ ⠀⣠⠖⡭⢉⣿⣯⣿⣯⣿⣿⣿⣟⣧⠛⢉⣤⣶⣾⣿⣿⠋⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
${c5}⠀⠀⠀⠀⠀⠀⠀⠀⣴⣫⠓⢱⣯⣿⢿⠋⠛⢛⠟⠯⠶⢟⣿⣯⣿⣿⣿⣿⣿⣿⣦⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
${c2}⠀⠀⠀⠀⠀⠀⢀⡮⢁⣴⣿⣿⣿⠖⣠⠐⠉⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠉⠉⠛⠛⠛⢿⣶⣄⠀⠀⠀⠀⠀⠀⠀
${c3}⠀⠀⠀⠀⢀⣤⣷⣿⣿⠿⢛⣭⠒⠉⠀⠀⠀⣀⣀⣄⣤⣤⣴⣶⣶⣶⣿⣿⣿⣿⣿⠿⠋⠁⠀⠀⠀⠀⠀⠀⠀⠀
${c1}⠀⢀⣶⠏⠟⠝⠉⢀⣤⣿⣿⣶⣾⣿⣿⣿⣿⣿⣿⣟⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣧⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
${c6}⢴⣯⣤⣶⣿⣿⣿⣿⣿⡿⣿⣯⠉⠉⠉⠉⠀⠀⠀⠈⣿⡀⣟⣿⣿⢿⣿⣿⣿⣿⣿⣦⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
${c5}⠀⠀⠀⠉⠛⣿⣧⠀⣆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⠃⣿⣿⣯⣿⣦⡀⠀⠉⠻⣿⣦⠀⠀⠀⠀⠀⠀⠀⠀⠀
${c3}⠀⠀⠀⠀⠀⠀⠉⢿⣮⣦⠀⠀⠀⠀⠀⠀⠀⠀⠀⣼⣿⠀⣯⠉⠉⠛⢿⣿⣷⣄⠀⠈⢻⣆⠀⠀⠀⠀⠀⠀⠀⠀
${c2}⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠢⠀⠀⠀⠀⠀⠀⠀⢀⢡⠃⣾⣿⣿⣦⠀⠀⠀⠙⢿⣿⣤⠀⠙⣄⠀⠀⠀⠀⠀⠀⠀
${c6}⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⢋⡟⢠⣿⣿⣿⠋⢿⣄⠀⠀⠀⠈⡄⠙⣶⣈⡄⠀⠀⠀⠀⠀⠀
${c1}⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠐⠚⢲⣿⠀⣾⣿⣿⠁⠀⠀⠉⢷⡀⠀⠀⣇⠀⠀⠈⠻⡀⠀⠀⠀⠀⠀
${c4}⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢢⣀⣿⡏⠀⣿⡿⠀⠀⠀⠀⠀⠀⠙⣦⠀⢧⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
${c3}⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⠿⣧⣾⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠙⣮⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
${c5}⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠙⠛⠀⠀⠀⠀
`,
	}
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
	parts := strings.Split(output_lines[1], "\\")

	return strings.TrimSpace(parts[1]) + "@" + strings.TrimSpace(parts[0])
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
	var mints = makePlural(minutes, "min")

	var date_str string
	if hours > 0 {
		date_str = fmt.Sprintf("%d %s, %d %s", hours, hour, minutes, mints)
	} else {
		date_str = fmt.Sprintf("%d %s", minutes, mints)
	}

	return date_str
}

func getShell() string {
	// assume that powershell is being used, so we just get Version
	cmd := exec.Command("powershell", "$PsVersionTable.PSVersion.ToString()")
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	return "powershell " + string(output)
}

func getResolution() string {
	cmd := exec.Command("wmic", "path", "Win32_VideoController", "get", "CurrentHorizontalResolution")
	output, err := cmd.Output()

	if err != nil {
		return "Unknown Resolution"
	}

	output_lines := strings.Split(string(output), "\n")
	horizontal := strings.TrimSpace(output_lines[1])

	cmd = exec.Command("wmic", "path", "Win32_VideoController", "get", "CurrentVerticalResolution")
	output, err = cmd.Output()
	output_lines = strings.Split(string(output), "\n")

	if err != nil {
		return "Unknown Resolution"
	}

	vertical := strings.TrimSpace(output_lines[1])

	return fmt.Sprintf("%sx%s", horizontal, vertical)
}

func getCPU() string {
	cmd := exec.Command("wmic", "cpu", "get", "Name")
	output, err := cmd.Output()

	if err != nil {
		return "Unknown CPU"
	}
	output_lines := strings.Split(string(output), "\n")

	return output_lines[1]
}

func getGPU() string {
	cmd := exec.Command("wmic", "path", "win32_VideoController", "get", "name")
	output, err := cmd.Output()

	if err != nil {
		return "Unknown CPU"
	}
	output_lines := strings.Split(string(output), "\n")

	return output_lines[1]
}

func displayInfo(tag string, data string) {
	fmt.Println(c1 + tag + ":\033[0m " + data)
}

func displayLogo(logo string) {
	logo_data := logos[logo]
	logo_data = strings.Replace(logo_data, "${c1}", c1, -1)
	logo_data = strings.Replace(logo_data, "${c2}", c2, -1)
	logo_data = strings.Replace(logo_data, "${c3}", c3, -1)
	logo_data = strings.Replace(logo_data, "${c4}", c4, -1)
	logo_data = strings.Replace(logo_data, "${c5}", c5, -1)
	logo_data = strings.Replace(logo_data, "${c6}", c6, -1)

	fmt.Println(logo_data + no_colour)
}

func main() {
	getHost()

	displayLogo("dragon")
	fmt.Println(c1 + getLoggedInUser() + "\033[0m")
	println()
	displayInfo("OS", getOs())
	displayInfo("Host", getHostName())
	displayInfo("Kernel", getKernel())
	displayInfo("Uptime", getUptime())
	// usually package count is displayed after this but it doesn't really work for windows
	displayInfo("Shell", getShell())
	displayInfo("Resolution", getResolution())
	displayInfo("CPU", getCPU())
	displayInfo("GPU", getGPU())
}
