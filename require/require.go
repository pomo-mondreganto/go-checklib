package require

import (
	"errors"
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"

	"github.com/pomo-mondreganto/go-checklib"
	o "github.com/pomo-mondreganto/go-checklib/require/options"

	"github.com/google/go-cmp/cmp"
)

func Equal(c *checklib.C, expected, actual any, public string, opts ...o.Option) {
	if err := validateEqual(expected, actual); err != nil {
		c.Finish(
			checklib.VerdictCheckFailed,
			"error in checker",
			fmt.Sprintf(
				"Invalid operation: %#v == %#v (%s)",
				expected,
				actual,
				err,
			),
		)
	}
	if diff := cmp.Diff(expected, actual); diff != "" {
		info := o.GetExitInfo(
			public,
			fmt.Sprintf("%s: Equal: %s", getCaller(1), diff),
			opts...,
		)
		c.Finish(info.Verdict, info.Public, info.Private)
	}
}

func Error(c *checklib.C, err error, public string, opts ...o.Option) {
	if err == nil {
		info := o.GetExitInfo(
			public,
			fmt.Sprintf("%s: expected error, got: %v", getCaller(1), err),
			opts...,
		)
		c.Finish(info.Verdict, info.Public, info.Private)
	}
}

func NoError(c *checklib.C, err error, public string, opts ...o.Option) {
	if err != nil {
		info := o.GetExitInfo(
			public,
			fmt.Sprintf("%s: unexpected error: %v", getCaller(1), err),
			opts...,
		)
		c.Finish(info.Verdict, info.Public, info.Private)
	}
}

func getCaller(depth int) string {
	_, file, lineNo, ok := runtime.Caller(depth + 1)
	if !ok {
		return "<unknown>:<0>"
	}
	return fmt.Sprintf("%s:%d", filepath.Base(file), lineNo)
}

func validateEqual(expected, actual any) error {
	if expected == nil && actual == nil {
		return nil
	}

	if isFunction(expected) || isFunction(actual) {
		return errors.New("cannot take func type as argument")
	}
	return nil
}

func isFunction(arg interface{}) bool {
	if arg == nil {
		return false
	}
	return reflect.TypeOf(arg).Kind() == reflect.Func
}
