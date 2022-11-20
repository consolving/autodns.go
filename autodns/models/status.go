package models

type Status struct {
	Chardata string   `xml:",chardata"`
	Code     Code     `xml:"code"`
	Text     []string `xml:"text"`
	Type     string   `xml:"type"`
}
