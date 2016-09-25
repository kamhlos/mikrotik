/*
	mikrotik (RouterOS) reports uptime as:
	1w4d02:07:25
*/

package uptime

import (
	"strconv"
	"errors"
)

// convert to seconds
func UptimeToSecs(up string) (int, error) {
	invalid := errors.New("not a valid uptime string:", up)
	l := len(up)

	// minimum is 00:00:00
	if l < 8 {
		return 0, invalid
	}

	// get seconds
	s, err := strconv.Atoi(string(up[l-2:]))

	// get minutes
	m, err := strconv.Atoi(string(up[l-5:l-3]))

	// get hours
	h, err := strconv.Atoi(string(up[l-8:l-6]))

	// get days
	var sd = "0"
	if l > 9 {
		sd = string(up[l-10:l-9])
	}
	d, err := strconv.Atoi(sd)

	// get weeks
	var sw = "0"
	if l > 11 {
		sw = string(up[:l-11])
	}
	w, err := strconv.Atoi(sw)

	if err != nil {
		return 0, invalid
	}
	
	upsecs := w * 604800 + d * 86400 + h * 3600 + m * 60 + s

	return upsecs, nil
}