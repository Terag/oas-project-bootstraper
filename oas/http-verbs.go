package oas

import "encoding/json"

type HttpVerb int

const (
	GET = iota
	POST
	PUT
	PATCH
	DELETE
	HEAD
	OPTIONS
	TRACE
)

// Convert a HttpVerb the associated string
func (p HttpVerb) String() string {
	return [...]string{
		"GET",
		"POST",
		"PUT",
		"PATCH",
		"DELETE",
		"HEAD",
		"OPTIONS",
		"TRACE",
	}[p]
}

// StringToHttpVerb Convert a string to a HttpVerb type
func StringToHttpVerb(s string) HttpVerb {
	return map[string]HttpVerb{
		"GET":  GET,
		"POST": POST,
		"PUT":  PUT,
		"PATCH": PATCH,
		"DELETE": DELETE,
		"HEAD": HEAD,
		"OPTIONS": OPTIONS,
		"TRACE": TRACE,
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
