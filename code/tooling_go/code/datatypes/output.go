package datatypes


type Testsuites struct {
	Testsuites []Testsuite `xml:"testsuites"`
}

type Testsuite struct {
	Name      string     `xml:"name,attr"`
	Timestamp string     `xml:"timestamp,attr,omitempty"`
	Testcases []Testcase `xml:"testcase"`
}

type Testcase struct {
	Classname string `xml:"classname,attr,omitempty"`
	Name string `xml:"name,attr,omitempty"`
	QualifiedName string `xml:"qualifiedName,attr,omitempty"`
	Failure *Failure `xml:"failure,omitempty"`
	Skipped *Skipped `xml:"skipped,omitempty"`
}
