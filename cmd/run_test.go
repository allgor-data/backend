package cmd

import (
	"os"
	"syscall"
	"testing"
	"time"

	"gotest.tools/v3/assert"
)

func TestRun(t *testing.T) {
	args := os.Args[0:1]
	args = append(args, "--playground")

	var err error

	go func() {
		err = Run(args)
	}()

	time.Sleep(time.Second * 1)

	syscall.Kill(syscall.Getpid(), syscall.SIGINT)

	assert.NilError(t, err)
}

func TestRun_Help(t *testing.T) {
	args := os.Args[0:1]
	args = append(args, "-h")

	os.Stdout = nil

	err := Run(args)

	assert.NilError(t, err)
}
