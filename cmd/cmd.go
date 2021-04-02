package cmd

import (
	"fmt"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
)

// ExecuteCmd is a wrapper for os.exec()
func ExecuteCmd(cmd string) []byte {
	if runtime.GOOS == "windows" {
		//cmd = exec.Command("tasklist")
		fmt.Println("You need to switch to Linux, stoopid!")
	}
	cmdString := "\"$ " + cmd + "\""
	fmt.Println("\nExecuting ", cmdString)
	output, err := exec.Command("bash", "-c", cmd).CombinedOutput()
	if err != nil {
		fmt.Println("\n", fmt.Sprint(err)+": "+string(output))
	} else {
		if string(output) != "" {
			fmt.Println("\nOutput=>", string(output))
		}
		fmt.Println("\n\u2713\u2713\u2713\u2713\u2713\u2713Ran successfully\u2713\u2713\u2713\u2713\u2713\u2713")
		fmt.Println("-----")
	}
	return output
}

func ParseCmdOutput(output []byte, parseType string, reg string, matchGr int) interface{} {
	match := regexp.MustCompile(reg).FindStringSubmatch(string(output))
	var result interface{}
	if parseType == "float" {
		if match != nil {
			if i, err := strconv.ParseFloat(match[matchGr], 64); err == nil {
				result = i
			}
		}
	} else if parseType == "string" {
		if match != nil {
			result = match[matchGr]
		}
	}
	return result
}
