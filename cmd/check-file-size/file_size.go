package main

import (
	"fmt"
	"github.com/yieldbot/ybsensupluginfile/cmd/check-file-size/Godeps/_workspace/src/github.com/yieldbot/ybsensuplugin/ybsensupluginutil"
	"os"
	"path/filepath"
)

func testFile(f string) string {
	file, err := os.Open(f)
	if err != nil {
		ybsensupluginutil.EHndlr(err)
	}
	defer file.Close()

	// get the file size
	stat, err := file.Stat()
	if err != nil {
		ybsensupluginutil.EHndlr(err)
	}

	fmt.Printf("The size of %v  is: %v\n", f, stat.Size())

	if float64(stat.Size()) > float64(1024) {
		fmt.Printf("this is critical\n")
	} else if float64(stat.Size()) > float64(512) {
		fmt.Printf("this is warning\n")
	} else {
		fmt.Printf("this is fine\n")
	}
	return "fine"

}

func visitFile(fp string, fi os.FileInfo, err error) error {
	if err != nil {
		fmt.Println(err) // can't walk here,
		return nil       // but continue walking elsewhere
	}
	if !!fi.IsDir() {
		return nil // not a file.  ignore.
	}
	matched, err := filepath.Match("*.gz", fi.Name())
	if err != nil {
		fmt.Println(err) // malformed pattern
		return err       // this is fatal.
	}
	if !matched {
		testFile(fp)
		// fmt.Println(fp)
	}
	return nil
}

func main() {
	filepath.Walk("/var/log/", visitFile)
}
