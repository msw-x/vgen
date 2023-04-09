package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/msw-x/moon/app"
	"github.com/msw-x/vgen"
)

func main() {
	err := run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	args := app.Args()
	if len(args) == 0 {
		return errors.New("language not specified")
	}
	lang := args[0]
	srcDir := "./"
	dstDir := "./"
	if len(args) > 1 {
		dstDir = args[1]
	}
	v, err := vgen.Gen(srcDir)
	if err != nil {
		return err
	}
	switch lang {
	case "go":
		return vgen.NewGo().GenFile(v, dstDir)
	}
	return fmt.Errorf("incorrect language: %s", lang)
}
