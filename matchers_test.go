package matchers

import (
	"testing"
)

func TestEqualTo(t *testing.T) {
	m := EqualTo{true}
	if !m.Match(true) {
		t.Error("true is not true")
	}
	if m.Match(false) {
		t.Error("true is false")
	}
}

func TestAssertThat(t *testing.T) {
	f := new(testing.T)
	AssertThat(f, true, EqualTo{true})
	if f.Failed() {
		t.Error("true is not true")
	}
	AssertThat(f, true, EqualTo{false})
	if !f.Failed() {
		t.Error("true is false")
	}
}

func TestIs(t *testing.T) {
	AssertThat(t, true, Is{true})
	AssertThat(t, true, Is{EqualTo{true}})
	AssertThat(t, Expect{true, Is{false}}, Fails{})
}

func TestNot(t *testing.T) {
	AssertThat(t, true, Not{false})
	AssertThat(t, Expect{true, Not{true}}, Fails{})
}

func TestAllOf(t *testing.T) {
	AssertThat(t, true, AllOf{Is{true}, Not{false}})
	AssertThat(t, Expect{true, AllOf{Is{false}, Not{true}}}, Fails{})
	AssertThat(t, Expect{true, AllOf{Is{true}, Not{true}}}, Fails{})
	AssertThat(t, Expect{true, AllOf{Is{false}, Not{true}}}, Fails{})
}

func TestAnyOf(t *testing.T) {
	AssertThat(t, true, AnyOf{Is{true}, Not{true}})
	AssertThat(t, true, AnyOf{Is{false}, Not{false}})
	AssertThat(t, true, AnyOf{Is{true}, Not{false}})
	AssertThat(t, Expect{true, AnyOf{Is{false}, Not{true}}}, Fails{})
}

func TestFails(t *testing.T) {
	AssertThat(t, Expect{Expect{true, Is{true}}, Fails{}}, Fails{})
}

func TestFailsPanic(t *testing.T) {
	defer func() {
		e := recover()
		AssertThat(t, e, Is{Not{nil}})
	}()
	AssertThat(t, true, Fails{})
}