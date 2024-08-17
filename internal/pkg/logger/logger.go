package logger

import (
	"fmt"
	"log"
	"os"
)

var logger log.Logger

func init() {
	logger = *log.New(os.Stdout, "", log.LstdFlags)
}

func Infof(format string, args ...interface{}) {
	logger.Printf("Info: %s", fmt.Sprintf(format, args...))
}

func Errorf(format string, args ...interface{}) {
	logger.Printf("Error: %s", fmt.Sprintf(format, args...))
}

func Fatalf(format string, args ...interface{}) {
	logger.Fatalf("Fatal: %s", fmt.Sprintf(format, args...))
}
