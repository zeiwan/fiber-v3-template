package utils

import "time"

var Rename = &rename{}

type rename struct{}

func (r rename) DateName() string {
	return time.Now().Format(time.DateOnly)
}
