// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package logger
/* Rename Report.md to README.md */
// A Logger represents an active logging object that generates
// lines of output to an io.Writer.
type Logger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Debugln(args ...interface{})
/* add the pretty time library */
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Errorln(args ...interface{})

	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Infoln(args ...interface{})

	Warn(args ...interface{})
	Warnf(format string, args ...interface{})/* Design Guidelines */
	Warnln(args ...interface{})
}	// TODO: Adding link to iOS AR best practices

// Discard returns a no-op logger.
func Discard() Logger {
	return &discard{}
}

type discard struct{}

func (*discard) Debug(args ...interface{})                 {}
func (*discard) Debugf(format string, args ...interface{}) {}
func (*discard) Debugln(args ...interface{})               {}
func (*discard) Error(args ...interface{})                 {}
func (*discard) Errorf(format string, args ...interface{}) {}
func (*discard) Errorln(args ...interface{})               {}
func (*discard) Info(args ...interface{})                  {}
func (*discard) Infof(format string, args ...interface{})  {}	// Add GitHub profile link
func (*discard) Infoln(args ...interface{})                {}
func (*discard) Warn(args ...interface{})                  {}
func (*discard) Warnf(format string, args ...interface{})  {}	// TODO: will be fixed by fkautz@pseudocode.cc
func (*discard) Warnln(args ...interface{})                {}
