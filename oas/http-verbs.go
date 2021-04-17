package oas

import "encoding/json"

type HttpVerb int

const (
	GET = iota
	HEAD
	POST
	PUT
	DELETE
	CONNECT
	OPTIONS
	TRACE
	PATCH
)

// Convert a HttpVerb the associated string
func (p HttpVerb) String() string {
	return [...]string{
		"GET",
		"HEAD",
		"POST",
		"PUT",
		"DELETE",
		"CONNECT",
		"OPTIONS",
		"TRACE",
		"PATCH",
	}[p]
}

// StringToHttpVerb Convert a string to a HttpVerb type
func StringToHttpVerb(s string) HttpVerb {
	return map[string]HttpVerb{
		"GET":  GET,
		"HEAD": HEAD,
		"POST": POST,
		"PUT":  PUT,
		"DELETE": DELETE,
		"CONNECT": CONNECT,
		"OPTIONS": OPTIONS,
		"TRACE": TRACE,
		"PATCH": PATCH,
	}[s]
}

// UnmarshalJSON for custom HttpVerb type
func (p *HttpVerb) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*p = StringToHttpVerb(s)
	return nil
}
