package checklib

import (
	"fmt"
	"os"
)

type Verdict string

const (
	VerdictOK          Verdict = "OK"
	VerdictMumble      Verdict = "MUMBLE"
	VerdictCorrupt     Verdict = "CORRUPT"
	VerdictDown        Verdict = "DOWN"
	VerdictCheckFailed Verdict = "CHECK FAILED"
)

func (v Verdict) Code() int {
	switch v {
	case VerdictOK:
		return 101
	case VerdictMumble:
		return 102
	case VerdictCorrupt:
		return 103
	case VerdictDown:
		return 104
	default:
		return 110
	}
}

type status struct {
	verdict Verdict
	public  string
	private string
}

func (s *status) write() {
	_, _ = fmt.Fprintf(os.Stdout, s.public+"\n")
	_, _ = fmt.Fprintf(os.Stderr, s.private+"\n")
}

func OK(c *C, public, private string, privateArgs ...any) {
	c.Finish(VerdictOK, public, getPrivate(public, private, privateArgs...))
}

func Mumble(c *C, public, private string, privateArgs ...any) {
	c.Finish(VerdictMumble, public, getPrivate(public, private, privateArgs...))
}

func Corrupt(c *C, public, private string, privateArgs ...any) {
	c.Finish(VerdictCorrupt, public, getPrivate(public, private, privateArgs...))
}

func Down(c *C, public, private string, privateArgs ...any) {
	c.Finish(VerdictDown, public, getPrivate(public, private, privateArgs...))
}

func CheckFailed(c *C, public, private string, privateArgs ...any) {
	c.Finish(VerdictCheckFailed, public, getPrivate(public, private, privateArgs...))
}

func getPrivate(public, private string, privateArgs ...any) string {
	if private == "" {
		return public
	}
	return fmt.Sprintf(private, privateArgs...)
}
