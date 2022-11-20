package models

import (
	"encoding/xml"
)

type Task struct {
	XMLName xml.Name `xml:"task"`
	Text    string   `xml:",chardata"`
	Code    Code     `xml:"code"`
	Default *Default `xml:"default"`
	Key     *TaskKey `xml:"key,omitempty"`
	Zone    *Zone    `xml:"zone"`
}

type TaskKey string

type Default struct {
	Text                  string            `xml:",chardata"`
	Comment               *Comment          `xml:"comment"`
	ResourceRecordAdds    []*ResourceRecord `xml:"rr_add"`
	ResourceRecordRemoves []*ResourceRecord `xml:"rr_rem"`
}

type Comment string

func NewTaskDefault() *Default {
	return &Default{}
}

func NewTaskWithCode(code Code) (*Task, error) {
	return &Task{
		Code: code,
	}, nil
}

func NewTaskWithKey(key string) (*Task, error) {
	code, err := GetValidCode(key)
	if err != nil {
		return nil, err
	}
	return NewTaskWithCode(code)
}

func (t *Task) WithZone(zone *Zone) *Task {
	t.Zone = zone
	return t
}

func (t *Task) WithDefault(d *Default) *Task {
	t.Default = d
	return t
}

func (d *Default) AddRecordAdd(rr *ResourceRecord) *Default {
	d.ResourceRecordAdds = append(d.ResourceRecordAdds, rr)
	return d
}

func (d *Default) AddRecordRem(rr *ResourceRecord) *Default {
	out := make([]*ResourceRecord, 0)
	for _, rra := range d.ResourceRecordAdds {
		if rra.Hash() != rr.Hash() {
			out = append(out, rra)
		}
	}
	d.ResourceRecordAdds = out
	return d
}

func (d *Default) RemRecordAdd(rr *ResourceRecord) *Default {
	d.ResourceRecordRemoves = append(d.ResourceRecordRemoves, rr)
	return d
}

func (d *Default) RemRecordRem(rr *ResourceRecord) *Default {
	out := make([]*ResourceRecord, 0)
	for _, rra := range d.ResourceRecordRemoves {
		if rra.Hash() != rr.Hash() {
			out = append(out, rra)
		}
	}
	d.ResourceRecordRemoves = out
	return d
}
