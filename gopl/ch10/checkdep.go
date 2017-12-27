package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type ListJson struct {
	Name string
	Deps []string
}

// Check all packages that depends on os.Args[1]
// ./checkdep encoding/xml
func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Please input argument to be checked")
		os.Exit(1)
	}

	// 1. get all packages in the workspace
	absdir, err := filepath.Abs(".")
	if err != nil {
		fmt.Println("get abs failed,err=", err)
		os.Exit(1)
	}
	fmt.Println("abs=", absdir)
	cmd := exec.Command("go", "list", "...")
	cmd.Env = append(os.Environ(), "GOPATH="+absdir)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println("Could not establish stdout pipeline")
		os.Exit(1)
	}

	s := string(stdout)
	all := strings.Split(s, "\n")
	all = all[:len(all)-2] //erase the empty element in array

	// 2. for each package, check if it depends on os.Args[1]
	for _, pkg := range all {
		cmd = exec.Command("go", "list", "-json", pkg)
		cmd.Env = append(os.Environ(), "GOPATH="+absdir)
		stdout, err := cmd.Output()
		if err != nil {
			fmt.Println("Could not get dependency of", pkg)
			continue
		}
		var listjson ListJson
		err = json.Unmarshal(stdout, &listjson)
		if err != nil {
			fmt.Println("Could not Unmarshal result", err)
			continue
		}

		for _, dep := range listjson.Deps {
			if dep == os.Args[1] {
				fmt.Fprintf(os.Stderr, "package %s depends on %s\n", pkg, os.Args[1])
				break
			}
		}
	}
}
