# syslog

[![GoDoc](https://godoc.org/github.com/builtdock/builtdock/logger/syslog?status.svg)](https://godoc.org/github.com/builtdock/builtdock/logger/syslog)

Package syslog implements a syslog server library. It is based on RFC 3164, as 
such it does not properly parse packets with an RFC 5424 header format.
