package media

import (
	"database/sql/driver"
	"errors"
	"strings"
)

type Files []string

func (f *Files) Scan(src any) error {
	bytes, ok := src.([]byte)
	if !ok {
		return errors.New("src value cannot cast to []byte")
	}
	*f = strings.Split(string(bytes), ",")
	return nil
}

func (f Files) Value() (driver.Value, error) {
	if len(f) == 0 {
		return nil, nil
	}
	return strings.Join(f, ","), nil
}
