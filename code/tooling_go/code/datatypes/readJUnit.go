package datatypes


// TODO: The XML structure can have two different root elements: <testsuites> and <testsuite>. Tool must be able to process both.
// Currently, <testsuite> is assumed as root
type JUnitTestsuites struct {
	Testsuites []JUnitTestsuite `xml:"testsuites"`
}

type JUnitTestsuite struct {
	Name      string     `xml:"name,attr"`
	Timestamp string     `xml:"timestamp,attr,omitempty"`
	Testcases []JUnitTestcase `xml:"testcase"`
}

type JUnitTestcase struct {
	Classname string `xml:"classname,attr,omitempty"`
	Name string `xml:"name,attr,omitempty"`
	Failure *Failure `xml:"failure,omitempty"`
	Skipped *Skipped `xml:"skipped,omitempty"`
}
