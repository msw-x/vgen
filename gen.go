package vgen

import (
	"fmt"
	"path"
	"strings"

	"github.com/msw-x/moon/fs"
)

func HasRepository(dir string) bool {
	return fs.Exist(path.Join(dir, ".git"))
}

func Gen(path string) (s string, err error) {
	if !HasRepository(path) {
		err = fmt.Errorf("not found git repository: %s", path)
		return
	}
	cmd := gitCmd{path: path}
	var status repoStatus
	status, err = cmd.repoStatus()
	if err == nil {
		components := []string{fmtTime(status.time)}
		if status.branch != "master" && status.branch != "main" {
			if status.branch == "HEAD (no branch)" {
				//
			} else {
				components = append(components, status.branch)
			}
		}
		if !status.pure {
			components = append(components, fmtRelativeTime(status.time))
		}
		s = strings.Join(components, "-")
	}
	return
}
