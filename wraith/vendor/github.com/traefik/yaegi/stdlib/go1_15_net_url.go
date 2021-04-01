// Code generated by 'yaegi extract net/url'. DO NOT EDIT.

// +build go1.15,!go1.16

package stdlib

import (
	"net/url"
	"reflect"
)

func init() {
	Symbols["net/url"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Parse":           reflect.ValueOf(url.Parse),
		"ParseQuery":      reflect.ValueOf(url.ParseQuery),
		"ParseRequestURI": reflect.ValueOf(url.ParseRequestURI),
		"PathEscape":      reflect.ValueOf(url.PathEscape),
		"PathUnescape":    reflect.ValueOf(url.PathUnescape),
		"QueryEscape":     reflect.ValueOf(url.QueryEscape),
		"QueryUnescape":   reflect.ValueOf(url.QueryUnescape),
		"User":            reflect.ValueOf(url.User),
		"UserPassword":    reflect.ValueOf(url.UserPassword),

		// type definitions
		"Error":            reflect.ValueOf((*url.Error)(nil)),
		"EscapeError":      reflect.ValueOf((*url.EscapeError)(nil)),
		"InvalidHostError": reflect.ValueOf((*url.InvalidHostError)(nil)),
		"URL":              reflect.ValueOf((*url.URL)(nil)),
		"Userinfo":         reflect.ValueOf((*url.Userinfo)(nil)),
		"Values":           reflect.ValueOf((*url.Values)(nil)),
	}
}
