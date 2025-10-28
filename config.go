package main

//
// This file acts as a config file that can be modified to change what happens when you run the program
//
// just change the contents of the config() function to make any modifications
//

func config() {
	displayLogo("windows")

	moveInfoLeft(0)

	// printWhitespace(3)
	displayInfo("", getLoggedInUser())
	displayInfo("", "----------")

	displayInfo("OS", getOs())
	// displayInfo("Host", getHostName())
	displayInfo("Kernel", getKernel())
	displayInfo("Uptime", getUptime())

	displayInfo("Shell", getShell())
	displayInfo("Resolution", getResolution())
	displayInfo("CPU", getCPU())
	displayInfo("GPU", getGPU())
}
