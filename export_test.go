package eqtest

import "github.com/google/go-cmp/cmp"

type DiffFunc = diffFunc

type AssertionParams struct {
	T    T
	Opts []cmp.Option
	Diff DiffFunc
}

func NewFrom(params AssertionParams) *Assertion {
	return &Assertion{
		t:    params.T,
		opts: params.Opts,
		diff: params.Diff,
	}
}
