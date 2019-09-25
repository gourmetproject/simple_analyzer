package main

import "github.com/gourmetproject/gourmet"

// SimpleAnalyzer is a simple analyzer for the Gourmet Project. This analyzer is an instance of
// the gourmet.Analyzer interface because it implements the requisite Filter and Analyze functions.
// More information about Gourmet analyzers can be found in the Gourmet docs at
// https://godoc.org/github.com/gourmetproject/gourmet .
type SimpleAnalyzer struct{}

// SimpleResult is a simple implementation of the gourmet.Result interface. This object is the
// return value of the SimpleAnalyzer.Analyze function. If this Analyzer identifies a Connection
// it wants to analyze, then this object will be added to that gourmet.Connection object before it
// is logged to the Gourmet sensor log file.
//
// Notice that SimpleResult is simply an integer. gourmet.Result can be anything you want, such as
// a primitive type (int, float64, bool, etc.), a map, an array, or a struct. 
type SimpleResult int


func NewAnalyzer() gourmet.Analyzer {
	return &SimpleAnalyzer{}
}

// Key is an implementation of the gourmet.Result Key function. This function tells the Gourmet
// sensor that, for every connection we analyze, add the gourmet.Result to the Connection JSON
// object with the JSON key "payload_byte_count".
func (sr *SimpleResult) Key() string {
	return "payload_byte_count"
}

// Filter is an implementation of the gourmet.Analyzer Filter function. This filter simply checks if the
// Connection has a payload.
func (sa *SimpleAnalyzer) Filter(c *gourmet.Connection) bool {
	if len(c.Payload.Bytes()) > 0 {
		return true
	}
	return false
}

// Analyze is an implementation of the gourmet.Analyzer Analyze function. This function returns
// the number of bytes in the Connection payload.
func (sa *SimpleAnalyzer) Analyze(c *gourmet.Connection) (gourmet.Result, error) {
	var result SimpleResult
	result = SimpleResult(len(c.Payload.Bytes()))
	return &result, nil
}
