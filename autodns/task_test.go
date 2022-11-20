package autodns

import (
	"testing"
)

func TestValidCodes(t *testing.T) {
	for k, v := range TASK_CODES {
		rcode, err := GetValidCode(k)
		if err != nil {
			t.Errorf("err should be nil, but was %v", err)
		}
		if v != rcode {
			t.Errorf("rcode should be %s, but was %s", v, rcode)
		}
	}
}
