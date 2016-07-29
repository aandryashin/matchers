package httpresp

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	. "github.com/aandryashin/matchers"
)

func TestCode(t *testing.T) {
	AssertThat(t, &http.Response{StatusCode: 200}, Code{http.StatusOK})
	AssertThat(t, Expect{&http.Response{StatusCode: 500}, Code{http.StatusOK}}, Fails{})
}

type JsonMessage struct {
	Message string `json:"message"`
}

func TestIsJson(t *testing.T) {
	resp := &http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewReader([]byte(`{"message" : "message"}`)))}

	var msg JsonMessage
	AssertThat(t, resp, IsJson{&msg})
	AssertThat(t, msg.Message, Is{"message"})
}

func TestIsJsonType(t *testing.T) {
	resp := &http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewReader([]byte(`{"message" : 0}`)))}
	var msg JsonMessage
	AssertThat(t, Expect{resp, IsJson{&msg}}, Fails{})
}

func TestIsJsonNonRef(t *testing.T) {
	resp := &http.Response{Body: ioutil.NopCloser(bytes.NewReader([]byte(`{}`)))}
	var msg JsonMessage
	AssertThat(t, Expect{resp, IsJson{msg}}, Fails{})
}
