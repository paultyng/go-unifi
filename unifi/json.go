package unifi

import (
	"strconv"
	"strings"
)

// emptyStringInt was created due to the behavior change in
// Go 1.14 with json.Number's handling of empty string.
type emptyStringInt struct {
	val int

	// Predicate that allows (un)marshalling functions to set custom
	// logic for handling what the JSON paylod would consider an "empty"
	// value.
	//
	// Allows for sending zero `0` in the JSON payload while allowing the
	// API to specify a magic value to interpret as blank, e.g. `-1` or other
	isEmpty func(val int) bool
}

func zeroAsBlank(i int) bool {
	return i == 0
}

// Default implementation. Zero converts to/from blank string in JSON payload.
func EmptyStringIntZero(val int) emptyStringInt {
	return emptyStringInt{
		val:     val,
		isEmpty: zeroAsBlank,
	}
}

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
	*e = EmptyStringIntZero(i)
	return nil
}

func (e *emptyStringInt) MarshalJSON() ([]byte, error) {
	b := []byte(`""`)

	if !e.isEmpty(e.val) {
		b = []byte(strconv.Itoa(int(e.val)))
	}

	return b, nil
}
