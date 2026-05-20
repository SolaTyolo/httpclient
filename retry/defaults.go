package retry

import "time"

const (
	DefaultRetryWaitMin = 2 * time.Second
	DefaultRetryWaitMax = 10 * time.Second
	DefaultRetryMax     = 3
	DefaultTimeout      = 5 * time.Second
)
