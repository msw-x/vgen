package vgen

import (
	"errors"
	"os/exec"
	"strings"
	"time"

	"msw/moon"
)

func HasRepository(path string) bool {
	return moon.PathExist(moon.PathJoin(path, ".git"))
}

func Gen(path string) string {
	if !HasRepository(path) {
		moon.Panicf("not found git repository: %s", path)
	}
	cmd := gitCmd{path: path}
	status := cmd.repoStatus()
	components := []string{formatTime(status.time)}
	if status.branch != "master" && status.branch != "main" {
		if status.branch == "HEAD (no branch)" {
			//
		} else {
			components = append(components, status.branch)
		}
	}
	if !status.pure {
		components = append(components, formatRelativeTime(status.time))
	}
	return strings.Join(components, "-")
}

func GenQuiet(path string) (v string, err error) {
	defer moon.Recover(func(e string) {
		err = errors.New(e)
	})
	v = Gen(path)
	return
}
