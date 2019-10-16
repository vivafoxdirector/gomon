package main

import (
	"github.com/fatih/color"
	"github.com/k0kubun/go-ansi"
	"github.com/mitchellh/colorstring"
)

func go_ansi_test() {
	ansi.Print("aaa")
}

func color_test() {
	color.Output = ansi.NewAnsiStdout()
	color.Cyan("fatih/color")
}

func colorstring_test() {
	colorstring.Fprintln(ansi.NewAnsiStdout(), "[green]mitchellh/colorstring")
}
