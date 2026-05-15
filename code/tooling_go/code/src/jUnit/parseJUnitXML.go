package jUnit

import shared "github.com/Zip-creations/optimize_CI_deterministic_builds/code/tooling_go/code/src/shared"


type JUnitTestsuites struct {
	Testsuites []JUnitTestsuite `xml:"testsuites"`
}

type JUnitTestsuite struct {
	Name      string     `xml:"name,attr"`
	Testcases []JUnitTestcase `xml:"testcase"`
}

type JUnitTestcase struct {
	Classname string `xml:"classname,attr,omitempty"`
	Name string `xml:"name,attr,omitempty"`
	Failure *shared.Failure `xml:"failure,omitempty"`
	Skipped *shared.Skipped `xml:"skipped,omitempty"`
}
