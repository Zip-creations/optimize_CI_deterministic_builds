package main

import "encoding/xml"

type TestStatus string
const (
    StatusPassed  TestStatus = "passed"
    StatusFailed  TestStatus = "failed"
    StatusSkipped TestStatus = "skipped"
)

type Report struct {
	XMLName xml.Name `xml:"report"`
	Name string `xml:"name,attr"`
	TotalFailed int `xml:"totalFailed,attr"`
	TotalSkipped int `xml:"totalSkipped,attr"`
	TestSuites []TestSuiteReport `xml:"testsuite"`
}

type TestSuiteReport struct {
	Name string `xml:"name,attr"`
	TestCases []TestCaseReport `xml:"testcase"`
}

type TestCaseReport struct {
	Name string `xml:"name,attr"`
	Result TestStatus `xml:"result,attr"`
}
