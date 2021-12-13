package data_types

import "time"

type InputRecord interface {
	// Processes an input record and returns an output record.
	myfunc(time.Time) interface{}
}
