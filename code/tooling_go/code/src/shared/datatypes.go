package shared


type Failure struct {
	Message string `xml:"message,attr"`
	Type string `xml:"type,attr"`
	Content string `xml:",chardata"`
}

type Skipped struct {
	Message string `xml:"message,attr"`
}
