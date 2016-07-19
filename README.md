# Matchers
[![Build Status](https://travis-ci.org/aandryashin/matchers.svg?branch=master)](https://travis-ci.org/aandryashin/matchers)
[![Coverage](https://codecov.io/github/aandryashin/matchers/coverage.svg)](https://codecov.io/gh/aandryashin/matchers)

**Matchers** is simple package to allow go programmers write java like tests. It is compatible with go test, self described and self tested (see matchers_test.go).

```go
package main

import (
        "testing"

        . "github.com/aandryashin/matchers"
)

func TestSample(t *testing.T) {
        // Assertion
        AssertThat(t, true, Is{true})

        // Expectation
        AssertThat(t, Expect{true, Not{true}}, Fails{})
}
```

It is easy to implement your own custom matcher and combine it with others.

```go
package main

import (
        "fmt"
        "testing"

        . "github.com/aandryashin/matchers"
)

type Contains struct {
        S string
}

func (m Contains) Match(i interface{}) bool {
        for _, v := range i.([]string) {
                if v == m.S {
                        return true
                }
        }
        return false
}

func (m Contains) String() string {
        return fmt.Sprintf("contains %v", m.S)
}

func TestContains(t *testing.T) {
        AssertThat(t, []string{"one", "two", "three"}, Contains{"two"})
        AssertThat(t, Expect{[]string{"one", "two", "three"}, Contains{"four"}}, Fails{})
}
```
