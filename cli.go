package main

import (
	"fmt"
	"os"
	//	"path"

	"github.com/jawher/mow.cli"
)

func exit() {
	fmt.Println("Here's nothing")
	os.Exit(0)
}

func banner() {
	fmt.Println("BABANNER")
}

// ToFile saves results to given file.
func toFile(filepath string, result interface{}) {
	fmt.Println("Saved to ", filepath)
	/*	dir := path.Dir(filepath)

		if _, err := os.Stat(dir); os.IsNotExist(err) {
			errFatal(err)
		}

		if _, err := os.Stat(filepath); os.IsExist(err) {
			errFatal(err)
		}

		file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0666)
		errFatal(err)
		defer file.Close()

		file.WriteString(result)
		errFatal(err)
	*/
}

func Init() {
	goluhn := cli.App("goluhn", "Luhn checksum calculator")
	goluhn.Spec = "-N | [-v] | [[--output=<filename> FILENAME] | [-bc] COMMAND ARG...[COMMAND ARG...]]"

	// CLI global options
	var (
		output = goluhn.StringOpt("o output", "", "output to file")

		none    = goluhn.BoolOpt("N none", false, "nothing will be shown")
		version = goluhn.BoolOpt("v version", false, "print app version")
		banner  = goluhn.BoolOpt("b banner", false, "print ASCII banner")
		colored = goluhn.BoolOpt("c colored", false, "colored output")
	)

	// Commands
	/*
	   cmd.Spec = "[OPTIONS] [COMAND [OPTIONS]]"
	*/
	goluhn.Command("check", "Check Luhn checksum for given number", func(cmd *cli.Cmd) {
		cmd.Spec = "NUMBER"
		number := cmd.IntArg("number", 0, "integer number to calculate its Luhn")
		cmd.Action = func() {
			ch := CheckLuhn(*number)
			fmt.Sprintf("%t", ch)
		}
	})
	
	if 

	goluhn.Run(os.Args)
}
