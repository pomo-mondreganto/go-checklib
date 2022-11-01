package checklib

import (
	"context"
	"encoding/json"
	"os"
	"strconv"
	"time"
)

// Run returns exit code.
func Run(checker Checker) int {
	verdict := make(chan Verdict, 1)

	info := checker.Info()

	// Separate goroutine terminated by runtime.Goexit.
	go func() {
		ctx, cancel := context.WithTimeout(
			context.Background(),
			time.Second*time.Duration(info.Timeout),
		)
		defer cancel()

		c := &C{Context: ctx}

		defer func() {
			s := &status{
				verdict: c.verdict,
				public:  c.public,
				private: c.private,
			}
			if !c.finished {
				s = &status{
					verdict: VerdictCheckFailed,
					public:  "error in checker",
					private: "checker did not report status",
				}
			}
			s.write()
			verdict <- s.verdict
		}()

		if len(os.Args) < 2 {
			CheckFailed(c, "error calling checker", "missing command")
		}

		switch cmd := os.Args[1]; cmd {
		case "info":
			data, _ := json.Marshal(info)
			c.Finish(VerdictOK, string(data), "")
		case "check":
			if len(os.Args) < 3 {
				CheckFailed(c, "error calling checker", "missing host")
			}
			checker.Check(c, os.Args[2])
		case "put":
			if len(os.Args) < 6 {
				CheckFailed(c, "error calling checker", "missing arguments")
			}

			vuln, err := strconv.Atoi(os.Args[5])
			if err != nil {
				CheckFailed(c, "error calling checker", "invalid vuln number")
			}

			checker.Put(c, os.Args[2], os.Args[3], os.Args[4], vuln)
		case "get":
			if len(os.Args) < 6 {
				CheckFailed(c, "error calling checker", "missing arguments")
			}

			vuln, err := strconv.Atoi(os.Args[5])
			if err != nil {
				CheckFailed(c, "error calling checker", "invalid vuln number")
			}

			checker.Get(c, os.Args[2], os.Args[3], os.Args[4], vuln)
		default:
			CheckFailed(c, "error calling checker", "bad command: %s", cmd)
		}
	}()

	return (<-verdict).Code()
}
