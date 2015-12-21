// Get the number of open files for a process and compare that against /proc/<pid>/limits and alert if
// over the given threshold.
//
//
// LICENSE:
//   Copyright 2015 Yieldbot. <devops@yieldbot.com>
//   Released under the MIT License; see LICENSE
//   for details.

package main

import (
	"flag"
	"fmt"
	"github.com/yieldbot/ybsensupluginfile/cmd/check-file-handles/Godeps/_workspace/src/github.com/yieldbot/ybsensuplugin/ybsensupluginutil"
	"github.com/yieldbot/ybsensupluginfile/cmd/check-file-handles/Godeps/_workspace/src/github.com/yieldbot/ybsensupluginfile/ybfilesys"
	"os"
)

// Calculate if the value is over a threshold
func determineThreshold(limit float64, threshold float64, numFD float64) bool {
	alarm := true
	tLimit := threshold / float64(100) * limit

	if numFD > float64(tLimit) {
		alarm = true
	} else {
		alarm = false
	}
	return alarm
}

func main() {

	// set commandline flags
	AppPtr := flag.String("app", "sbin/init", "the process name")
	WarnPtr := flag.Int("warn", 75, "the alert warning threshold percentage")
	CritPtr := flag.Int("crit", 75, "the alert critical threshold percentage")
	DebugPtr := flag.Bool("debug", false, "print debugging info instead of an alert")
	JavaAppPtr := flag.Bool("java", false, "java apps process detection is different")

	flag.Parse()
	app := *AppPtr
	warnThreshold := *WarnPtr
	critThreshold := *CritPtr
	ybsensupluginutil.Debug = *DebugPtr
	ybfilesys.JavaApp = *JavaAppPtr

	var appPid string
	var sLimit, hLimit, openFd float64

	if app != "" {
		appPid = ybfilesys.GetPid(app)
		sLimit, hLimit, openFd = ybfilesys.GetFileHandles(appPid)
		if ybsensupluginutil.Debug {
			fmt.Printf("warning threshold: %v percent, critical threshold: %v percent\n", warnThreshold, critThreshold)
			fmt.Printf("this is the number of open files at the specific point in time: %v\n", openFd)
			fmt.Printf("app pid is: %v\n", appPid)
			fmt.Printf("This is the soft limit: %v\n", sLimit)
			fmt.Printf("This is the hard limit: %v\n", hLimit)
			os.Exit(0)
		}
		if determineThreshold(hLimit, float64(critThreshold), openFd) {
			fmt.Printf("%v is over %v percent of the the open file handles hard limit of %v\n", app, critThreshold, hLimit)
			os.Exit(2)
		} else if determineThreshold(sLimit, float64(warnThreshold), openFd) {
			fmt.Printf("%v is over %v percent of the open file handles soft limit of %v\n", app, warnThreshold, sLimit)
			os.Exit(1)
		} else {
			fmt.Printf("There was an error calculating the thresholds. Check to make sure everything got convert to a float64.\n")
			fmt.Printf("If unsure of the use, consult the documentation for examples and requirements\n")
			os.Exit(ybsensupluginutil.MonitoringErrorCodes["RUNTIME_ERROR"])
		}
	} else {
		fmt.Printf("Please enter a process name to check. \n")
		fmt.Printf("If unsure consult the documentation for examples and requirements\n")
		os.Exit(ybsensupluginutil.MonitoringErrorCodes["CONFIG_ERROR"])
	}
}
