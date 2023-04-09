package vgen

import (
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/msw-x/moon/ustring"
)

type gitCmd struct {
	path string
}

type repoStatus struct {
	branch string
	time   time.Time
	pure   bool
	late   bool
}

func (o *gitCmd) exec(arg ...string) (string, error) {
	cmd := exec.Command("git", arg...)
	cmd.Dir = o.path
	out, err := cmd.CombinedOutput()
	s := string(out)
	s = strings.TrimSuffix(s, "\n")
	return s, err
}

func (o *gitCmd) status() (branch string, pure bool, err error) {
	var s string
	s, err = o.exec("status", "-s", "-b")
	if err != nil {
		return
	}
	lines := strings.Split(s, "\n")
	if len(lines) == 0 {
		err = errors.New("git status fail: lines is empty")
		return
	}
	branch = lines[0]
	prefix := "## "
	if strings.HasPrefix(branch, prefix) {
		branch = strings.TrimPrefix(branch, prefix)
	} else {
		err = fmt.Errorf("git status fail: branch not found: %s", branch)
		return
	}
	branch, _ = ustring.SplitPair(branch, "...")
	lines = lines[1:]
	pure = len(lines) == 0
	return
}

func (o *gitCmd) lastHash(all bool) (string, error) {
	cmd := []string{"log", "-1", `--pretty=format:"%h"`}
	if all {
		cmd = append(cmd, "--all")
	}
	return o.exec(cmd...)
}

func (o *gitCmd) lastTime() (t time.Time, err error) {
	cmd := []string{"log", "-1", `--pretty=format:"%ct"`}
	var s string
	s, err = o.exec(cmd...)
	if err == nil {
		var i int
		i, err = strconv.Atoi(ustring.TrimQuotes(s))
		t = time.Unix(int64(i), 0)
	}
	return
}

func (o *gitCmd) lastComit() (yes bool, err error) {
	var h1, h2 string
	h1, err = o.lastHash(false)
	if err == nil {
		h2, err = o.lastHash(true)
		if err == nil {
			yes = h1 == h2
		}
	}
	return
}

func (o *gitCmd) repoStatus() (s repoStatus, err error) {
	s.branch, s.pure, err = o.status()
	if err == nil {
		s.time, err = o.lastTime()
		if err == nil {
			s.late, err = o.lastComit()
		}
	}
	return
}
