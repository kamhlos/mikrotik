/*
	mikrotik (RouterOS) reports uptime as:
	1w4d02:07:25
*/

package uptime

import (
	"fmt"
	"strconv"
	"strings"
)

// accepts a routerOS uptime string, returns an int or an error
func uptimeToSecs(s string) (int, error) {

	var up, secs, mins, hours, days, weeks int
	l := len(s)

	// minimum string length should be 8 chars (00:00:01)
	if l < 8 {
		return up, fmt.Errorf("not a valid uptime string")
	}

	// seconds range from 00 to 59
	ss := string(s[l-2:])
	secs, err := strconv.Atoi(ss)
	if err != nil || secs < 0 || secs > 59 {
		return up, fmt.Errorf("not a valid uptime string")
	}

	// minutes range from 00 to 59
	sm := string(s[l-5 : l-3])
	mins, err = strconv.Atoi(sm)
	if err != nil || mins < 0 || mins > 59 {
		return up, fmt.Errorf("not a valid uptime string")
	}

	// hours range from 00 to 23
	sh := string(s[l-8 : l-6])
	hours, err = strconv.Atoi(sh)
	if err != nil || hours < 0 || hours > 23 {
		return up, fmt.Errorf("not a valid uptime string")
	}

	// days range from 1 to 6 - if "d" is present
	if i := strings.Index(s, "d"); i != -1 {
		// d is found, i -1 is the index of the number of days
		sd := string(s[i-1])
		days, err = strconv.Atoi(sd)
		if err != nil || days < 1 || days > 6 {
			return up, fmt.Errorf("not a valid uptime string")
		}
	}

	// weeks range from 1 to 52 - if "w" is present
	if i := strings.Index(s, "w"); i != -1 {
		// w is found, 0 to i is where weeks number is
		sw := string(s[0:i])
		weeks, err = strconv.Atoi(sw)
		if err != nil || weeks < 1 || weeks > 52 {
			return up, fmt.Errorf("not a valid uptime string")
		}
	}

	// not looking for years; adds unnecessary complexity
	// TODO: breaks when a year is found

	// upsecs := w * 604800 + d * 86400 + h * 3600 + m * 60 + s
	up = secs + (60 * mins) + (3600 * hours) + (86400 * days) + (604800 * weeks)

	return up, nil
}
