package autodns

import (
	"encoding/xml"
)

type Auth struct {
	XMLName  xml.Name `xml:"auth"`
	Chardata string   `xml:",chardata"`
	User     string   `xml:"user"`
	Password string   `xml:"password"`
	Context  string   `xml:"context"`
}
