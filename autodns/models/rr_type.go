package models

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type RRType string

const (
	A     = RRType("A")
	AAAA  = RRType("AAAA")
	MX    = RRType("MX")
	CNAME = RRType("CNAME")
	NS    = RRType("NS")
	PTR   = RRType("PTR")
	TXT   = RRType("TXT")
	HINFO = RRType("HINFO")
	SPF   = RRType("SPF")
	SRV   = RRType("SRV")
	NAPTR = RRType("NAPTR")
)

var valid_types = []RRType{
	A, AAAA, MX, CNAME, NS, PTR, TXT, HINFO, SRV, NAPTR,
}

func NewRRType(v string) (RRType, error) {
	value := strings.ToUpper(v)
	if !valid(value) {
		return "", fmt.Errorf("invalid RR type %s", value)
	}
	return RRType(value), nil
}

func valid(value string) bool {
	for _, r := range valid_types {
		if r == RRType(value) {
			return true
		}
	}
	return false
}

func (rrType *RRType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	d.DecodeElement(&v, &start)
	t, err := NewRRType(v)
	if err != nil {
		return err
	}
	rrType = &t
	return nil
}

func (rrType *RRType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(*rrType, start)
}
