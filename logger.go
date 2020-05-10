package rmq

import (
	"fmt"
	"log"
	"os"
)

type Logger interface {
	Errf(t string, opts ...interface{})
	Panicf(t string, opts ...interface{})
}

type DefaultLogger struct{}

func (DefaultLogger) Errf(t string, opts ...interface{}) {
	fmt.Fprintf(os.Stderr, t, opts)
}

func (DefaultLogger) Panicf(t string, opts ...interface{}) {
	log.Panicf(t, opts...)
}

var logger Logger = DefaultLogger{}

func SetLogger(log Logger) {
	logger = log
}
