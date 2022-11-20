package models

import (
	"crypto/sha256"
	"encoding/xml"
	"fmt"
)

type Zone struct {
	Text              string             `xml:",chardata"`
	Name              string             `xml:"name"`
	Changed           *Changed           `xml:"changed"`
	Created           *Created           `xml:"created"`
	SystemNs          *SystemNameServer  `xml:"system_ns"`
	NsAction          *NsAction          `xml:"ns_action"`
	WwwInclude        *WwwInclude        `xml:"www_include"`
	AllowTransferFrom *AllowTransferFrom `xml:"allow_transfer_from"`
	Main              *Main              `xml:"main"`
	Soa               *Soa               `xml:"soa"`
	NameServers       []*NameServer      `xml:"nserver"`
	ResourceRecords   []*ResourceRecord  `xml:"rr"`
	Free              []*Free            `xml:"free"`
	Domainsafe        *Domainsafe        `xml:"domainsafe"`
	Owner             *User              `xml:"owner"`
	UpdatedBy         *User              `xml:"updated_by"`
	NsGroup           *NsGroup           `xml:"ns_group"`
	PurgeType         *PurgeType         `xml:"purge_type"`
}

type Changed string
type Created string
type SystemNameServer string
type NsAction string
type WwwInclude string
type AllowTransferFrom string
type Domainsafe string
type NsGroup string
type PurgeType string

type User struct {
	Text    string `xml:",chardata"`
	User    string `xml:"user"`
	Context string `xml:"context"`
}

type Free struct {
	XMLName xml.Name `xml:"free"`
	Text    string   `xml:",chardata"`
}

type Main struct {
	Text  string `xml:",chardata"`
	Value string `xml:"value"`
	Ttl   string `xml:"ttl"`
}

type Soa struct {
	Text    string `xml:",chardata"`
	Level   string `xml:"level,omitempty"`
	Refresh string `xml:"refresh"`
	Retry   string `xml:"retry"`
	Expire  string `xml:"expire"`
	Ttl     string `xml:"ttl"`
	Email   string `xml:"email"`
	Default string `xml:"default"`
}

type NameServer struct {
	XMLName xml.Name `xml:"nserver"`
	Text    string   `xml:",chardata"`
	Name    string   `xml:"name"`
	Ttl     string   `xml:"ttl,omitempty"`
}

type ResourceRecord struct {
	Text  string  `xml:",chardata"`
	Name  string  `xml:"name"`
	Ttl   string  `xml:"ttl"`
	Type  *RRType `xml:"type"`
	Pref  string  `xml:"pref,omitempty"`
	Value string  `xml:"value"`
}

func NewZone(name string) (*Zone, error) {
	return &Zone{
		Name: name,
	}, nil
}

func (zone *Zone) WithSystemNS(sysNameServer SystemNameServer) *Zone {
	zone.SystemNs = &sysNameServer
	return zone
}

func (zone *Zone) WithFree(text string) *Zone {
	free := &Free{Text: text}
	zone.Free = append(zone.Free, free)
	return zone
}

func (zone *Zone) WithMain(value string, ttl string) *Zone {
	main := &Main{Value: value, Ttl: ttl}
	zone.Main = main
	return zone
}

func NewResourceRecord(params map[string]string) (*ResourceRecord, error) {
	rrType, err := NewRRType(params["type"])
	if err != nil {
		return nil, err
	}
	return &ResourceRecord{
		Name:  params["name"],
		Ttl:   params["ttl"],
		Type:  &rrType,
		Pref:  params["pref"],
		Value: params["value"],
	}, nil
}

func (rr *ResourceRecord) Hash() string {
	h := sha256.New()
	s := fmt.Sprintf("%s/%s/%s/%s/%s", rr.Name, rr.Ttl, rr.Type, rr.Pref, rr.Value)
	h.Write([]byte(s))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}
