package matchers

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

type Matcher interface {
	Match(interface{}) bool
	String() string
}

type Mortal interface {
	Fatal(args ...interface{})
}

type Protect struct {
	*testing.T
}

func (t Protect) Fatal(args ...interface{}) {
	t.Error(args...)
}

type EqualTo struct {
	V interface{}
}

func (m EqualTo) Match(i interface{}) bool {
	return reflect.DeepEqual(m.V, i)
}

func (m EqualTo) String() string {
	return fmt.Sprintf("equal to %v(%v)", reflect.TypeOf(m.V), m.V)
}

type Is struct {
	V interface{}
}

func (m Is) matcher() Matcher {
	switch m.V.(type) {
	case Matcher:
		return m.V.(Matcher)
	}
	return &EqualTo{m.V}
}

func (m Is) Match(i interface{}) bool {
	return m.matcher().Match(i)
}

func (m Is) String() string {
	return fmt.Sprintf("is %v", m.matcher())
}

type Not struct {
	V interface{}
}

func (m Not) Match(i interface{}) bool {
	return !Is{m.V}.Match(i)
}

func (m Not) String() string {
	return fmt.Sprintf("not %v", m.V)
}

type AllOf []Matcher

func (all AllOf) Match(v interface{}) bool {
	for _, m := range all {
		if !m.Match(v) {
			return false
		}
	}
	return true
}

func (all AllOf) String() string {
	s := []string{}
	for _, m := range all {
		s = append(s, fmt.Sprintf("%v", m))
	}
	return strings.Join(s, ", and ")
}

type AnyOf []Matcher

func (any AnyOf) Match(v interface{}) bool {
	for _, m := range any {
		if m.Match(v) {
			return true
		}
	}
	return false
}

func (any AnyOf) String() string {
	s := ""
	for i, m := range any {
		s += fmt.Sprintf("%v", m)
		if i < len(any)-1 {
			s += ", or "
		}
	}
	return s
}

type Fails struct {
}

func (m Fails) Match(i interface{}) bool {
	err := i.(Expect).Confirm()
	return err != nil
}

func (m Fails) String() string {
	return fmt.Sprintf("fails")
}

type Expect struct {
	I interface{}
	M Matcher
}

func (m Expect) String() string {
	return fmt.Sprintf("%v %v", m.I, m.M)
}

func (e Expect) Confirm() error {
	if !e.M.Match(e.I) {
		return errors.New(fmt.Sprintf("%v(%v) %v", reflect.TypeOf(e.I), e.I, e.M))
	}
	return nil
}

func AssertThat(t Mortal, i interface{}, m Matcher) {
	err := Expect{i, m}.Confirm()
	if err != nil {
		t.Fatal(fmt.Sprintf("expect that: %v", err))
	}
}
