package datatypes


type Testsuites struct {
	Testsuites []Testsuite `xml:"testsuites"`
}

type Testsuite struct {
	Name      string     `xml:"name,attr"`
	Testcases []Testcase `xml:"testcase"`
}

type Testcase struct {
	Classname string `xml:"classname,attr,omitempty"`
	Name string `xml:"name,attr,omitempty"`
	QualifiedName string `xml:"qualifiedName,attr,omitempty"`
	Result TestStatus `xml:"result,attr"`  // if status is failed or skipped, additional info is added
	Failure *Failure `xml:"failure,omitempty"`
	Skipped *Skipped `xml:"skipped,omitempty"`
}

type TestStatus string
const (
    StatusPassed  TestStatus = "passed"
    StatusFailed  TestStatus = "failed"
    StatusSkipped TestStatus = "skipped"
	StatusNotExecuted TestStatus = "notExecuted"
)

func (t Testcase) hasRun() bool {
    return t.Result != "notExecuted"
}
