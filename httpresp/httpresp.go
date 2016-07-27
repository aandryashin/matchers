package httpresp

import (
	"fmt"
	"net/http"
)

type Code struct {
	C int
}

func (m Code) Match(i interface{}) bool {
	return i.(*http.Response).StatusCode == m.C
}

func (m Code) String() string {
	return fmt.Sprintf("response code %v", m.C)
}
