package main

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

func initLogging() {
	// Setup logging
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.JSONFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			s := strings.Split(f.Function, ".")
			funcName := s[len(s)-1]
			return funcName, fmt.Sprintf("%s:%d", path.Base(f.File), f.Line)
		},
	})
	logrus.SetOutput(os.Stdout)

	logrus.Infof("Log level: %s", logrus.GetLevel())
}
