package main

import (
	"fmt"
	"os"

	"github.com/bbfh-dev/mend/cli"
	"github.com/bbfh-dev/parsex/v2"
)

var Runtime = parsex.Program{
	Data: &cli.Options,
	Name: "mend",
	Desc: "HTML template processor designed to, but not limited to be used to generate static websites",
	Exec: cli.Main,
}.Runtime().SetVersion("1.0.1-alpha.2").SetPosArgs("html files...")

func main() {
	err := Runtime.Run(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}
