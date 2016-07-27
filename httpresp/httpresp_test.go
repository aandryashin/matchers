package httpresp

import (
	"net/http"
	"testing"
	. "github.com/aandryashin/matchers"
)

func TestTypeOf(t *testing.T) {
	AssertThat(t, &http.Response{StatusCode: 200}, Code{http.StatusOK})
	AssertThat(t, Expect{&http.Response{StatusCode: 500}, Code{http.StatusOK}}, Fails{})
}
