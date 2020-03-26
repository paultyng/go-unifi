package unifi

import (
	"strconv"
	"strings"
)

// emptyStringInt was created due to the behavior change in
// Go 1.14 with json.Number's handling of empty string.
type emptyStringInt int

func (e *emptyStringInt) UnmarshalJSON(b []byte) error {
	if len(b) == 0 {
		return nil
	}
	if string(b) == `""` {
		return nil
	}
	var err error
	s := string(b)
	if strings.HasPrefix(s, `"`) && strings.HasSuffix(s, `"`) {
		s, err = strconv.Unquote(s)
		if err != nil {
			return err
		}
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	*e = emptyStringInt(i)
	return nil
}
