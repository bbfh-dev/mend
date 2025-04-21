package main

import (
	"os"

	"github.com/bbfh-dev/mend/cli"
	"github.com/bbfh-dev/parsex/v2"
)

var CLI = parsex.Program{
	Data: &cli.Options,
	Name: "mend",
	Desc: "Mend is a simple HTML template processor designed to, but not limited to be used to generate static websites",
	Exec: cli.Main,
}.Runtime().
	SetVersion("0.2.0-alpha").
	SetPosArgs("html files...")

func main() {
	err := CLI.Run(os.Args[1:])
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
	}
}
