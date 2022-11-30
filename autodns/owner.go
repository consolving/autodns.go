package autodns

import "encoding/xml"

type Owner struct {
	XMLName  xml.Name `xml:"owner"`
	Chardata string   `xml:",chardata"`
	User     string   `xml:"user"`
	Context  string   `xml:"context"`
}
