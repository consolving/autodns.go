package autodns

import (
	"encoding/xml"
	"strings"
)

type Request struct {
	XMLName  xml.Name `xml:"request"`
	Chardata string   `xml:",chardata"`
	Auth     *Auth
	Owner    *Owner
	Language string `xml:"language"`
	Task     *Task
}

func ParseRequest(data string) (*Request, error) {
	var replacer = strings.NewReplacer("&#xA;", "", "&#x9;", "", "\n", "", "\t", "")
	var request *Request
	data = replacer.Replace(data)
	err := xml.Unmarshal([]byte(data), &request)
	if err != nil {
		return nil, err
	}
	return request, nil
}

func (r *Request) WithTask(task *Task) *Request {
	r.Task = task
	return r
}
