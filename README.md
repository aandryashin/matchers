# Matchers

**Matchers** is simple package to allow go programmers write java like tests.

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
