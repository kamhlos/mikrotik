/*
	mikrotik (RouterOS) reports uptime as:
	1w4d02:07:25
*/

package uptime

import (
	"fmt"
	"strconv"
	"errors"
)

// convert to seconds
func uptimeToSecs(up string) (int, error) {
	invalid := errors.New(fmt.Sprintf("not a valid uptime string: %s", up))
	l := len(up)

	// minimum is 00:00:00
	if l < 8 {
		return 0, invalid
	}

	// get seconds
	ss := string(up[l-2:])
	s, err := strconv.Atoi(ss)

	// get minutes
	sm := string(up[l-5:l-3])
	m, err := strconv.Atoi(sm)

	// get hours
	sh := string(up[l-8:l-6])
	h, err := strconv.Atoi(sh)

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