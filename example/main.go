package main

import (
	"errors"
	"os"

	"github.com/pomo-mondreganto/go-checklib"
	"github.com/pomo-mondreganto/go-checklib/require"
	o "github.com/pomo-mondreganto/go-checklib/require/options"
)

type Checker struct{}

func (ch *Checker) Info() *checklib.CheckerInfo {
	return &checklib.CheckerInfo{
		Vulns:      1,
		Timeout:    10,
		AttackData: true,
		Puts:       1,
		Gets:       1,
	}
}

func (ch *Checker) Check(c *checklib.C, _ string) {
	require.Equal(c, "a", "a", "bad cmp", o.Corrupt())
	require.Error(c, errors.New("kek"), "bad error", o.Mumble())
	checklib.OK(c, "OK", "")
}

func (ch *Checker) Put(c *checklib.C, _, flagID, flag string, _ int) {
	checklib.OK(c, flag, flagID)
}

func (ch *Checker) Get(c *checklib.C, _, _, _ string, _ int) {
	checklib.OK(c, "OK", "OK")
}

func main() {
	os.Exit(checklib.Run(&Checker{}))
}
