// Get the number of open files for a process and compare that against /proc/<pid>/limits and alert if
// over the given threshold.
//
//
// LICENSE:
//   Copyright 2015 Yieldbot. <devops@yieldbot.com>
//   Released under the MIT License; see LICENSE
//   for details.

package filesys

import (
	"fmt"
	"github.com/yieldbot/ybsensuplugin/ybsensupluginutil"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

//GetPid returns the pid for the desired process
func GetPid(app string) string {
	pidExp := regexp.MustCompile("[0-9]+")
	termExp := regexp.MustCompile(`pts/[0-9]`)
	appPid := ""

	/// the pid for the binary
	goPid := os.Getpid()
	if ybsensupluginutil.Debug {
		fmt.Printf("golang binary pid: %v\n", goPid)
	}

	psAEF := exec.Command("ps", "-aef")

	out, err := psAEF.Output()
	if err != nil {
		ybsensupluginutil.EHndlr(err)
	}

	psAEF.Start()

	lines := strings.Split(string(out), "\n")

	if !JavaApp {
		for i := range lines {
			if !strings.Contains(lines[i], strconv.Itoa(goPid)) && !termExp.MatchString(lines[i]) {
				words := strings.Split(lines[i], " ")
				for j := range words {
					if app == words[j] {
						appPid = pidExp.FindString(lines[i])
					}
				}
			}
		}
	} else {
		for i := range lines {
			if strings.Contains(lines[i], app) && !strings.Contains(lines[i], strconv.Itoa(goPid)) && !termExp.MatchString(lines[i]) {
				appPid = pidExp.FindString(lines[i])

			}
		}
	}
	if appPid == "" {
		fmt.Printf("No process with the name " + app + " exists.\n")
		fmt.Printf("If unsure consult the documentation for examples and requirements\n")
		os.Exit(ybsensupluginutil.MonitoringErrorCodes["CONFIG_ERROR"])
	}
	return appPid
}

//GetFileHandles returns the current number of open file handles for a process
func GetFileHandles(pid string) (float64, float64, float64) {
	var _s, _h string
	var s, h float64
	limitExp := regexp.MustCompile("[0-9]+")
	filename := `/proc/` + pid + `/limits`
	fdLoc := "/proc/" + pid + "/fd"
	numFD := 0.0

	limits, err := ioutil.ReadFile(filename)
	if err != nil {
		ybsensupluginutil.EHndlr(err)
	}

	lines := strings.Split(string(limits), "\n")
	for i := range lines {
		if strings.Contains(lines[i], "open files") {
			limits := limitExp.FindAllString(lines[i], 2)
			_s = limits[0]
			_h = limits[1]

			s, err = strconv.ParseFloat(_s, 64)
			if err != nil {
				ybsensupluginutil.EHndlr(err)
				os.Exit(2)
			}
			h, err = strconv.ParseFloat(_h, 64)
			if err != nil {
				ybsensupluginutil.EHndlr(err)
				os.Exit(2)
			}
		}
	}

	files, _ := ioutil.ReadDir(fdLoc)
	for _ = range files {
		numFD++
	}
	if numFD == 0.0 {
		fmt.Printf("There are no open file descriptors for the process, did you use sudo?\n")
		fmt.Printf("If unsure of the use, consult the documentation for examples and requirements\n")
		os.Exit(ybsensupluginutil.MonitoringErrorCodes["PERMISSION_ERROR"])
	}
	return s, h, numFD
}
