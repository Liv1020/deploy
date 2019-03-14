package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

var out *os.File

func logFileName() string {
	now := time.Now()
	return fmt.Sprintf("%s/log/log-%s.txt", baseDir, now.Format("2006-01-02"))
}

// Println Println
func Println(a ...interface{}) {
	_, _ = fmt.Fprintln(out, a...)
}

// Print Print
func Print(a ...interface{}) {
	_, _ = fmt.Fprint(out, a...)
}

// InitLogFile InitLogFile
func InitLogFile() {
	name := logFileName()
	if _, err := os.Stat(name); err != nil {
		dir := filepath.Dir(name)
		if err := os.MkdirAll(dir, 0777); err != nil {
			log.Fatalln(err)
			return
		}

		var err error
		out, err = os.OpenFile(logFileName(), os.O_RDWR|os.O_CREATE, 0777)
		if err != nil {
			log.Fatalln(err)
			return
		}
	}
}

// LogHandler LogHandler
func LogHandler(c *gin.Context) {
	c.File(logFileName())
}

var baseDir string

func init() {
	var err error
	baseDir, err = filepath.Abs("./")
	if err != nil {
		log.Fatalln(err)
	}
}
