package vgen

import (
	"os/exec"
	"strings"
	"time"
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

func (this *gitCmd) exec(arg ...string) string {
	cmd := exec.Command("git", arg...)
	cmd.Dir = this.path
	out, err := cmd.CombinedOutput()
	s := string(out)
	moon.TrimSuffix(&s, "\n")
	moon.Check(err, s)
	return s
}

func (this *gitCmd) status() (branch string, pure bool) {
	s := this.exec("status", "-s", "-b")
	lines := strings.Split(s, "\n")
	if len(lines) == 0 {
		moon.Panic("git status fail: lines is empty")
	}
	branch = lines[0]
	if !moon.TrimPrefix(&branch, "## ") {
		moon.Panicf("git status fail: branch not found: %s", branch)
	}
	branch, _ = moon.SplitPair(branch, "...")
	lines = lines[1:]
	pure = len(lines) == 0
	return
}

func (this *gitCmd) lastHash(all bool) string {
	cmd := []string{"log", "-1", `--pretty=format:"%h"`}
	if all {
		cmd = append(cmd, "--all")
	}
	return this.exec(cmd...)
}

func (this *gitCmd) lastTime() time.Time {
	cmd := []string{"log", "-1", `--pretty=format:"%ct"`}
	i := moon.ToInt(moon.TrimQuotesThru(this.exec(cmd...)))
	return time.Unix(int64(i), 0)
}

func (this *gitCmd) lastComit() bool {
	return this.lastHash(false) == this.lastHash(true)
}

func (this *gitCmd) repoStatus() repoStatus {
	branch, pure := this.status()
	return repoStatus{
		branch: branch,
		time:   this.lastTime(),
		pure:   pure,
		late:   this.lastComit(),
	}
}
