package main

import (
	stdlog "log"
	"strings"
	"testing"

	"github.com/go-logr/stdr"
	"github.com/stretchr/testify/assert"
)

type testLogStruct struct {
	level   int
	message string
}

var testLogs = []struct {
	verbosity int
	log       []testLogStruct
}{
	{1,
		[]testLogStruct{
			{0, "CRITICAL(?)"},
			{1, "INFO"},
			{3, "DEBUG"},
		},
	},
	{3,
		[]testLogStruct{
			{0, "CRITICAL(?)"},
			{1, "INFO"},
			{3, "DEBUG"},
		},
	},
}

func TestLogVerbosity(t *testing.T) {
	for _, l := range testLogs {
		stdr.SetVerbosity(l.verbosity)
		var logBuf strings.Builder
		log := stdr.New(stdlog.New(&logBuf, "", stdlog.Lshortfile))
		for _, le := range l.log {
			log.V(le.level).Info(le.message)
		}

		if l.verbosity > 1 {
			assert.Contains(t, logBuf.String(), "DEBUG")
		} else {
			assert.NotContains(t, logBuf.String(), "DEBUG")
		}
	}
}
