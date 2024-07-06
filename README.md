# Eqtest - Equality test assertions with go-cmp

[![Go Reference](https://pkg.go.dev/badge/github.com/adzil/eqtest.svg)](https://pkg.go.dev/github.com/adzil/eqtest)
[![codecov](https://codecov.io/github/adzil/eqtest/graph/badge.svg?token=O54SMZGI1T)](https://codecov.io/github/adzil/eqtest)

Eqtest provides equality test assertions API using go-cmp. It is not designed to compete or replace existing assertion frameworks such as testify, but rather to complement where it lacks such as proper equality comparison for complex types.

## Quick Start

To start using `eqtest`, simply call `eqtest.New` from your Go test files:

```go
package mypkg_test

import (
    "testing"

    "github.com/adzil/eqtest"
)

func TestSomething_Simple(t *testing.T) {
    eqtest.Equal(t, "hello", "world") // This should fail the test.
}

func TestSomething_WithNew(t *testing.T) {
    eqt := eqtest.New(t)

    eqt.Equal("hello", "world") // This should fail the test.
}
```

## Using `cmp.Option`

There are three ways to add `cmp.Option` during the assertion:

### 1. During `New` initialization

This will be useful if all of the assertions require the options such as transformers.

```go
eqtest.New(t, cmp.Transformer(...))
```

### 2. Chain the option using `With`

This will be useful if there are subset of the assertions that need to have the same options such as filters.

```go
eqt := eqtest.New(t).With(cmpopts.IgnoreFields(MyStruct{}, "FieldOne"))

eqt.Equal(...)
eqt.Equal(...)
```

### 3. During the call to `Equal` or `MustEqual`

This will be useful if an assertion needs a specific filter that doesn't needed by the others.

```go
eqtest.New(t).Equal(expected, actual,
    cmpopts.IgnoreFields(MyStruct{}, "FieldTwo"),
)
```

## Reusing `cmp.Option`s from existing `Assertion`

If there are subtests that sharing the same `cmp.Option` as the current test, the existing `Assertion` can be cloned easily using a different `*testing.T` by calling `Assertion.Using`:

```go
func TestSomething(t *testing.T) {
    eqt := eqtest.New(t,
        cmpopts.SomeOption(...),
        cmpopts.SomeOption(...),
    )

    t.Run("subtest with identical cmp options", func(t *testing.T) {
        eqt := eqt.With(t)

        // Now the eqt will refer to the subtest t instead of the parent t.
        eqt.Equal(...)
    })
}
```
