package ldetesting

import (
	"time"
)

func (p *ShouldBeImportWithCustomType) unmarshalFirst(s string) (time.Time, error) {
	return time.Date(2019, 9, 4, 11, 6, 27, 0, time.UTC), nil
}

func (p *ShouldBeImportWithCustomType) unmarshalSecond(s string) (time.Time, error) {
	return time.Date(2019, 9, 4, 11, 6, 37, 0, time.UTC), nil
}
