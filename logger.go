package main

import (
	"os"

	logger "github.com/sirupsen/logrus"
)

var rlog = logger.New()

func init() {
	// Log as JSON instead of the default ASCII formatter.
	rlog.SetFormatter(&logger.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	rlog.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	rlog.SetLevel(logger.InfoLevel)
}
