package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
)

var (
	flgHelp   bool
	flgEcho   string
	goVersion string
)

func parseCmdLineFlags() {
	flag.BoolVar(&flgHelp, "help", false, "if true,show help")
	flag.StringVar(&flgEcho, "echo", "李林dji", "")
	flag.StringVar(&goVersion, "go-version", runtime.Version(), "show the go version")
	flag.Parse()
}

func main() {
	parseCmdLineFlags()
	if flgHelp {
		flag.Usage()
		os.Exit(0)
	}
	fmt.Printf("flag -echo: '%s' \n", flgEcho)
	remainArgs := flag.Args()
	for _, arg := range remainArgs {
		fmt.Printf("remaining arg: %s \n", arg)
	}

}
