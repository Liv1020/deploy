package main

import (
	"bufio"
	"io"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var lock = new(OptimisticLock)

// DeployHandler DeployHandler
func DeployHandler(c *gin.Context) {
	if c.Request.Method == "OPTIONS" {
		return
	}

	path := c.Query("path")
	if strings.Contains(path, "favicon.ico") {
		return
	}

	if !lock.Lock() {
		c.String(http.StatusOK, "Already has task, wait a moment. Or see <a href=\"http://%s/log\">deploy log</a>", c.Request.Host)
		return
	}

	InitLogFile()

	if runtime.GOOS == "windows" {
		go updateGitFolder(path[1:])
	} else {
		go updateGitFolder(path)
	}

	c.String(http.StatusOK, "ok")
}

func updateGitFolder(path string) {
	defer func() {
		lock.UnLock()
	}()

	Println("======================= Date ========================")
	now := time.Now()
	Println(now.Format(time.RFC3339))
	Println("======================= Path ========================")
	Println(path)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		Println("Not exists! Ignored")
		return
	}
	if err := os.Chdir(path); err != nil {
		Println("Access denied")
		return
	}
	if _, err := os.Stat(".git"); err == nil {
		Println("================ Revert all changes =================")
		runCommand(exec.Command("git", "reset", "--hard", "HEAD"))
		Println("===================== Pulling =======================")
		runCommand(exec.Command("git", "pull"))
	}

	if _, err := os.Stat("build.sh"); err == nil {
		Println("========= Running external build.sh script ==========")
		runCommand(exec.Command("./build.sh"))
	}

	if _, err := os.Stat("deploy.sh"); err == nil {
		Println("========= Running external deploy.sh script =========")
		runCommand(exec.Command("./deploy.sh"))
	}

	Println("======================== Done =======================")
	Println("")
}

func runCommand(cmd *exec.Cmd) {
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		Println(err)
		return
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		Println(err)
		return
	}

	if err := cmd.Start(); err != nil {
		Println(err)
		return
	}

	outReader := bufio.NewReader(stdout)
	errReader := bufio.NewReader(stderr)

	go func() {
		for {
			line, err2 := outReader.ReadString('\n')
			if err2 != nil || io.EOF == err2 {
				break
			}
			Print(line)
		}
	}()

	go func() {
		for {
			line, err2 := errReader.ReadString('\n')
			if err2 != nil || io.EOF == err2 {
				break
			}
			Print(line)
		}
	}()

	if err := cmd.Wait(); err != nil {
		Println(err)
		return
	}
}
