package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegex(t *testing.T) {
	var tests = []struct {
		fqdn        string
		hostname    string
		nodeSetName string
	}{
		{"cv-stg-al002-1-007-ct-jp2v-dev.lineinfra-dev.com",
			"cv-stg-al002-1-007-ct-jp2v-dev",
			"al002-1",
		},
		{"cv-test-test-001-0-001-ct-jp2v-dev.lineinfra-dev.com",
			"cv-test-test-001-0-001-ct-jp2v-dev",
			"test-001-0",
		},
		{"cv-stg-gonz-n-prod-1-005-nucleo-jp2v-dev.lineinfra-dev.com",
			"cv-stg-gonz-n-prod-1-005-nucleo-jp2v-dev",
			"gonz-n-prod-1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.fqdn, func(t *testing.T) {
			hostname, nodeSetName := parse(tt.fqdn)
			assert.Equal(t, tt.hostname, hostname)
			assert.Equal(t, tt.nodeSetName, nodeSetName)
		})
	}
}

func parse(fqdn string) (hostname string, nodeSetName string) {
	hostname = strings.Split(fqdn, ".")[0]

	var nodeSetNameFrag []string
	hostnameFrag := strings.Split(hostname, "-")
	for i := 2; i < len(hostnameFrag)-4; i++ {
		nodeSetNameFrag = append(nodeSetNameFrag, hostnameFrag[i])
	}
	nodeSetName = strings.Join(nodeSetNameFrag, "-")
	return
}

func TestSplit(t *testing.T) {
	src := `al003
al004
al005
al006
al007
al007-dev
al007-prod
al008
al008-dev
al008-prod`
	rr := strings.Split(src, "\n")
	assert.Equal(t, []string{
		"al003",
		"al004",
		"al005",
		"al006",
		"al007",
		"al007-dev",
		"al007-prod",
		"al008",
		"al008-dev",
		"al008-prod",
	}, rr)
}
