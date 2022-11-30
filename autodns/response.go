package autodns

import "encoding/xml"

type Response struct {
	XMLName  xml.Name `xml:"response"`
	Chardata string   `xml:",chardata"`
	Result   struct {
		Chardata string     `xml:",chardata"`
		Data     *Data      `xml:"data"`
		Status   *Status    `xml:"status"`
		Msg      []*Message `xml:"msg"`
	} `xml:"result"`
	StID string `xml:"stid"`
}

type Data struct {
	Chardata string `xml:",chardata"`
	Zone     *Zone  `xml:"zone"`
}

type Message struct {
	Chardata string   `xml:",chardata"`
	Text     []string `xml:"text"`
	Code     Code     `xml:"code"`
	Type     string   `xml:"type"`
	Object   struct {
		Chardata string `xml:",chardata"`
		Type     string `xml:"type"`
		Value    string `xml:"value"`
	} `xml:"object"`
}
