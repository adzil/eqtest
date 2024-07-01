# Eqtest - Equality test assertions with go-cmp

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
    eqtest.Equal(t, "hello", "world")
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
