package checklib

import (
	"context"
	"runtime"
	"sync"
)

// C is like testing.T, the "core" object.
type C struct {
	context.Context

	mu sync.Mutex

	verdict  Verdict
	public   string
	private  string
	finished bool
}

func (c *C) SetPublic(s string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.public = s
}

func (c *C) SetPrivate(s string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.private = s
}

func (c *C) Finish(verdict Verdict, public, private string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.verdict = verdict
	c.public = public
	c.private = private
	c.finished = true

	runtime.Goexit()
}
