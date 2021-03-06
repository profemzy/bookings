package helpers

import "time"

func CurrentDateTime() string {
	dt := time.Now()
	return dt.String()
}
