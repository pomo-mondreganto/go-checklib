package o

import (
	"fmt"

	"github.com/pomo-mondreganto/go-checklib"
)

type ExitInfo struct {
	Verdict checklib.Verdict
	Public  string
	Private string
}

func GetExitInfo(public, private string, opts ...Option) *ExitInfo {
	info := &ExitInfo{
		Verdict: checklib.VerdictMumble,
		Public:  public,
		Private: private,
	}
	for _, opt := range opts {
		opt(info)
	}
	return info
}

type Option func(info *ExitInfo)

func Private(s string) Option {
	return func(info *ExitInfo) {
		info.Private = s
	}
}

func Privatef(format string, args ...any) Option {
	return func(info *ExitInfo) {
		info.Private = fmt.Sprintf(format, args...)
	}
}

func OK() Option {
	return func(info *ExitInfo) {
		info.Verdict = checklib.VerdictOK
	}
}

func Mumble() Option {
	return func(info *ExitInfo) {
		info.Verdict = checklib.VerdictMumble
	}
}

func Corrupt() Option {
	return func(info *ExitInfo) {
		info.Verdict = checklib.VerdictCorrupt
	}
}

func Down() Option {
	return func(info *ExitInfo) {
		info.Verdict = checklib.VerdictDown
	}
}

func CheckFailed() Option {
	return func(info *ExitInfo) {
		info.Verdict = checklib.VerdictCheckFailed
	}
}
