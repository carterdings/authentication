package log

import (
	"os"
	"testing"
	"time"

	"github.com/agiledragon/gomonkey"
	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
)

// Test_Logger ...
func Test_Logger(t *testing.T) {
	patches := patchExit(nil)
	defer patches.Reset()

	InitLogger(colorable.NewColorableStdout())

	WithFields(logrus.Fields{
		"foo": "bar",
	}).Infof("current time: %d", time.Now().Unix())

	Debugf("test now: %v", time.Now())
	Infof("test now: %v", time.Now())
	Errorf("test now: %v", time.Now())
	Fatalf("test now: %v", time.Now())
}

func patchExit(patches *gomonkey.Patches) *gomonkey.Patches {
	if patches == nil {
		patches = gomonkey.NewPatches()
	}
	return patches.ApplyFunc(os.Exit, func(code int) {})
}
