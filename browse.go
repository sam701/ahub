package main

import (
	"bytes"
	"os/exec"

	"strings"

	"fmt"

	"github.com/urfave/cli"
)

func Browse(ctx *cli.Context) error {
	repo := getCurrentUser(getCurrentPullUrl())
	suffix := ""
	if user := ctx.String("user"); user != "" {
		repo.user = user
	}
	if ctx.Bool("pulls") {
		suffix = "/pulls"
	} else if ctx.Bool("issues") {
		suffix = "/issues"
	}

	openBrowser(fmt.Sprintf("https://github.com/%s/%s%s", repo.user, repo.repo, suffix))
	return nil
}

type Repo struct {
	user string
	repo string
}

func getCurrentUser(pullUrl string) *Repo {
	start := strings.LastIndex(pullUrl, ":")
	end := strings.LastIndex(pullUrl, "/")
	return &Repo{
		user: pullUrl[start+1 : end],
		repo: strings.Replace(pullUrl[end+1:], ".git", "", 1),
	}
}

func getCurrentPullUrl() string {
	cmd := exec.Command("git", "remote", "get-url", "origin")

	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Run()

	return strings.TrimSpace(string(buf.Bytes()))
}

func openBrowser(url string) {
	exec.Command("open", url).Run()
}
