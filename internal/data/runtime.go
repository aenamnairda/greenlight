package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)
	quotedJsonValue := strconv.Quote(jsonValue)

	return []byte(quotedJsonValue), nil
}

func (r *Runtime) UnmarshalJSON(jsonValue []byte) error {
	unquotedJsonValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return errInvalidRuntimeFormat
	}

	parts := strings.Split(unquotedJsonValue, " ")
	if len(parts) != 2 || parts[1] != "mins" {
		return errInvalidRuntimeFormat
	}

	i, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return errInvalidRuntimeFormat
	}

	*r = Runtime(i)
	return nil
}

var errInvalidRuntimeFormat = errors.New("invalid runtime format")
