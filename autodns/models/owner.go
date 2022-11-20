package models

import "encoding/xml"

type Owner struct {
	XMLName xml.Name `xml:"owner"`
	Text    string   `xml:",chardata"`
	User    string   `xml:"user"`
	Context string   `xml:"context"`
}
