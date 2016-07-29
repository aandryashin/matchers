package httpresp

import (
	"encoding/json"
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

type IsJson struct {
	P interface{}
}

func (m IsJson) Match(r interface{}) bool {
	err := json.NewDecoder(r.(*http.Response).Body).Decode(m.P)
	if err != nil {
		panic(err)
	}
	return true
}

func (m IsJson) String() string {
	return fmt.Sprintf("is json representation of %T", m.P)
}
