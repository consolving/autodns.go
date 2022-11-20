package autodns

import (
	"encoding/xml"
	"errors"
	"fmt"
	"strings"
)

type Code struct {
	Number        string
	Status        string
	StatusMessage string
	Error         error
}

const (
	STATUS_S  string = "Success"
	STATUS_E  string = "Error"
	STATUS_N  string = "Notification"
	STATUS_EF string = "Error Function"

	ZONE_CREATE         string = "ZONE_CREATE"
	ZONE_UPDATE         string = "ZONE_UPDATE"
	ZONE_UPDATE_BULK    string = "ZONE_UPDATE_BULK"
	ZONE_INFO           string = "ZONE_INFO"
	ZONE_COMMENT_UPDATE string = "ZONE_COMMENT_UPDATE"
)

var TASK_CODES = map[string]Code{
	ZONE_CREATE:         {Number: "0201"},
	ZONE_UPDATE:         {Number: "0202"},
	ZONE_UPDATE_BULK:    {Number: "0202001"},
	ZONE_INFO:           {Number: "0205"},
	ZONE_COMMENT_UPDATE: {Number: "0202004"},
}

func GetValidCode(key string) (Code, error) {
	if !contains(TASK_CODES, key) {
		return Code{Number: ""}, fmt.Errorf("task code invalid")
	}
	return TASK_CODES[key], nil
}

func contains(m map[string]Code, e string) bool {
	for k := range m {
		if k == e {
			return true
		}
	}
	return false
}

func (code *Code) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var err error
	var v, statusMessage, status string
	d.DecodeElement(&v, &start)
	sLen := len(v)
	number := v
	if strings.HasPrefix(v, "S") {
		status = "S"
		statusMessage = STATUS_S
		number = v[1:sLen]
	}
	if strings.HasPrefix(v, "N") {
		status = "N"
		statusMessage = STATUS_N
		number = v[1:sLen]
	}
	if strings.HasPrefix(v, "E") {
		status = "E"
		statusMessage = STATUS_E
		number = v[1:sLen]
		err = errors.New("error: Es ist ein Fehler aufgetreten")
	}
	if strings.HasPrefix(v, "EF") {
		status = "EF"
		statusMessage = STATUS_EF
		number = v[2:sLen]
		err = errors.New("error: Hinweis auf einen Fehler bei der Verarbeitung ergänzt durch den spezifischen Fehlercode")
	}

	*code = Code{Number: number, Status: status, StatusMessage: statusMessage, Error: err}
	return nil
}

func (code *Code) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(code.Status+code.Number, start)
}

func (code *Code) String() string {
	return fmt.Sprintf("%s%s", code.Status, code.Number)
}
