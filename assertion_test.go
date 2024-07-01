package eqtest_test

import (
	"fmt"
	"runtime"
	"slices"
	"strings"
	"testing"

	"github.com/adzil/eqtest"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestAssertion_Equal_IsEqual(t *testing.T) {
	mt := NewT(t)
	assert := eqtest.NewFrom(eqtest.AssertionParams{
		T: mt,
		Diff: func(expected, actual any, opts ...cmp.Option) (diff string) {
			return ""
		},
	})

	assert.Equal(nil, nil)

	if mt.Result.Fail || mt.Result.Abort {
		t.Error("fail or fail now is unexpectedly called")
	}
}

func TestAssertion_Equal_IsNotEqual(t *testing.T) {
	mt := NewT(t)
	assert := eqtest.NewFrom(eqtest.AssertionParams{
		T: mt,
		Diff: func(expected, actual any, opts ...cmp.Option) (diff string) {
			return "diff here"
		},
	})

	assert.Equal(nil, nil)

	if !mt.Result.Fail {
		t.Error("fail is not called")
	}

	if mt.Result.Abort {
		t.Error("fail now unexpectedly called")
	}

	var notEqualMsg bool
	for _, log := range mt.Result.Logs {
		if strings.Contains(log, "Not equal") {
			notEqualMsg = true
		}
	}

	if !notEqualMsg {
		t.Error("not equal message not logged")
	}
}

func TestAssertion_MustEqual_IsEqual(t *testing.T) {
	mt := NewT(t)
	assert := eqtest.NewFrom(eqtest.AssertionParams{
		T: mt,
		Diff: func(expected, actual any, opts ...cmp.Option) (diff string) {
			return ""
		},
	})

	assert.MustEqual(nil, nil)

	if mt.Result.Fail || mt.Result.Abort {
		t.Error("fail or fail now is unexpectedly called")
	}
}

func TestAssertion_MustEqual_IsNotEqual(t *testing.T) {
	mt := NewT(t)
	assert := eqtest.NewFrom(eqtest.AssertionParams{
		T: mt,
		Diff: func(expected, actual any, opts ...cmp.Option) (diff string) {
			return "diff here"
		},
	})

	assert.MustEqual(nil, nil)

	if !mt.Result.Abort {
		t.Error("fail now is not called")
	}
}

func TestAssertion_WithOpts(t *testing.T) {
	// Fill expectedOpts with 6 [cmpopts.IgnoreInterfaces] options.
	var expectedOpts []cmp.Option
	for range 6 {
		expectedOpts = append(expectedOpts, cmpopts.IgnoreInterfaces(struct{ error }{}))
	}

	mt := NewT(t)
	assert := eqtest.NewFrom(eqtest.AssertionParams{
		T: mt,
		Diff: func(expected, actual any, opts ...cmp.Option) (diff string) {
			if !slices.Equal(expectedOpts, opts) {
				t.Error("cmp.Option given is not equal")
			}

			return ""
		},
		Opts: expectedOpts[0:2],
	})

	assert.With(expectedOpts[2:4]...).
		Equal(nil, nil, expectedOpts[4:6]...)

	if mt.Result.Fail || mt.Result.Abort {
		t.Error("fail or fail now is unexpectedly called")
	}
}

func TestAssertion_Using(t *testing.T) {
	mt := NewT(t)
	mut := NewT(t)
	assert := eqtest.NewFrom(eqtest.AssertionParams{
		T: mt,
		Diff: func(expected, actual any, opts ...cmp.Option) (diff string) {
			return "diff here"
		},
	})

	assert.Using(mut).
		Equal(nil, nil)

	if mt.Result.Fail || mt.Result.Abort {
		t.Error("fail or fail now is unexpectedly called on mt")
	}

	if !mut.Result.Fail {
		t.Error("fail is not called on mut")
	}
}

func TestEqual(t *testing.T) {
	// Just a simple test to ensure that the Equal call is working.
	eqtest.Equal(t, "hello", "hello")
}

func TestMustEqual(t *testing.T) {
	// Just a simple test to ensure that the MustEqual call is working.
	eqtest.MustEqual(t, "hello", "hello")
}

type TResult struct {
	Logs  []string
	Fail  bool
	Abort bool
}

type T struct {
	Result  TResult
	t       eqtest.T
	helpers map[string]struct{}
}

func getCallerFunc(skip int) string {
	var pcs [1]uintptr
	runtime.Callers(skip+1, pcs[:]) // Skip Callers function.

	frames := runtime.CallersFrames(pcs[:])
	frame, _ := frames.Next()

	if frame.Function == "" {
		panic("unable to identify the function name of the caller")
	}

	return frame.Function
}

func NewT(t eqtest.T) *T {
	return &T{t: t}
}

func (t *T) Helper() {
	t.t.Helper()

	if t.helpers == nil {
		t.helpers = make(map[string]struct{})
	}

	fnName := getCallerFunc(2) // Skip getCallerPC and Helper function.
	t.helpers[fnName] = struct{}{}
}

func (t *T) checkHelper() {
	t.t.Helper()

	fnName := getCallerFunc(3) // Skip the caller of this function.
	if _, ok := t.helpers[fnName]; !ok {
		t.t.Log("method called without helper:", fnName)
		t.t.Fail()
	}
}

func (t *T) Fail() {
	t.t.Helper()
	t.checkHelper()

	t.Result.Fail = true
}

func (t *T) FailNow() {
	t.t.Helper()
	t.checkHelper()

	t.Result.Fail = true
	t.Result.Abort = true
}

func (t *T) Log(args ...any) {
	t.t.Helper()
	t.checkHelper()

	t.Result.Logs = append(t.Result.Logs, fmt.Sprint(args...))
}
