package main

import (
	"net/url"
	"reflect"
	"testing"
)

func FuzzTest(f *testing.F) {
	f.Add("x=1&y=2")
	f.Fuzz(func(t *testing.T, queryStr string) {
		query, err := url.ParseQuery(queryStr)
		if err != nil {
			t.Skip()
		}

		queryStr2 := query.Encode()
		query2, err := url.ParseQuery(queryStr2)
		if err != nil {
			t.Fatalf("parseQuery fail")
		}
		if !reflect.DeepEqual(query, query2) {
			t.Errorf("ParseQuery gave")
		}

	})
}
