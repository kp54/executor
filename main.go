package main

import (
	"io"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func main() {
	if len(os.Args) < 3 {
		os.Stdout.WriteString("usage: executor TARGET EXIT [ARGS...]\n")
		return
	}

	cmd := exec.Command(os.Args[1], os.Args[3:]...)

	cstdin, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}
	defer cstdin.Close()
	cstdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	defer cstdout.Close()
	cstderr, err := cmd.StderrPipe()
	if err != nil {
		panic(err)
	}
	defer cstderr.Close()

	// stdin, stdout, stderr をつなぐ (こんな雑でいいのか?)
	go func() {
		for {
			io.Copy(cstdin, os.Stdin)
		}
	}()
	go func() {
		for {
			io.Copy(os.Stdout, cstdout)
		}
	}()
	go func() {
		for {
			io.Copy(os.Stderr, cstderr)
		}
	}()

	// SIGTERM をトラップ
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM)

	// onSIGTERM
	go func() {
		for {
			<-ch
			// os.Stdin がバッファされてる前提 (何か入ってたら死ぬよねこれ)
			io.WriteString(cstdin, os.Args[2]+"\n")
		}
	}()

	cmd.Run()
}
