package main

import (
	firefox "xallet.app/tab/firefox/engine"
	"os"
)

func main() {
	config := firefox.Config{
		Name: "librewolf", MacProfileRoot: "librewolf", WindowsProfileRoot: "librewolf", LinuxProfileRoot: ".librewolf",
	}
	os.Exit(firefox.Run(config, os.Args[1:], os.Stdin, os.Stdout, os.Stderr))
}
