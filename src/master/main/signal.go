package main

import (
	"os"
	"os/signal"
	"syscall"
	logger "github.com/shengkehua/xlog4go"
	"server/httpserver"
	"time"
)

func signal_proc() {
	c := make(chan os.Signal, 1)

	signal.Notify(c, syscall.SIGINT, syscall.SIGALRM, syscall.SIGTERM, syscall.SIGUSR1)

	// Block until a signal is received.
	sig := <-c

	logger.Warn("Signal received: %v", sig)

	httpserver.HttpListener.Close()
	for _, handler := range httpserver.Uri2Handler {
		handler.Close()
	}

	time.Sleep(500 * time.Millisecond)

	logger.Warn("send quit signal")
	G_QuitChan <- 1
}