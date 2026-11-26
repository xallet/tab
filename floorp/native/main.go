package main

import (
	firefox "xallet.app/tab/firefox/engine"
	"os"
)

func main() {
	config := firefox.Config{
		Name: "floorp", MacProfileRoot: "Floorp", WindowsProfileRoot: "Floorp", LinuxProfileRoot: ".floorp",
	}
	os.Exit(firefox.Run(config, os.Args[1:], os.Stdin, os.Stdout, os.Stderr))
}
