### uptime
RouterOS returns uptime as a string e.g. 1w4d02:07:25.
`uptime.go` converts the uptime string into seconds.
Note if uptime is more than a year, "y" char breaks the func.