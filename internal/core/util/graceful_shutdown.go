package util

import (
	"os"
	"syscall"
)

var InterruptSignals = []os.Signal{
	os.Interrupt,
	syscall.SIGTERM,
	syscall.SIGINT,
}