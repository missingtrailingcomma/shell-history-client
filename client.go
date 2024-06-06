package main

import (
	"flag"
	"fmt"
)

var (
	pid  = flag.String("pid", "", "the PID of the process")
	ppid = flag.String("ppid", "", "the PPID of the process")

	debug = flag.Bool("debug", false, "turn on debug mode or not")
)

type commandInfo struct {
	pid  string
	ppid string
}

type envInfo struct {
	debug bool
}

func main() {
	flag.Parse()

	ci := commandInfo{
		pid:  *pid,
		ppid: *ppid,
	}

	ei := envInfo{
		debug: *debug,
	}

	fmt.Printf("debug: %+v %+v\n", ci, ei)
}
