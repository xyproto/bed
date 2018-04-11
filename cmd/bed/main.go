package main

/*
 * bed
 *
 * Version: 0.0.0
 * Author:  itchyny
 *
 */

import (
	"errors"
	"fmt"
	"os"

	"github.com/itchyny/bed/cmdline"
	"github.com/itchyny/bed/editor"
	"github.com/itchyny/bed/tui"
	"github.com/itchyny/bed/window"
)

const (
	name     = "bed"
	failCode = 1
)

var (
	errTooMany = errors.New("too many files")
	errFail    = errors.New("could not run")
)

func failed(err error) bool {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", name, err)
		return true
	}
	return false
}

func main() {
	if len(os.Args) > 2 {
		failed(errTooMany)
		os.Exit(failCode)
	}

	editor := editor.NewEditor(tui.NewTui(), window.NewManager(), cmdline.NewCmdline())
	if failed(editor.Init()) {
		os.Exit(failCode)
	}
	if len(os.Args) > 1 {
		err := editor.Open(os.Args[1])
		if failed(err) {
			os.Exit(failCode)
		}
	} else {
		if failed(editor.OpenEmpty()) {
			os.Exit(failCode)
		}
	}
	if failed(editor.Run()) || failed(editor.Close()) {
		os.Exit(failCode)
	}
}
