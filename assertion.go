package eqtest

import (
	"github.com/adzil/eqtest/internal/indentstr"
	"github.com/adzil/eqtest/internal/slices"
	"github.com/google/go-cmp/cmp"
)

type T interface {
	Helper()
	Log(args ...any)
	Fail()
	FailNow()
}

type diffFunc func(expected, actual any, opts ...cmp.Option) (diff string)

// Assertion implements equality test assertions using go-cmp.
type Assertion struct {
	t    T
	opts []cmp.Option
	diff diffFunc
}

func (a *Assertion) clone() *Assertion {
	na := *a
	return &na
}

// New returns a new [Assertion] using a instance of [*testing.T].
func New(t T, opts ...cmp.Option) *Assertion {
	return &Assertion{
		t:    t,
		opts: opts,
		diff: cmp.Diff,
	}
}

// Using returns a copy of [Assertion] that uses a different [*testing.T].
func (a *Assertion) Using(t T) *Assertion {
	na := a.clone()
	na.t = t
	return na
}

// With adds new [cmp.Option] to the [Assertion] options.
func (a *Assertion) With(opts ...cmp.Option) *Assertion {
	na := a.clone()
	na.opts = slices.Merge(a.opts, opts)
	return na
}

// Equal asserts the equality of expected and actual.
func (a *Assertion) Equal(expected, actual any, opts ...cmp.Option) bool {
	a.t.Helper()

	opts = slices.Merge(a.opts, opts)

	diff := a.diff(expected, actual, opts...)
	if diff == "" {
		return true
	}

	var buf indentstr.Builder
	buf.WriteString("\nError:  Not equal")
	buf.WriteString("\nDiff:   ")
	buf.SetIndent(8)
	buf.WriteString(diff)

	a.t.Log(buf.String())
	a.t.Fail()
	return false
}

// MustEqual asserts the equality of expected and actual, and immediately stop
// the test execution on failure.
func (a *Assertion) MustEqual(expected, actual any, opts ...cmp.Option) bool {
	a.t.Helper()

	if a.Equal(expected, actual, opts...) {
		return true
	}

	a.t.FailNow()
	return false
}

// Equal is a convenience function for [New] followed by an [Assertion.Equal].
func Equal(t T, expected, actual any, opts ...cmp.Option) bool {
	t.Helper()
	return New(t).Equal(expected, actual, opts...)
}

// MustEqual is a convenience function for [New] followed by a
// [Assertion.MustEqual].
func MustEqual(t T, expected, actual any, opts ...cmp.Option) bool {
	t.Helper()
	return New(t).MustEqual(expected, actual, opts...)
}
